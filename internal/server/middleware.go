package server

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	"github.com/stevenandrewcarter/terradex/internal/models"
)

func ConfigCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("[TRC] Initializing Config Context...")
		host := viper.Get("elasticsearch.host")
		if host == nil {
			host = viper.Get("elasticsearch_host")
		}
		log.Print(host)
		ctx := context.WithValue(r.Context(), "elasticsearch.host", host)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Initialize the connection to the data store as a middleware context call. This means that if the database is not
// available on any call it will fail here first.
func DatastoreCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("[TRC] Initializing Database Context...")
		db, err := models.NewDatabase()
		if err != nil {
			log.Printf("[ERR] %s", err.Error())
			http.Error(w, http.StatusText(500)+" - "+err.Error(), 500)
			return
		}
		log.Print("[TRC] Database Initialized...")
		ctx := context.WithValue(r.Context(), "database", db)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Authenticate the request, all requests must be authenticated in order to Audit Terraform State change requests.
// TODO: The username and password will need to be validated against a Datastore as well
func AuthenticateCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, authOK := r.BasicAuth()
		if !authOK || username != "username" || password != "password" {
			log.Print("[WRN] Could not authenticate user")
			http.Error(w, http.StatusText(401), 401)
			return
		}
		log.Printf("[TRC] User '%s' authenticated", username)
		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Extract the Project ID from the URL and load the Project into the context. This must occur After the Database context
// otherwise the Project will not be able to load.
func ProjectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projectID := chi.URLParam(r, "projectID")
		log.Printf("[TRC] Extracting Project ID from the url '%s'", projectID)
		if projectID == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		log.Printf("[TRC] Attempting to load ProjectID '%s'", projectID)
		db := r.Context().Value("database").(models.Database)
		project, err := db.GetProjectByID(projectID)
		if err != nil {
			log.Printf("[ERR] %s", err.Error())
			http.Error(w, http.StatusText(500)+" - "+err.Error(), 500)
			return
		}
		ctx := context.WithValue(r.Context(), "project", project)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
