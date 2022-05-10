package client

import (
	"github.com/cocm1324/system-view/pkg/request"
	"github.com/cocm1324/system-view/pkg/service"
)

// Client struct is to hold data for client component
// - Eps stands for event per second.
// - Target hold pointer to target.
type Client struct {
	Eps      float64
	Target   *service.Service
	Requests []request.Request
}

func New(target *service.Service, Eps float64) *Client {
	requests := make([]request.Request, 0)
	return &Client{
		Requests: requests,
	}
}

func Request(target *service.Service) {

}

func Start() {

}

func Event() {

}

func Kill() {

}
