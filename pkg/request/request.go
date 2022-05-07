package request

import "errors"

type Request struct {
	Id    int
	State int // 0 -> pending, 1 -> completed, 2 -> completed with err, 3 -> failed
}

func (r *Request) Process(isErr bool) error {
	if r.State != 0 {
		err := errors.New("request: cannot process already processed request")
		return err
	}
	if isErr {
		r.State = 2
		return nil
	}
	r.State = 1
	return nil
}

func (r *Request) Fail() error {
	if r.State != 0 {
		err := errors.New("request: cannot fail already processed request")
		return err
	}
	r.State = 3
	return nil
}

type RequestGen struct {
	LastIssued int
}

func New() *RequestGen {
	return &RequestGen{
		LastIssued: 0,
	}
}

func (g *RequestGen) Issue(count int) []Request {
	req := make([]Request, count)
	for i := range req {
		g.LastIssued = g.LastIssued + 1
		req[i] = Request{
			Id:    g.LastIssued,
			State: 0,
		}
	}
	return req
}
