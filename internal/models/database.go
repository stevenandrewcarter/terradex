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

func NewDatabase() (Database, error) {
	db := DatabaseElastic{}
	if err := db.Initialize(); err != nil {
		return nil, err
	} else {
		return &db, nil
	}
}
