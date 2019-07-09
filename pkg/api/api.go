package api

import (
	"encoding/json"
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

	render.Respond(w, r, &alarms)
}

func (api *API) PostAlert(w http.ResponseWriter, r *http.Request) {
	var alarm model.Alarm

	if err := json.NewDecoder(r.Body).Decode(&alarm); err != nil {
		http.Error(w, "error decoding json body", 500)
		return
	}

	err := api.db.CreateAlarm(&alarm)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v\n", err), 500)
		return
	}

	w.Write([]byte("OK"))
}

func (api *API) Prune(w http.ResponseWriter, r *http.Request) {
	err := api.db.DeleteAllAlarms()
	if err != nil {
		http.Error(w, fmt.Sprintf("%v\n", err), 500)
		return
	}

	w.Write([]byte("OK"))
}

func (api *API) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := api.db.DeleteAlarm(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v\n", err), 500)
		return
	}

	w.Write([]byte("OK"))
}

func (api *API) Seed(w http.ResponseWriter, r *http.Request) {
	api.db.Seed()
	w.Write([]byte("OK"))
}
