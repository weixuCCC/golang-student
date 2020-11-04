package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
	t.Log(strings.Join(parts, ""))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	i, err := strconv.Atoi("10")
	if err != nil {
		t.Log(err)
	}
	t.Log(i * i)
}
