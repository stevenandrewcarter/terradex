package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

type Database struct {
	Client *elasticsearch.Client
}

func NewDatabase() (*Database, error) {
	db := Database{}
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	db.Client = es
	return &db, nil
}

func (d *Database) GetProjectByID(id string) (*Project, error) {
	query := `
{
  "query" : { 
    "match" : { "id" : "%s" } 
  },
  "sort": [ { "created_date": { "order": "asc"} } ]
}`
	queryFmt := fmt.Sprintf(query, id)
	res, err := d.Client.Search(
		d.Client.Search.WithContext(context.Background()),
		d.Client.Search.WithIndex("test"),
		d.Client.Search.WithBody(strings.NewReader(queryFmt)),
		d.Client.Search.WithTrackTotalHits(true),
		d.Client.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		} else if res.StatusCode == 404 {
			return nil, nil
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			return nil, errors.New((e["error"].(map[string]interface{})["reason"]).(string))
		}
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	hits := r["hits"]
	total := hits.(map[string]interface{})["total"].(float64)
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(total), // int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	if total == 0 {
		return nil, errors.New("Could not find a Project with Id " + id)
	}
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		project := &Project{Id: id}
		project.State = hit.(map[string]interface{})["_source"].(map[string]interface{})["State"].(map[string]interface{})
		return project, nil
	}
	return nil, nil
}

func (d *Database) GetProjectBySlug(id string) (*Project, error) {
	return nil, nil
}

func (d *Database) NewProject(project *Project) error {
	body, err := project.toJSON()
	if err != nil {
		return err
	}
	log.Print(body)
	req := esapi.IndexRequest{
		Index:   "terradex",
		Body:    body,
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), d.Client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d, %s", res.Status(), 1, res.Body)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			return errors.New("Error parsing the response body: " + err.Error())
		} else {
			// Print the response status and indexed document version.
			return errors.New("[" + res.Status() + "] " + r["result"].(string) + "; version=" + strconv.Itoa(int(r["_version"].(float64))))
		}
	}
	return res.Body.Close()
}
