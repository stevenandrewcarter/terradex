package models

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"log"
	"reflect"
	"time"
)

var indexName = "terradex"

type Database struct {
	Client *elastic.Client
}

func NewDatabase() (*Database, error) {
	db := Database{}
	es, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	db.Client = es
	indexParams := `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"doc":{
				"properties": {
					"Id": {
						"type":"keyword"
					},
					"CreatedDate": {
						"type":"keyword"
					}
				}
			}
		}
	}`
	exists, err := es.IndexExists(indexName).Do(context.Background())
	if err != nil {
		return nil, err
	}
	if !exists {
		// Create an index
		_, err = db.Client.CreateIndex(indexName).BodyJson(indexParams).Do(context.Background())
		if err != nil {
			return nil, err
		}
	}
	return &db, nil
}

func (d *Database) GetProjectByID(id string) (*Project, error) {
	termQuery := elastic.NewTermQuery("Id", id)
	res, err := d.Client.Search().
		Index(indexName).
		Query(termQuery).
		Sort("CreatedDate", false).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	total := res.Hits.TotalHits
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[---] %d hits; took: %dms",
		int(total),
		res.TookInMillis,
	)
	if total == 0 {
		return nil, errors.New("Could not find a Project with Id " + id)
	}
	var ptyp Project
	// Print the ID and document source for each hit.
	for _, item := range res.Each(reflect.TypeOf(ptyp)) {
		if t, ok := item.(Project); ok {
			log.Printf(" * ID=%s, %s", t.Id, t.State)
			return &t, nil
		}
	}
	return nil, nil
}

func (d *Database) GetProjectBySlug(id string) (*Project, error) {
	return nil, nil
}

func (d *Database) NewProject(project *Project) error {
	project.CreatedDate = time.Now()
	body, err := project.ToJSON()
	if err != nil {
		return err
	}
	log.Print(body)
	_, err = d.Client.Index().
		Index(indexName).
		Type("doc").
		BodyJson(project).
		Refresh("wait_for").
		Do(context.Background())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	return nil
}
