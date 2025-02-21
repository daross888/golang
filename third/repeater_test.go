package third

import (
	"testing"
)

func TestRepeater(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Repeater failed: expected %s, got %s", expected, repeated)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
