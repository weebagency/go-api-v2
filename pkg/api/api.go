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

	//logger := log.New(os.Stdout, "server: ", log.Lshortfile)

	h.HandleFunc("/", mw.Decorate(a.indexHandler, mw.WithLogger))

	//http.Handle("/", mw.WithLogger(a.indexHandler))

	//h.HandleFunc("/", mw.WithLogger(a.indexHandler))

}
