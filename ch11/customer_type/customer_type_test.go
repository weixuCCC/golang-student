package customer_type_test

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(io int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestFn(t *testing.T) {
	ret := timeSpent(slowFunc)
	t.Log(ret(10))
}
