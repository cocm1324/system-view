package main

import (
	"time"

	"github.com/cocm1324/system-view/pkg/service"
)

func main() {
	s1 := service.New(1000)
	s2 := service.New(2000)
	s3 := service.New(3000)
	s4 := service.New(4000)

	s1.Start()
	s2.Start()
	s3.Start()
	s4.Start()

	time.Sleep(1 * time.Second)
	s1.Kill()
	time.Sleep(1 * time.Second)
	s2.Kill()
	time.Sleep(1 * time.Second)
	s3.Kill()
	time.Sleep(1 * time.Second)
	s4.Kill()
}
