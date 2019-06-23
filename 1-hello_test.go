package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s' want '%s'", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Aditya", "")
		want := "Hello, Aditya"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Aditya", "Spanish")
		want := "Hola, Aditya"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Aditya", "French")
		want := "Bonjour, Aditya"
		assertCorrectMessage(t, got, want)
	})
}
