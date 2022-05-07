// Package service is an abstraction of system component in system design
// It can represent server, database system, loadbalancer, cache, etc.
// Key characteristics of these system component is something that it takes(or input) data and outputs.
// It will have some kind of thoughput, maximum limit of processing performance, etc, just as services do in server.
package service

import (
	"time"

	"github.com/cocm1324/system-view/pkg/request"
)

type Service struct {
	Up               bool
	ProcessTimeMilli int
	MaxRequest       int
	RebootTime       int
	AutoReboot       bool
	kill             chan bool
}

func New(up bool, processTimeMilli int, maxRequest int, rebootTime int, autoReboot bool) *Service {
	return &Service{
		Up:               up,
		ProcessTimeMilli: processTimeMilli,
		MaxRequest:       maxRequest,
		RebootTime:       rebootTime,
		AutoReboot:       autoReboot,
	}
}

func (s *Service) Process(request *request.Request, isErr bool) {
	if !s.Up {
		request.Fail()
	}
	if s.ProcessTimeMilli == 0 {
		request.Process(isErr)
	}

	go func() {
		time.Sleep(time.Duration(s.ProcessTimeMilli) * time.Millisecond)

		request.Process(isErr)
	}()

}

func (s *Service) Kill() {
	s.kill <- true
}

func (s *Service) Boot() {

}
