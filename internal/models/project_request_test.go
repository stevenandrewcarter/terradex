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

func TestGetProjectRequest(t *testing.T) {
	tests := map[string]struct {
		project *Project
	}{
		"unknownID":  {project: nil},
		"successful": {project: &Project{}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			projectRequest := ProjectRequest{}
			projectRequest.Project = tc.project
			if err := projectRequest.Bind(nil); err != nil {
				t.Error("Unexpected error")
			}
		})
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

}
