package models

type Database interface {
	Initialize() error
	GetProjectByID(id string) (*Project, error)
	DeleteLockByID(id string) error
	HasLockForID(id string) (bool, error)
	WriteLock(project *Project) error
	NewProject(project *Project) error
}

func connectToDatabase(cb func(Database) error) error {
	db := DatabaseElastic{}
	err := db.Initialize()
	if err != nil {
		return err
	}
	return cb(&db)
}

func newDatabase() (Database, error) {
	db := DatabaseElastic{}
	return db.Initialize()
}
