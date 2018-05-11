package sum

import "testing"

func TestSum(t *testing.T){
	numbers := []int{10, 4, 7, 9}
	actual := Sum(numbers)
	expected := 30

	if actual != expected {
		t.Errorf("Expected '%d', but got '%d', inputs: '%v'", expected, actual, numbers)
	}
}