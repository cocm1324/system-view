package client

type Client interface {
	Start()
	Request()
	Kill()
}
