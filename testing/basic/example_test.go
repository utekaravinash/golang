package main

import "testing"

// TestAdd tests add function
func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum != 3 {
		t.Errorf("Sum should be 3, got %v", sum)
	}
}
