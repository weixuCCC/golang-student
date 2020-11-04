package func_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a)
	t.Log(b)
}

func Sum(ops ...int) int {
	sum := 0
	for _, op := range ops {
		sum += op
	}
	return sum
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(1, 2, 3))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("有点困")
	panic("err")
	// fmt.Println("能不能出来？")
}
