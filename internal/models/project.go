package models

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"time"
)

// Projects are the container object that Terraform uses to represent a State. All resources, variables, data, etc
// are stored within the Project.
type Project struct {
	Id          string                 `json:"id"`
	State       map[string]interface{} `json:"state,omitempty"`
	CreatedDate time.Time              `json:"created_date"`
	Username    string                 `json:"username,omitempty"`
	Type        string                 `json:"type"`
}

// Convert the given interface into a JSON string.
func convertToJSON(item interface{}) (*bytes.Buffer, error) {
	if b, err := json.Marshal(item); err != nil {
		return nil, err
	} else {
		return bytes.NewBuffer(b), nil
	}
}

// Convert the entire Project to a JSON string, this is used to store the Project into the database
func (e *Project) ToJSON() (*bytes.Buffer, error) {
	return convertToJSON(e)
}

// Retrieve the State of the Project. State is passed back to Terraform after it has locked the Project.
func (e *Project) GetState() (*bytes.Buffer, error) {
	return convertToJSON(e.State)
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
	return connectToDatabase(func(db Database) error {
		if locked, err := db.HasLockForID(e.Id); err != nil {
			return err
		} else if !locked {
			return db.WriteLock(e)
		}
		return nil
	})
}

func (e *Project) Unlock() error {
	return connectToDatabase(func(db Database) error {
		if locked, err := db.HasLockForID(e.Id); err != nil {
			return err
		} else if locked {
			return db.DeleteLockByID(e.Id)
		}
		return nil
	})
}

func (e *Project) Save() error {
	return connectToDatabase(func(db Database) error {
		return db.NewProject(e)
	})
}

func GetProjectById(id string) (*Project, error) {
	var project *Project
	db, err := NewDatabase()
	if err != nil {
		return nil, err
	} else {
		if project, err = db.GetProjectByID(id); err != nil {
			log.Printf("No existing project exists for %s - %s", id, err.Error())
		}
		return project, nil
	}
}
