package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s', want '%s'", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Matt", "")
		want := "Hello, Matt"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying hello to world when empty string supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Passing a second parameter with spanish language", func(t *testing.T) {
		got := Hello("Matt", "Spanish")
		want := "Hola, Matt"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Passing a second parameter with French language", func(t *testing.T) {
		got := Hello("Matt", "French")
		want := "Bonjour, Matt"
		assertCorrectMessage(t, got, want)
	})
}
