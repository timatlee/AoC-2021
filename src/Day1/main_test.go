package main

import (
	"testing"
)

func TestFind_deeper_count(t *testing.T) {
	numbers := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	result := find_deeper_count(numbers)
	if result != 7 {
		t.Errorf("Result should be 7")
	}
}

func TestAddArray(t *testing.T) {
	result := addArray([]int{1, 2})
	if result != 3 {
		t.Errorf("Result should be 3")
	}

	result = addArray([]int{1})
	if result != 1 {
		t.Errorf("Result should be 1")
	}
}

func TestFind_deeper_count_sliding_window(t *testing.T) {
	numbers := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	result := find_deeper_count_sliding_window(numbers, 3)
	if result != 5 {
		t.Errorf("Result should be 5")
	}
}
