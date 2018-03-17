package main

import (
	"context"
	"net"
	"net/http"

	"github.com/oklog/oklog/pkg/group"
	"github.com/weebagency/go-api-v2/pkg/api"
	"github.com/weebagency/go-api-v2/pkg/state"
)

func main() {
	// Init
	sm := state.NewStateMachine()
	api := api.API{sm.Actionc}
	h := http.NewServeMux()
	h.HandleFunc("/", api.IndexHandler)

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
