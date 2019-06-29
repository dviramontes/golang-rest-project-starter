package api

import (
	"net/http"
)

type API struct {}

func New() *API {
	return &API{}
}

func (api *API) Get(w http.ResponseWriter, r *http.Request) {}
