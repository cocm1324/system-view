package request

// Request type is request object
type Request struct {
	StatusCode int
	Done       chan bool
}

func New(timeoutMilli int) *Request {
	done := make(chan bool)
	return &Request{
		Done: done,
	}
}

func (r *Request) Response(statusCode int) {
	r.StatusCode = statusCode
	r.Done <- true
}
