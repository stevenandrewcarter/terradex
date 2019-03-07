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
					"id": {
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
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("id", id))
	query.MustNot(elastic.NewTermQuery("type", "lock"))
	res, err := d.Client.Search().
		Index(indexName).
		Query(query).
		Sort("created_date", false).
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
			log.Printf(" * Id=%s, %s", t.Id, t.State)
			return &t, nil
		}
	}
	return nil, nil
}

func (d *Database) DeleteLockByID(id string) error {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("id", id))
	query.Filter(elastic.NewTermQuery("type", "lock"))
	_, err := d.Client.DeleteByQuery().
		Index(indexName).
		Type("doc").
		Query(query).
		Do(context.Background())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	return nil
}

func (d *Database) GetLockByID(id string) error {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("id", id))
	query.Filter(elastic.NewTermQuery("type", "lock"))
	res, err := d.Client.Search().
		Index(indexName).
		Query(query).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return err
	}
	total := res.Hits.TotalHits
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[---] %d hits; took: %dms",
		int(total),
		res.TookInMillis,
	)
	if total == 0 {
		return errors.New("Could not find a Project with Id " + id)
	}
	return nil
}

func (d *Database) WriteLock(id string) error {
	project := Project{
		Id:          id,
		Type:        "lock",
		CreatedDate: time.Now(),
	}
	body, err := project.ToJSON()
	if err != nil {
		return err
	}
	log.Printf("Writing new Lock Entry - %s", body)
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

func (d *Database) GetProjectBySlug(id string) (*Project, error) {
	return nil, nil
}

func (d *Database) NewProject(project *Project) error {
	project.CreatedDate = time.Now()
	body, err := project.ToJSON()
	if err != nil {
		return err
	}
	log.Printf("Writing new Project Entry - %s", body)
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
