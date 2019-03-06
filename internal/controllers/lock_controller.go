package controllers

import (
	"log"
	"net/http"
)

func LockProject(w http.ResponseWriter, r *http.Request) {
	projectID := r.Context().Value("projectID").(string)
	username := r.Context().Value("username").(string)
	log.Printf("Locking... %s for %s", projectID, username)
	w.WriteHeader(200)
}
