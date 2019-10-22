package models

import "testing"

func TestDatabaseElastic_Initialize(t *testing.T) {
	db := DatabaseElastic{}
	err := db.Initialize()
	if err != nil {
		t.Errorf("Expect the database to be initialized!")
	}
}
