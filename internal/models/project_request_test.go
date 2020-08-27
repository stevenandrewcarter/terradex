package models

import (
	"testing"
)

func TestProjectRequest_ToJSON(t *testing.T) {
	projectRequest := ProjectRequest{}
	_, err := projectRequest.ToJSON()
	if err != nil {
		t.Errorf("Expected the Project to cast to JSON")
	}
}

func TestProjectRequest_BindShouldFail(t *testing.T) {
	projectRequest := ProjectRequest{}
	err := projectRequest.Bind(nil)
	if err == nil {
		t.Error("Error expected")
	}
}

func TestProjectRequest_Bind(t *testing.T) {
	projectRequest := ProjectRequest{}
	projectRequest.Project = &Project{}
	err := projectRequest.Bind(nil)
	if err != nil {
		t.Error("Unexpected error")
	}
}
