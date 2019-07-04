package di

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Aditya")

	got := buffer.String()
	want := "Hello, Aditya"

	if got != want {
		t.Errorf("Got '%s' Want '%s'", got, want)
	}
}
