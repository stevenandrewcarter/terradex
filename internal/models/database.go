package models

type Database interface {
	Initialize() (Database, error)
	GetProjectByID(id string) (*Project, error)
	DeleteLockByID(id string) error
	GetLockByID(id string) error
	WriteLock(project *Project) error
	NewProject(project *Project) error
}

func newDatabase() (Database, error) {
	db := DatabaseElastic{}
	return db.Initialize()
}
