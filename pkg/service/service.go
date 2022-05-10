// Package service is an abstraction of system component in system design
// It can represent server, database system, loadbalancer, cache, etc.
// Key characteristics of these system component is something that it takes(or input) data and outputs.
// It will have some kind of thoughput, maximum limit of processing performance, etc, just as services do in server.
package service

import (
	"fmt"
	"time"

	"github.com/cocm1324/system-view/pkg/request"
	"github.com/cocm1324/system-view/pkg/runner"
)

// Service is abstraction of service itself
//
type Service struct {
	Up               bool
	ProcessTimeMilli int
	lastIssued       int
	runners          map[int]runner.Runner
	kill             chan bool
	killed           chan bool
	runnerDone       chan int
}

func New(processTimeMilli int) *Service {
	return &Service{
		lastIssued:       0,
		ProcessTimeMilli: processTimeMilli,
	}
}

func (s *Service) Start() {
	// prepare struct before start goroutine
	kill := make(chan bool)
	killed := make(chan bool)

	s.Up = true
	s.kill = kill
	s.killed = killed
	s.runners = make(map[int]runner.Runner)

	// start goroutine
	go s.start()
}

func (s *Service) Kill() {
	s.Up = false
	s.kill <- true
	<-s.killed
	close(s.kill)
	close(s.killed)
}

func (s *Service) Request(r *request.Request) {
	s.runners[s.lastIssued] = *runner.New(s.lastIssued, s.ProcessTimeMilli)
	s.lastIssued += 1
}

func (s *Service) start() {
L:
	for {
		// do something
		select {
		case <-s.kill:
			// do something before destroy
			fmt.Printf("service: killed\n")
			break L
		default:
			// do something
		}
		time.Sleep(500 * time.Millisecond)
	}

	s.killed <- true
}
