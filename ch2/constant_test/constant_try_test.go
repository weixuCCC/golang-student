package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wedenday
)

const (
	read = 1 << iota // 0001
	wirt             // 0010
	ccc              // 0011
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestRead(t *testing.T) {
	a := 2 // 0010
	t.Log(a&read == read, a&wirt == wirt, a&ccc == ccc)
}
