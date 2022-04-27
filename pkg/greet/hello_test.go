package greet

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Prabhat", "English")
		want := "Hello, Prabhat"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, in Spanish", func(t *testing.T) {
		got := Hello("Prabhat", "Spanish")
		want := "Hola, Prabhat"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

}
