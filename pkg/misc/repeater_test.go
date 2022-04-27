package misc

import "testing"

func TestRepeater(t *testing.T) {
	got := Repeater("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("a", 5)
	}
}
