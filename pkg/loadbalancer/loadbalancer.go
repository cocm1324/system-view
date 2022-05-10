package loadbalancer

type LB struct {
	Up     bool
	kill   chan bool
	killed chan bool
}
