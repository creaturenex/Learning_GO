package Arrays_Slices

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	checkSumMsg := func(t testing.TB, actual, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v to be %v", expected, actual)
		}
	}
	t.Run("Make the sums of some slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSumMsg(t, actual, expected)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSumMsg(t, actual, expected)
	})
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}
	// cannot do equality operators with Slices, instead
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v to be %v", expected, actual)
	}
}

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
