package controllers

import (
	"github.com/stevenandrewcarter/terradex/internal/models"
	"log"
	"net/http"
	"time"
)

func UnlockProject(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value("projectID") == nil {
		log.Print("Please provide a projectID in order to lock the project.")
		w.WriteHeader(400)
		return
	}
	if r.Context().Value("username") == nil {
		log.Print("Please provide a username in order to lock the project.")
		w.WriteHeader(400)
		return
	}
	projectID := r.Context().Value("projectID").(string)
	username := r.Context().Value("username").(string)
	log.Printf("Unlocking... %s for %s", projectID, username)
	project := models.Project{
		Id:          projectID,
		Username:    username,
		CreatedDate: time.Now(),
	}
	if err := project.Unlock(); err != nil {
		w.WriteHeader(500)
	}
	w.WriteHeader(200)
}
