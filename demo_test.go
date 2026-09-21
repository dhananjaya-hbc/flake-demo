package demo

import "testing"

func TestStable(t *testing.T) {
	// always passes
}

func TestAlwaysFails(t *testing.T) {
	t.Fatal("this test is genuinely broken, not flaky")
}

var flakyCounter int

func TestFlaky(t *testing.T) {
	flakyCounter++
	if flakyCounter%2 == 0 {
		t.Fatal("flaky failure")
	}
}
