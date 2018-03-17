package state

import "context"

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
