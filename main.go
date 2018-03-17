package main

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/oklog/oklog/pkg/group"
)

type API struct {
	action chan func()
}

type StateMachine struct {
	state   string
	actionc chan func()
}

func main() {
	// Init
	sm := newStateMachine()
	api := API{sm.actionc}
	h := http.NewServeMux()
	h.HandleFunc("/", api.indexHandler)

	l, _ := net.Listen("tcp", ":8080")

	ctx, cancel := context.WithCancel(context.Background())

	var g group.Group

	g.Add(func() error {
		return sm.Run(ctx)
	}, func(error) {
		cancel()
	})

	g.Add(func() error {
		return http.Serve(l, h)
	}, func(error) {
		l.Close()
	})

	g.Run()
}

func newStateMachine() *StateMachine {
	sm := &StateMachine{
		"initial",
		make(chan func()),
	}

	return sm
}

func (sm *StateMachine) Run(ctx context.Context) error {
	for {
		select {
		case f := <-sm.actionc:
			f()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

/*
func (sm *StateMachine) index() int {
	c := make(chan int)
	sm.actionc <- func() {
		if sm.state == "initial" {
			sm.state = "index"
		}
		c <- 1
	}
	return <-c
}
*/

func (a *API) indexHandler(w http.ResponseWriter, r *http.Request) {
	ok := make(chan []byte)

	a.action <- func() {

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
