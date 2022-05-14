package client

import (
	"time"

	"github.com/cocm1324/system-view/pkg/request"
	"github.com/cocm1324/system-view/pkg/service"
)

// Client struct is to hold data for client component
// - Eps stands for event per second.
// - Target hold pointer to target.
type SClient struct {
	DurationMilli int
	Target        *service.Service
	Requests      map[int]request.Request
	lastIssued    int
	kill          chan bool
	killed        chan bool
	reqDone       chan int
}

func NewSClient(target *service.Service, durationMilli int) *SClient {
	requests := make(map[int]request.Request)
	return &SClient{
		lastIssued:    0,
		DurationMilli: durationMilli,
		Target:        target,
		Requests:      requests,
	}
}

func (s *SClient) Start() {
	s.kill = make(chan bool)
	s.killed = make(chan bool)
	s.reqDone = make(chan int, 10)
	go s.start()
}

func (s *SClient) start() {
L:
	for {
		select {
		case done := <-s.reqDone:
			// do something before delete
			delete(s.Requests, done)
			s.Request()
		case <-s.kill:
			break L
		default:
			s.Request()
		}
		time.Sleep(time.Duration(s.DurationMilli) * time.Millisecond)
	}

	s.killed <- true
}

func (s *SClient) Request() {
	s.Requests[s.lastIssued] = *request.New()
	go s.request(s.lastIssued)
	s.lastIssued++
}

func (s *SClient) request(id int) {
	<-s.Requests[id].Done
	s.reqDone <- id
}

func (s *SClient) Kill() {
	s.kill <- true
	<-s.killed
	close(s.kill)
	close(s.killed)
}
