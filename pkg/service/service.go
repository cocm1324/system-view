// Package service is an abstraction of system component in system design
// It can represent server, database system, loadbalancer, cache, etc.
// Key characteristics of these system component is something that it takes(or input) data and outputs.
// It will have some kind of thoughput, maximum limit of processing performance, etc, just as services do in server.
package service

import (
	"fmt"
	"time"

	"github.com/cocm1324/system-view/pkg/request"
)

// Service is abstraction of service itself
//
type Service struct {
	Up               bool
	ProcessTimeMilli int
	kill             chan bool
	killed           chan bool
	runners          []Runner
}

type Runner struct {
	done   chan bool
	kill   chan bool
	killed chan bool
}

func New(processTimeMilli int) *Service {
	return &Service{
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
	s.runners = make([]Runner, 0)

	// start goroutine
	go s.start()
}

func (s *Service) Kill() {
	s.Up = false
	s.kill <- true
	<-s.killed
	close(s.killed)
}

func (s *Service) Do(r *request.Request) {
	done := make(chan bool)
	kill := make(chan bool)
	killed := make(chan bool)
	runner := Runner{
		done:   done,
		kill:   kill,
		killed: killed,
	}
	s.runners = append(s.runners, runner)
	go s.do(r, runner)
}

func (s *Service) do(r *request.Request, n Runner) {
	go func() {
		time.Sleep(time.Duration(s.ProcessTimeMilli) * time.Microsecond)
		n.done <- true
	}()
L:
	for {
		select {
		case <-n.kill:
			fmt.Printf("runner: killed\n")
			break L
		case <-n.done:
			break L
		default:
		}
		time.Sleep(500 * time.Millisecond)
	}
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
	close(s.kill)
	s.killed <- true
}
