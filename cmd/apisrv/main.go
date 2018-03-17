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
	var g group.Group
	ctx, cancel := context.WithCancel(context.Background())
	l, _ := net.Listen("tcp", ":8080")
	sm := state.NewStateMachine()
	a := &api.API{Action: sm.Actionc}

	// Routes
	h := http.NewServeMux()
	a.RegisterRoutes(h)

	// Add components
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
