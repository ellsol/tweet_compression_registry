package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

const apiVersionURL = "/api/v1"

type Api struct {
	Router *chi.Mux
	App    *App
}

func NewApi(app *App) Api {
	api := Api{}
	api.App = app

	api.Router = newRouter(&api)
	return api
}

func newRouter(api *Api) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Auth"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	routes := chi.NewRouter()
	routes.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("works"))
	})

	routes.Route("/tweet", func(r chi.Router) {
		r.Get("/", api.App.Controller.HandlePaginateTweets)
		r.Post("/", api.App.Controller.UploadTweet)
	})

	routes.Route("/tweet/bychecksum", api.App.Controller.RetrieveTweet())

	r.Mount(apiVersionURL, routes)

	return r
}
