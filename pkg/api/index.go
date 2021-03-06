package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *API) indexHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resOK   = make(chan []byte)
		jsonErr = make(chan error)
	)
	a.sm.Actionc <- func() {
		s := struct {
			Name string
		}{
			"Service name",
		}

		res, err := json.Marshal(s)
		if err != nil {
			jsonErr <- err
			return
		}
		resOK <- res
	}

	select {
	case err := <-jsonErr:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case ok := <-resOK:
		// Respond
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", ok)
	}
}
