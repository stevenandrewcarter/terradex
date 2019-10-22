package models

import "testing"

func TestDatabaseElastic_Initialize(t *testing.T) {
	db := DatabaseElastic{}
	err := db.Initialize()
	if err != nil {
		t.Errorf("Expected the Project to cast to JSON")
	}
}
