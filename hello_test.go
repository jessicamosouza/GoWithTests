package main

import "testing"

// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// In order to use the *testing.T type, you need to import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Jéssica", "Portuguese")
		want := "Olá, Jéssica"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Chloé", "French")
		want := "Bonjour, Chloé"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in japanese", func(t *testing.T) {
		got := Hello("Aiko", "Japanese")
		want := "こんにちは, Aiko"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	{
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}
