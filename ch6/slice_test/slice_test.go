package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0))
	t.Log(cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0))
	t.Log(cap(s0))

	s1 := []int{0, 2, 3, 4, 5, 6}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 4, 7)
	t.Log(s2, len(s2), cap(s2))

	s2 = append(s2, 5)
	t.Log(s2, len(s2), cap(s2))
}

func TestSliceGorwing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(s, len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	q2 := year[3:6]
	t.Log(q2, len(q2), cap(q2))
	year[4] = "ahahaha"
	t.Log(q2, len(q2), cap(q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "emmmmm"
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	if a == nil {
		t.Log("为空")
	} else {
		t.Log("不为空")
	}
}
