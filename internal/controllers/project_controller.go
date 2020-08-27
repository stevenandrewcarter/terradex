package controllers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/stevenandrewcarter/terradex/internal/models"
	"log"
	"net/http"
)

func GetProject(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value("project") == nil {
		log.Print("Please provide a project.")
		w.WriteHeader(400)
		return
	}
	project := r.Context().Value("project").(*models.Project)
	if project == nil {
		http.NotFound(w, r)
		return
	}
	log.Printf("[TRC] Loading Project %s", project.Id)
	if jsonBody, err := project.GetState(); err != nil {
		http.Error(w, http.StatusText(500)+" - "+err.Error(), 500)
	} else {
		w.WriteHeader(200)
		if _, err := w.Write(jsonBody.Bytes()); err != nil {
			http.Error(w, http.StatusText(500)+" - "+err.Error(), 500)
		}
	}
}

func SaveProject(w http.ResponseWriter, r *http.Request) {
	data := &models.ProjectRequest{}
	projectID := chi.URLParam(r, "projectID")
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
