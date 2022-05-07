package request

import "testing"

// Test_RequestGen is to test request generator features
// suit1 - issue 1 and check it issues
// suit2 - issue 10 and check if it issues with unique id
func Test_RequestGen(t *testing.T) {
	gen := New()

	test1 := gen.Issue(1)
	if test1[0].Id != 1 {
		t.Errorf("First generated id should be 1, got %d\n", test1[0].Id)
	}

	test2 := gen.Issue(10)
	check := make(map[int]bool)

	for _, v := range test2 {
		_, ok := check[v.Id]
		if ok {
			t.Errorf("Generated Id should be unique got duplicate %d\n", v.Id)
		}
		check[v.Id] = true
	}
}

// Test_Request is to test request functions
// suit1 - Do function is changing complete to true
// suit2 - Do function is changing error to true if err
func Test_Request(t *testing.T) {
	gen := New()

	test := gen.Issue(3)

	test[0].Process(false)
	if test[0].State != 1 {
		t.Errorf("State should be 1 after process with no err, got %d\n", test[0].State)
	}

	test[1].Process(true)
	if test[1].State != 2 {
		t.Errorf("State should be 2 after process with err, got %d\n", test[1].State)
	}

	test[2].Fail()
	if test[2].State != 3 {
		t.Errorf("State should be 3 after fails, got %d\n", test[2].State)
	}

	err := test[1].Fail()
	if err != nil {
		t.Errorf("Trying to fail not pending request should return err, got nil")
	}

	err = test[2].Process(true)
	if err != nil {
		t.Errorf("Trying to process not pending request should return err, got nil")
	}
}
