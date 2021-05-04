package main

import "testing"

func TestHello(t *testing.T) {
	//cases := []string{
	//	"Hello, Chris",
	//	"hello, Chris",
	//	"HELLO, Chris",
	//	"",
	//}

	//got := Hello("Chris")
	//for _, want := range cases {
	//	if got != want {
	//		t.Errorf("got %q want %q", got, want)
	//	}
	//}

	assertCorretMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorretMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorretMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorretMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("lunar", "French")
		want := "Bonjour, lunar"
		assertCorretMessage(t, got, want)
	})
}