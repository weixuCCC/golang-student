package client_test

import (
	"ch15/series"
	"fmt"
	"testing"
)

func TestPackage(t *testing.T) {
	ret := series.GetFibonacciSerie(5)
	fmt.Println(123)
	t.Log(ret)
}
