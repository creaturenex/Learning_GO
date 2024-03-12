package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish, when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, Mundo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Claude", "French")
		want := "Bonjour, Claude"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French, when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, Monde"
		assertCorrectMessage(t, got, want)
	})
}

// use testing.TB when using helper functions in tests, creates interface for testing.T testing.B
// t.Helper() identifies this function as a helper
// so if it fails, will point to invoked function not the helper function
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
