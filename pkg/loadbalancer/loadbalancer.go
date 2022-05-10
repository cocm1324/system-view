package loadbalancer

import "github.com/cocm1324/system-view/pkg/request"

type LB struct {
	Up bool

	kill   chan bool
	killed chan bool
}

func New() *LB {
	return &LB{}
}

func (l *LB) Request(r *request.Request) {

}
