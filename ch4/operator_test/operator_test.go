package operator_test

import "testing"

const (
	read = 1 << iota // 0001
	wirt             // 0010
	ccc              // 0011
)

func TestOperareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3, 3, 2}
	t.Log(a == b)
	t.Log(a == c)
	t.Log(b == c)
}

func TestBigClear(t *testing.T) {
	// a := 2 // 0010
	a := 7        // 0111
	a = a &^ read // 0010
	t.Log(a&read == read, a&wirt == wirt, a&ccc == ccc)
}
