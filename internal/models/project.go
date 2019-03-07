package models

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"time"
)

type Project struct {
	Id          string                 `json:"id"`
	State       map[string]interface{} `json:"state,omitempty"`
	CreatedDate time.Time              `json:"created_date"`
	Username    string                 `json:"username,omitempty"`
	Type        string                 `json:"type"`
}

func (e *Project) ToJSON() (*bytes.Buffer, error) {
	if b, err := json.Marshal(e); err != nil {
		return nil, err
	} else {
		return bytes.NewBuffer(b), nil
	}
}

func (e *Project) GetState() (*bytes.Buffer, error) {
	if b, err := json.Marshal(e.State); err != nil {
		return nil, err
	} else {
		return bytes.NewBuffer(b), nil
	}
}

func (e *Project) LoadState(reader io.Reader) error {
	if body, err := ioutil.ReadAll(reader); err != nil {
		log.Fatalf("%v", err)
	} else {
		if len(body) > 0 {
			if err = json.Unmarshal(body, &e.State); err != nil {
				log.Fatalf("%v - %v", err, string(body))
				return err
			}
		}
	}
	return nil
}

func (e *Project) Lock() error {
	if db, err := newDatabase(); err != nil {
		log.Fatal(err)
	} else {
		if err = db.GetLockByID(e.Id); err != nil {
			return db.WriteLock(e.Id)
		}
	}
	return nil
}

func (e *Project) Unlock() error {
	if db, err := newDatabase(); err != nil {
		return err
	} else {
		if err = db.GetLockByID(e.Id); err == nil {
			if err = db.DeleteLockByID(e.Id); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func (e *Project) Save() error {
	if db, err := newDatabase(); err != nil {
		log.Fatal(err)
	} else {
		return db.NewProject(e)
	}
}

func GetProjectById(id string) (*Project, error) {
	if db, err := newDatabase(); err != nil {
		return nil, err
	} else {
		var project *Project
		if project, err = db.GetProjectByID(id); err != nil {
			log.Printf("No existing project exists for %s - %s", id, err.Error())
		}
		return project, nil
	}
}
