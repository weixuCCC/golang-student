package error_test

import (
	"errors"
	"testing"
)

func GetFibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n should be in [0,~]")
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	s, err := GetFibonacci(-1)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(s)
	}

}
