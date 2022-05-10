package request

// Request type is request object
type Request struct {
	Id         int
	StatusCode int
	done       chan bool
}

func New() *Request {
	return &Request{}
}
