package models

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"log"
	"reflect"
	"sync"
	"time"
)

type DatabaseElastic struct {
	Client    *elastic.Client
	IndexName string `default:"terradex"`
}

var lock = &sync.Mutex{}

func (d *DatabaseElastic) newClient() error {
	log.Print("[TRC] Initializing new Elastic Client...")
	lock.Lock()
	defer lock.Unlock()
	if d.Client == nil {
		host := "http://localhost:9200"
		log.Printf("[TRC] No existing client. Creating new Elastic Client %s", host)
		err := errors.New("")
		d.Client, err = elastic.NewClient(
			elastic.SetURL(host),
			elastic.SetSniff(false))
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DatabaseElastic) createIndex(indexName string) error {
	log.Printf("[TRC] Validating Index '%s'", d.IndexName)
	ctx := context.Background()
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
					"created_date": {
						"type":"keyword"
					}
				}
			}
		}
	}`
	if exists, err := d.Client.IndexExists(d.IndexName).Do(ctx); err != nil {
		return err
	} else {
		if !exists {
			log.Print("[TRC] Index does not exist, creating new index...")
			// Create an index
			if _, err = d.Client.CreateIndex(d.IndexName).BodyJson(indexParams).Do(ctx); err != nil {
				return err
			}
		}
		return nil
	}
}

func (d *DatabaseElastic) Initialize() error {
	if err := d.newClient(); err != nil {
		return err
	} else {
		if d.IndexName == "" {
			d.IndexName = "terradex"
		}
		return d.createIndex(d.IndexName)
	}
}

func (d *DatabaseElastic) GetProjectByID(id string) (*Project, error) {
	log.Printf("[TRC] Searching for a Project with ID '%s'", id)
	query := elastic.NewBoolQuery().
		Must(elastic.NewTermQuery("id", id)).
		MustNot(elastic.NewTermQuery("type", "lock"))
	res, err := d.Client.Search().
		Index(d.IndexName).
		Query(query).
		Sort("created_date", false).
		From(0).Size(1).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	total := res.Hits.TotalHits
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[TRC] %d hits; took: %dms",
		int(total),
		res.TookInMillis,
	)
	if total == 0 {
		return nil, nil
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

func (d *DatabaseElastic) DeleteLockByID(id string) error {
	query := elastic.NewBoolQuery().
		Must(elastic.NewTermQuery("id", id)).
		Filter(elastic.NewTermQuery("type", "lock"))
	_, err := d.Client.DeleteByQuery().
		Index(d.IndexName).
		Type("doc").
		Query(query).
		Do(context.Background())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	return nil
}

// Search for a Lock for the given project id. If a lock has been created due to another instance of Terraform
// running, then it will not be possible to execute any Terraform changes until the Lock is released.
func (d *DatabaseElastic) HasLockForID(id string) (bool, error) {
	log.Printf("[TRC] Searching for existing lock for Project %s", id)
	query := elastic.NewBoolQuery().
		Must(elastic.NewTermQuery("id", id)).
		Filter(elastic.NewTermQuery("type", "lock"))
	res, err := d.Client.Search().
		Index(d.IndexName).
		Query(query).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return false, err
	}
	total := res.Hits.TotalHits
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[TRC] %d hits; took: %dms",
		int(total),
		res.TookInMillis,
	)
	return total == 0, nil
}

func (d *DatabaseElastic) WriteLock(project *Project) error {
	if body, err := project.ToJSON(); err != nil {
		return err
	} else {
		log.Printf("Writing new Lock Entry - %s", body)
		_, err = d.Client.Index().
			Index(d.IndexName).
			Type("doc").
			BodyJson(project).
			Refresh("wait_for").
			Do(context.Background())
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
	}
	return nil
}

func (d *DatabaseElastic) NewProject(project *Project) error {
	project.CreatedDate = time.Now()
	if body, err := project.ToJSON(); err != nil {
		return err
	} else {
		log.Printf("Writing new Project Entry - %s", body)
		_, err = d.Client.Index().
			Index(d.IndexName).
			Type("doc").
			BodyJson(project).
			Refresh("wait_for").
			Do(context.Background())
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
	}
	return nil
}
