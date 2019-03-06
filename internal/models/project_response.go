package models

import "net/http"

type ProjectResponse struct {
	*Project

	// User *UserPayload `json:"user,omitempty"`

	// We add an additional field to the response here.. such as this
	// elapsed computed property
	// Elapsed int64 `json:"elapsed"`
}

func NewProjectResponse(project *Project) *ProjectResponse {
	resp := &ProjectResponse{Project: project}

	//if resp.User == nil {
	//	if user, _ := dbGetUser(resp.UserID); user != nil {
	//		resp.User = NewUserPayloadResponse(user)
	//	}
	//}

	return resp
}

func (rd *ProjectResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	// rd.Elapsed = 10
	return nil
}
