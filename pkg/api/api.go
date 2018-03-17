package api

import (
	"net/http"
)

type API struct {
	Action chan func()
}

func (a *API) RegisterRoutes(h *http.ServeMux) {
	h.HandleFunc("/", a.indexHandler)
}
