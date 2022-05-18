package client

import (
	"time"

	"github.com/cocm1324/system-view/pkg/request"
	"github.com/cocm1324/system-view/pkg/service"
)

// Client struct is to hold data for client component
// - Eps stands for event per second.
// - Target hold pointer to target.
type Client struct {
	DurationMilli       int
	RequestTimeoutMilli int
	Target              *service.Service
	Requests            map[int]request.Request
	lastIssued          int
	kill                chan bool
	killed              chan bool
	reqDone             chan int
}

// New will return pointer to new Client struct instance
// Param target is pointer to target service. Client will send request to cirtain service and that is target. It should be set in first place.
// durationMilli is the duration between each request. When client starts, it produces requests until it stops. Duration is time between each request
func New(durationMilli int, requestTimeoutMilli int, target *service.Service) *Client {
	requests := make(map[int]request.Request)
	return &Client{
		lastIssued:          0,
		DurationMilli:       durationMilli,
		RequestTimeoutMilli: requestTimeoutMilli,
		Target:              target,
		Requests:            requests,
	}
}

// Start will start client goroutine by setting up channels and call go c.start(). It is kind of encapsulated goroutine method that outsider caller will call to start client
func (c *Client) Start() {
	c.kill = make(chan bool)
	c.killed = make(chan bool)
	c.reqDone = make(chan int, 10)
	go c.start()
}

// start will run main event loop for this client. It will wait for request done or kill channel
// reqDone channel is channel to receive done signal of each request
// kill channel is channel to detect kill signal from outside of this client, if kill outputs, the loop will break and cleanup goroutine will be executed
// At the end of each loop it will sleep for some millisecond
func (c *Client) start() {
L:
	for {
		select {
		case done := <-c.reqDone:
			// do something before delete
			delete(c.Requests, done)
			c.Request()
		case <-c.kill:
			break L
		default:
			c.Request()
		}
		time.Sleep(time.Duration(c.DurationMilli) * time.Millisecond)
	}

	c.killed <- true
}

//
func (c *Client) Request() {
	c.Requests[c.lastIssued] = *request.New()
	c.lastIssued++
}

func (c *Client) Kill() {
	c.kill <- true
	<-c.killed
	close(c.kill)
	close(c.killed)
}
