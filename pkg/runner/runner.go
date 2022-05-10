package runner

import (
	"fmt"
	"time"

	"github.com/cocm1324/system-view/pkg/request"
)

type Runner struct {
	done     chan bool
	kill     chan bool
	killed   chan bool
	duration int
}

func New(duration int) *Runner {
	done := make(chan bool)
	kill := make(chan bool)
	killed := make(chan bool)
	return &Runner{
		done:     done,
		kill:     kill,
		killed:   killed,
		duration: duration,
	}
}

func (r *Runner) Run(req *request.Request) {
	tick := 100
	tickCount := r.duration / tick

L:
	for i := 0; i < tickCount; i++ {
		select {
		case <-r.kill:
			fmt.Printf("runner: killed\n")
			break L
		default:
		}
		time.Sleep(time.Duration(tick) * time.Millisecond)
	}

}

func (r *Runner) Kill() {
	r.kill <- true

}
