package interface_test

import "testing"

type Programmer interface {
	WriteHelloWord() string
}

type GoProgrammer struct{}

func (g *GoProgrammer) WriteHelloWord() string {
	return "fmt.Println(\"Hello Word\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWord())
}
