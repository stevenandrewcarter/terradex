package controllers

import (
	"github.com/stevenandrewcarter/terradex/internal/models"
	"log"
	"net/http"
	"time"
)

func LockProject(w http.ResponseWriter, r *http.Request) {
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
