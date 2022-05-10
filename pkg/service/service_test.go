package service

import "testing"

func TestServiceRun(t *testing.T) {
	s := New(5000)

	if s.Up != false {
		t.Errorf("service: after new, Up should be false. Got %t\n", s.Up)
	}

	s.Start()
	if s.Up != true {
		t.Errorf("service: after Start, Up should be true. Got %t\n", s.Up)
	}

	s.Kill()
	if s.Up != false {
		t.Errorf("service: after Kill, Up should be false. Got %t\n", s.Up)
	}
}

func TestServiceRequest(t *testing.T) {

}
