package api

import (
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
	api.db.GetAllAlarms(&alarms)
	names := normalizeAlarms(&alarms)
	w.Write([]byte(tml.Compile(names)))
}

func normalizeAlarms(alarms *[]model.Alarm) []string {
	var names []string
	for _, a := range *alarms {
		names = append(names, a.Text)
	}

	return names
}
