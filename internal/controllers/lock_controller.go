package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/stevenandrewcarter/terradex/internal/models"
)

func LockProject(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Error: %s", err.Error())
		http.Error(w, http.StatusText(409), 409)
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(project); err != nil {
			http.Error(w, http.StatusText(500), 500)
		}
	}
}
