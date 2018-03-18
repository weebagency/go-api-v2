package api

import (
	"net/http"

	mw "github.com/weebagency/go-api-v2/pkg/middleware"
	"github.com/weebagency/go-api-v2/pkg/state"
)

type API struct {
	sm *state.StateMachine
}

func NewAPI(sm *state.StateMachine) *API {
	return &API{sm}
}

func (a *API) RegisterRoutes(h *http.ServeMux) {

	h.HandleFunc("/", mw.Decorate(a.indexHandler,
		mw.WithLogger))

	//http.Handle("/", mw.WithLogger(a.indexHandler))

}
