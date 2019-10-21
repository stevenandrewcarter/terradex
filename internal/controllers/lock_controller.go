package controllers

import (
	"github.com/stevenandrewcarter/terradex/internal/models"
	"log"
	"net/http"
	"time"
)

func LockProject(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Context())
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
	log.Printf("Locking... %s for %s", projectID, username)
	project := models.Project{
		Id:          projectID,
		Username:    username,
		CreatedDate: time.Now(),
		Type:        "lock",
	}
	if err := project.Lock(); err != nil {
		w.WriteHeader(409)
	}
	w.WriteHeader(200)
}
