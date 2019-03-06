package models

import (
	"errors"
	"net/http"
)

type ProjectRequest struct {
	*Project    `json:"project,omitempty"`
	ProtectedID string `json:"id"` // override 'id' json to have more control
}

func (a *ProjectRequest) Bind(r *http.Request) error {
	if a.Project == nil {
		return errors.New("missing required Project fields")
	}
	a.ProtectedID = "" // unset the protected ID
	return nil
}
