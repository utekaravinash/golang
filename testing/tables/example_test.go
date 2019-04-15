package main

import "testing"

// table of testes with input numbers and their expected output
var tests = []struct {
	numbers  []int
	expected int
}{
	{
		numbers:  []int{4},
		expected: 4,
	},
	{
		numbers:  []int{1, 2},
		expected: 3,
	},
	{
		numbers:  []int{1, 2, 3, 4},
		expected: 10,
	},
}

// TestAdd tests add function
func TestAdd(t *testing.T) {

	for _, test := range tests {
		sum := Add(test.numbers...)
		if sum != test.expected {
			t.Errorf("sum of %v was expected to be %v, instead got %v", test.numbers, test.expected, sum)
		}
	}

}
