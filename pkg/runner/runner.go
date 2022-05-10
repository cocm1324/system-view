package runner

import (
	"fmt"
	"time"

	"github.com/cocm1324/system-view/pkg/request"
)

type Runner struct {
	Id       int
	kill     chan bool
	killed   chan bool
	duration int
}

func New(id int, duration int) *Runner {
	kill := make(chan bool)
	killed := make(chan bool)
	return &Runner{
		Id:       id,
		kill:     kill,
		killed:   killed,
		duration: duration,
	}
}

func (r *Runner) Start(req *request.Request, dispatcher chan int) {
	go r.start(req, dispatcher)
}

func (r *Runner) Kill() {
	r.kill <- true
}

func (r *Runner) start(req *request.Request, dispatcher chan int) {
	tick := 100
	tickCount := r.duration / tick

	for i := 0; i < tickCount; i++ {
		select {
		case <-r.kill:
			fmt.Printf("runner: killed\n")
			req.Response(500)
			return
		default:
		}
		time.Sleep(time.Duration(tick) * time.Millisecond)
	}

	req.Response(200)
}
