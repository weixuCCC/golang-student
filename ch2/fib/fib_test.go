package fib_test

import (
	"testing"
)

func TestFibTest(t *testing.T) {
	a := 0
	b := 1
	var c int
	for i := 0; i < 10; i++ {
		c = a + b
		t.Log(b)
		a = b
		b = c
	}
}
