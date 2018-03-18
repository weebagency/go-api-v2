package api

import (
	"net/http"

	"github.com/weebagency/go-api-v2/pkg/state"
)

type API struct {
	sm *state.StateMachine
}

func NewAPI(sm *state.StateMachine) *API {
	return &API{sm}
}

func (a *API) RegisterRoutes(h *http.ServeMux) {

	http.Handle("/", middleware.Set(middleware.Logger())(a.indexHandler))
	//h.HandleFunc("/", a.indexHandler)

}
