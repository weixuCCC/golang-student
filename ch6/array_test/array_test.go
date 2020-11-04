package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr2 := [5]int{0, 1, 2, 3, 4}
	arr3 := [...]int{0, 2, 1, 3, 5}
	arr[0] = 3
	t.Log(arr[1], arr[0])
	t.Log(arr2, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{0, 2, 1, 3, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx, val := range arr3 {
		t.Log(idx, val)
	}
}

func TestArratSection(t *testing.T) {
	arr3 := [...]int{0, 2, 1, 3, 5}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)
}
