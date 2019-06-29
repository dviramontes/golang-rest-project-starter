package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dviramontes/golang-rest-project-starter/internal/config"
	"github.com/dviramontes/golang-rest-project-starter/pkg/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	conf := config.Read("/config.yml", map[string]interface{}{
		"ingest": true, // ingest by default
	})

	API := api.New()

	version := conf.GetString("version")
	log.Printf("version: %s\n", version)

	router := chi.NewRouter()

	// middleware setup
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,                         // log api request calls
		middleware.DefaultCompress,                // compress results, mostly gzipping assets and json
		middleware.StripSlashes,                   // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,                      // recover from panics without crashing server
		middleware.Timeout(3000*time.Millisecond), // Stop processing after 3 seconds
	)

	// obligatory health-check endpoint
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Get("/", API.Get)

	http.ListenAndServe(":3000", router)
}
