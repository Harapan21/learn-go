package main

import "testing"

var (
	a      int
	b      int
	c      int
	result int
)

func testPlus(t *testing.T) {
	a = 1
	b = 2
	result = 3

	t.Logf("Keliling : %b", plus(a, b))

	if plus(a, b) != result {
		t.Errorf("fail result must %b", result)
	}
}
