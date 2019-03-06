package controllers

import (
	"log"
	"net/http"
)

func UnlockProject(w http.ResponseWriter, r *http.Request) {
	projectID := r.Context().Value("projectID").(string)
	username := r.Context().Value("username").(string)
	log.Printf("Unlocking... %s for %s", projectID, username)
	w.WriteHeader(200)
}
