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

func TestGetProjectById_UnknownID(t *testing.T) {
	projectId := "1234"
	_, err := GetProjectById(projectId)
	if err != nil {
		t.Errorf("No Project ID with %v should be found", projectId)
	}
}

func TestGetProjectById(t *testing.T) {
	projectId := "123"
	_, err := GetProjectById(projectId)
	if err != nil {
		t.Errorf("Should be able to retrieve project with ID %v", projectId)
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
