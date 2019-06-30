package api

import (
	"net/http"

	"github.com/dviramontes/golang-rest-project-starter/pkg/pg"
)

type API struct {
	db *pg.DB
}

func New(db *pg.DB) *API {
	return &API{db}
}

func (api *API) Get(w http.ResponseWriter, r *http.Request) {}
