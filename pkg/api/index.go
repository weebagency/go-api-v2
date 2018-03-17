package api

import (
	"encoding/json"
	"net/http"
)

func (a *API) indexHandler(w http.ResponseWriter, r *http.Request) {
	ok := make(chan []byte)

	a.Action <- func() {

		s := struct {
			Name string
		}{
			"Service name",
		}

		res, err := json.Marshal(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ok <- res
	}

	select {
	case k := <-ok:
		w.Header().Set("Content-Type", "application/json")
		w.Write(k)
	}
}
