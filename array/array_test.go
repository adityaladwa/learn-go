package array

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 times", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("Got %d and want %d, given %v", got, want, numbers)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		actual := Sum(numbers)
		expected := 6
		if actual != expected {
			t.Errorf("Got %d and want %d, given %v", actual, expected, numbers)
		}
	})
}
