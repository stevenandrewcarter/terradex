package controllers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/stevenandrewcarter/terradex/internal/models"
	"log"
	"net/http"
)

func ProjectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var project *models.Project
		var projectID string
		var err error
		db, err := models.NewDatabase()
		if err != nil {
			log.Fatal(err)
		}
		if projectID = chi.URLParam(r, "projectID"); projectID != "" {
			project, err = db.GetProjectByID(projectID)
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}
		if r.Method == "POST" {
			project = &models.Project{Id: projectID}
		} else if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "project", project)
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
	// ctx := r.Context()
	db, err := models.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	var projectID string
	var project *models.Project
	if projectID = chi.URLParam(r, "projectID"); projectID != "" {
		project, err = db.GetProjectByID(projectID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
		}
	}
	// project, ok := ctx.Value("project").(*models.Project)
	w.WriteHeader(200)
	jsonBody, err := project.GetState()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	w.Write([]byte(jsonBody.Bytes()))
}

func SaveProject(w http.ResponseWriter, r *http.Request) {
	data := &models.ProjectRequest{}
	var projectID string
	if projectID = chi.URLParam(r, "projectID"); projectID != "" {
		data.ProtectedID = projectID
		data.Project = &models.Project{Id: projectID}
		data.LoadState(r.Body)
	}
	//if err := render.Bind(r, data); err != nil {
	//	render.Render(w, r, ErrInvalidRequest(err))
	//	return
	//}
	db, err := models.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	db.NewProject(data.Project)
	render.Status(r, http.StatusCreated)
	render.Render(w, r, models.NewProjectResponse(data.Project))
	//body := Body{}
	//body.parse(r.Body)
	//log.Print(html.EscapeString(r.URL.Path), " ", r.Method, " - ", body)
	//// save(body)
	//w.WriteHeader(200)
}
