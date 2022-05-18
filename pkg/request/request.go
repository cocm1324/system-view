package request

import "time"

const TICK int = 100

// Request type is request struct that will be passed from client to server. It is abstraction of HTTP request.
// StatusCode is to hold standard http status code. If it is 0, it is not done yet.
// ClientTimeoutStatus is flag that represents if this request is timedout by client and no longer in use
// Done channel is channel to notify this request is done. Request is created by client only and when it is created, the client starts goroutine of this request. Done channel will pass signal if this request is done to client. After client receive signal, it will shutdown goroutine.
type Request struct {
	StatusCode    int
	ClientTimeout chan bool
	Done          chan bool
}

// New will return pointer to newly created Request struct.
//
func New() *Request {
	return &Request{}
}

func (r *Request) Send(timeoutMilli int, id int, reqDone chan int) {
	done := make(chan bool)
	clientTimeout := make(chan bool)
	r.Done = done
	r.ClientTimeout = clientTimeout
	go r.send(timeoutMilli, id, reqDone)
}

func (r *Request) send(timeoutMilli int, id int, reqDone chan int) {
L:
	for {
		select {
		case <-r.Done:
			break L
		case <-r.ClientTimeout:
			break L
		default:
			timeoutMilli -= TICK
			if timeoutMilli < 0 {
				r.ClientTimeout <- true
			}
		}
		time.Sleep(time.Duration(TICK) * time.Millisecond)
	}
	reqDone <- id
}

func (r *Request) Response(statusCode int) {
	r.StatusCode = statusCode
	r.Done <- true
}
