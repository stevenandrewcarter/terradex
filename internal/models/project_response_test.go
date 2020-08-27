package models

import (
	"testing"
)

func TestProjectResponse_ToJSON(t *testing.T) {
	projectResponse := ProjectResponse{}
	_, err := projectResponse.ToJSON()
	if err != nil {
		t.Errorf("Expected the Project to cast to JSON")
	}
}

func TestNewProjectResponse(t *testing.T) {
	projectResponse := NewProjectResponse(nil)
	if projectResponse == nil {
		t.Error("ProjectResponse should not be nil")
	}
}

func TestProjectResponse_Render(t *testing.T) {
	projectResponse := NewProjectResponse(nil)
	response := projectResponse.Render(nil, nil)
	if response != nil {
		t.Error("Render should return nil response")
	}
}
