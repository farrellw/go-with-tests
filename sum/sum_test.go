package sum

import (
	"reflect"
	"testing"
)

func assertCorrectMessage(t *testing.T, expected int, actual int, params []int) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected '%d', but got '%d', inputs: '%v'", expected, actual, params)
	}
}

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		actual := Sum(numbers)
		expected := 6

		assertCorrectMessage(t, expected, actual, numbers)
	})
}

func assertDeepEqual(t *testing.T, expected []int, actual []int) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, wanted %v", actual, expected)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	assertDeepEqual(t, want, got)
}

func TestSumAllTails(t *testing.T) {
	t.Run("Make the sums of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 2, 9})
		want := []int{2, 11}

		assertDeepEqual(t, want, got)
	})

	t.Run("Saveful sum empty slices", func(t *testing.T){
		got := SumAllTails([]int{}, []int{2,9})
		want := []int{0, 9}

		assertDeepEqual(t, want, got)
	})
}
