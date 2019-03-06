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
	Id          string `json:id`
	State       map[string]interface{}
	CreatedDate time.Time
}

func (e *Project) ToJSON() (*bytes.Buffer, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	log.Printf("%s", string(b))
	return bytes.NewBuffer(b), nil
}

func (e *Project) GetState() (*bytes.Buffer, error) {
	b, err := json.Marshal(e.State)
	if err != nil {
		return nil, err
	}
	log.Printf("%s", string(b))
	return bytes.NewBuffer(b), nil
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
