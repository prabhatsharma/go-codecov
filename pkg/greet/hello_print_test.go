package greet

import "testing"

func TestHelloPrint(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Prabhat", "English")
		want := "Hello, Prabhat"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
