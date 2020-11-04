package map_ext_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = false
	mySet[3] = false
	delete(mySet, 3)
	for i := 1; i <= len(mySet); i++ {
		if mySet[i] {
			t.Logf("%d is true", i)
		} else {
			t.Logf("%d is false", i)
		}
	}
}
