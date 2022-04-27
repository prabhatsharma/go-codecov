package mt

import "testing"

func TestAdd(t *testing.T) {
	t.Run("1+1=2", func(t *testing.T) {
		got := Add(1, 1)
		want := 2

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
