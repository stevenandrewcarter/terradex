package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/stevenandrewcarter/terradex/internal/controllers"
)

func Routes() *chi.Mux {
	chi.RegisterMethod("LOCK")
	chi.RegisterMethod("UNLOCK")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(controllers.AuthenticateCtx)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Route("/{projectID}", func(r chi.Router) {
		r.Use(controllers.ProjectCtx)
		r.MethodFunc("LOCK", "/", controllers.LockProject)
		r.MethodFunc("UNLOCK", "/", controllers.UnlockProject)
		r.Get("/", controllers.GetProject)
		r.Post("/", controllers.SaveProject)
	})
	return r
}
