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
	tests := map[string]struct {
		projectID string
	}{
		"unknownID":  {projectID: "1234"},
		"successful": {projectID: "123"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := GetProjectById(tc.projectID)
			if err != nil {
				t.Errorf("Should be able to retrieve project with ID %v", tc.projectID)
			}
		})
	}
}

func TestProject_GetState(t *testing.T) {
	project := Project{}
	state, _ := project.GetState()
	if state == nil {
		t.Error("State should not be nil")
	}
}

func TestProject_Save(t *testing.T) {
	project := Project{}
	err := project.Save()
	if err != nil {
		t.Error("Project Save should not be nil")
	}
}

func TestProject_Lock(t *testing.T) {
	project := Project{}
	err := project.Lock()
	if err != nil {
		t.Error("Project should have been locked")
	}
}

func TestProject_LockCannotLockTwice(t *testing.T) {
	project := Project{}
	project.Lock()
	err := project.Lock()
	if err != nil {
		t.Error("Project should have been locked")
	}
}

func TestProject_Unlock(t *testing.T) {
	project := Project{}
	err := project.Unlock()
	if err != nil {
		t.Error("Project should have been unlocked")
	}
}
