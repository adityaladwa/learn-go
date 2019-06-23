package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	assertRepeat := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Expected %s and got %s", want, got)
		}
	}
	t.Run("Test to repeat a", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		assertRepeat(t, repeated, expected)
	})

	t.Run("Test to repeat b", func(t *testing.T) {
		repeated := Repeat("b")
		expected := "bbbbb"
		assertRepeat(t, repeated, expected)
	})
}
