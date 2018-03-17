package api

import (
	"net/http"
)

type API struct {
	Action chan func()
}

func NewAPI(ac chan func()) *API {
	return &API{ac}
}

func (a *API) RegisterRoutes(h *http.ServeMux) {
	h.HandleFunc("/", a.indexHandler)
}
