package controllers

import (
	"github.com/stevenandrewcarter/terradex/internal/models"
	"log"
	"net/http"
	"time"
)

func UnlockProject(w http.ResponseWriter, r *http.Request) {
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
