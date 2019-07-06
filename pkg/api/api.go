package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dviramontes/golang-rest-project-starter/pkg/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type API struct {
	db *model.DB
}

func New(db *model.DB) *API {
	return &API{db}
}

// **************
// * ALERTS API *
// **************

func (api *API) GetAlerts(w http.ResponseWriter, r *http.Request) {
	var alarms []model.Alarm

	err := api.db.GetAllAlarms(&alarms)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("%v\n", err), 500)
		return
	}

	names := normalizeAlarms(&alarms)
	render.Respond(w, r, names)
}

func (api *API) Prune(w http.ResponseWriter, r *http.Request) {
	err := api.db.DeleteAllAlarms()
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("%v\n", err), 500)
		return
	}

	w.Write([]byte("OK"))
}

func (api *API) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := api.db.DeleteAlarm(id)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("%v\n", err), 500)
		return
	}

	w.Write([]byte("OK"))
}

func (api *API) Seed(w http.ResponseWriter, r *http.Request) {
	api.db.Seed()
	w.Write([]byte("OK"))
}

func normalizeAlarms(alarms *[]model.Alarm) []string {
	var names []string
	for _, a := range *alarms {
		names = append(names, a.Text)
	}

	return names
}

// *************
// * TODOS API *
// *************

func (api *API) GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []model.Todo

	err := api.db.GetAllTodos(&todos)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("%v\n", err), 500)
		return
	}

	render.Respond(w, r, todos)
}
