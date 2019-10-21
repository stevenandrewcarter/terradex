package models

import (
	"testing"
)

func TestProject_ToJSON(t *testing.T) {
	project := Project{}
	_, err := project.ToJSON()
	if err != nil {
		t.Errorf("Expected the Project to cast to JSON")
	}
}

func TestGetProjectById(t *testing.T) {
	_, err := GetProjectById("123")
	if err != nil {
		t.Errorf("Should be able to retrieve project with ID 123")
	}
}