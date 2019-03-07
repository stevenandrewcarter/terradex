package controllers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/stevenandrewcarter/terradex/internal/models"
	"log"
	"net/http"
)

func AuthenticateCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := false
		username := "ANONYMOUS"
		if auth {
			username, password, authOK := r.BasicAuth()
			if authOK == false {
				http.Error(w, "Not authorized", 401)
				return
			}
			if username != "username" || password != "password" {
				http.Error(w, "Not authorized", 401)
				return
			}
			log.Printf("User %s authenticated", username)
		}
		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ProjectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var projectID string
		projectID = chi.URLParam(r, "projectID")
		if projectID == "" {
			render.Render(w, r, ErrNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "projectID", projectID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	projectID := r.Context().Value("projectID").(string)
	if project, err := models.GetProjectById(projectID); err != nil {
		http.Error(w, http.StatusText(500), 500)
	} else {
		if project != nil {
			if jsonBody, err := project.GetState(); err != nil {
				http.Error(w, http.StatusText(500), 500)
			} else {
				w.WriteHeader(200)
				w.Write([]byte(jsonBody.Bytes()))
			}
		}
	}
}

func SaveProject(w http.ResponseWriter, r *http.Request) {
	data := &models.ProjectRequest{}
	projectID := r.Context().Value("projectID").(string)
	data.ProtectedID = projectID
	data.Project = &models.Project{
		Id:       projectID,
		Username: r.Context().Value("username").(string),
		Type:     "state",
	}
	if err := data.LoadState(r.Body); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	if err := data.Save(); err != nil {
		http.Error(w, http.StatusText(500), 500)
	} else {
		render.Status(r, http.StatusCreated)
		render.Render(w, r, models.NewProjectResponse(data.Project))
	}
}
