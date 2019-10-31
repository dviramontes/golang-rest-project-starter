package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dviramontes/golang-rest-project-starter/internal/config"
	"github.com/dviramontes/golang-rest-project-starter/pkg/api"
	"github.com/dviramontes/golang-rest-project-starter/pkg/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	conf := config.Read("/config.yml", map[string]interface{}{
		"seed": true, // ingest by default
	})
	version := conf.GetString("version")
	seed := conf.GetBool("seed")
	connStr := conf.GetString("conn_str")

	log.Printf("version: %s\n", version)
	log.Printf("seeding: %t\n", seed)

	pgdb, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err, "Could not connect to postgres database")
	}
	defer pgdb.Close()

	DB := model.New(pgdb)
	API := api.New(DB)

	DB.Migrate()

	if seed {
		DB.Seed()
	}

	router := chi.NewRouter()

	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// middleware setup
	router.Use(
		cors.Handler,
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

	router.Route("/api", func(router chi.Router) {
		router.Route("/alerts", func(router chi.Router) {
			router.Get("/", API.GetAlerts)
			router.Post("/", API.PostAlert)
			router.Post("/seed", API.Seed)
			router.Put("/upvote", API.Upvote)
			router.Delete("/", API.Delete)
			router.Delete("/prune", API.Prune)
		})
	})

	http.ListenAndServe(":4000", router)
}
