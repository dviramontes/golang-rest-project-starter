package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dviramontes/golang-rest-project-starter/pkg/model"
	"github.com/dviramontes/golang-rest-project-starter/tml"
)

type API struct {
	db *model.DB
}

func New(db *model.DB) *API {
	return &API{db}
}

func (api *API) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var alarms []model.Alarm

	err := api.db.GetAllAlarms(&alarms)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("%v\n", err), 500)
		return
	}

	names := normalizeAlarms(&alarms)
	w.Write([]byte(tml.Compile(names)))
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

func normalizeAlarms(alarms *[]model.Alarm) []string {
	var names []string
	for _, a := range *alarms {
		names = append(names, a.Text)
	}

	return names
}
