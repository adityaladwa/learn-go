package iteration

import (
	"fmt"
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
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertRepeat(t, repeated, expected)
	})

	t.Run("Test to repeat b", func(t *testing.T) {
		repeated := Repeat("b", 10)
		expected := "bbbbbbbbbb"
		assertRepeat(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	got := Repeat("a", 5)
	fmt.Println(got)
	// Output: aaaaa
}
