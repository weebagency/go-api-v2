package state

import (
	"context"
	"log"
)

type StateMachine struct {
	state   string
	Actionc chan func()
}

func NewStateMachine() *StateMachine {
	sm := &StateMachine{
		"initial",
		make(chan func()),
	}

	return sm
}

func (sm *StateMachine) Run(ctx context.Context) error {
	for {
		select {
		case f := <-sm.Actionc:
			f()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (sm *StateMachine) Index() int {
	c := make(chan int)
	sm.Actionc <- func() {
		if sm.state == "initial" {
			sm.state = "index 5"
		}

		log.Println(sm.state)
		c <- 1
	}
	return <-c
}
