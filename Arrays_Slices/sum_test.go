package Arrays_Slices

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		actual := Sum(numbers)
		expected := 15
		assertCorrectMsg(t, actual, expected)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		actual := Sum(numbers)
		expected := 6
		assertCorrectMsg(t, actual, expected)
	})
}

func assertCorrectMsg(t testing.TB, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %d to be %d", expected, actual)
	}
}
