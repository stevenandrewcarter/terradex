package models

import "net/http"

type ProjectResponse struct {
	*Project `json:"project,omitempty"`
	// We add an additional field to the response here.. such as this
	// elapsed computed property
	// Elapsed int64 `json:"elapsed"`
}

func NewProjectResponse(project *Project) *ProjectResponse {
	return &ProjectResponse{Project: project}
}

func (rd *ProjectResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	// rd.Elapsed = 10
	return nil
}
