package controllers

import (
	"github.com/stevenandrewcarter/terradex/internal/models"
	"log"
	"net/http"
)

func UnlockProject(w http.ResponseWriter, r *http.Request) {
	projectID := r.Context().Value("projectID").(string)
	username := r.Context().Value("username").(string)
	log.Printf("Unlocking... %s for %s", projectID, username)
	db, err := models.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	err = db.GetLockByID(projectID)
	if err == nil {
		db.DeleteLockByID(projectID)
	} else {
		w.WriteHeader(409)
	}
	w.WriteHeader(200)
}
