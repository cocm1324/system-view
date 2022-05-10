package client

import (
	"github.com/cocm1324/system-view/pkg/loadbalancer"
	"github.com/cocm1324/system-view/pkg/request"
	"github.com/cocm1324/system-view/pkg/service"
)

type Client struct {
	requests []request.Request
}

func New() *Client {
	return &Client{}
}

func Request[C service.Service | loadbalancer.LB](target *C) {

}
