package main

import (
	"testing"
	"time"
)

func TestSleepSort(t *testing.T) {
	input := []int{3, 1, 4, 2}
	result := sleepSort(input)

	expected := []int{1, 2, 3, 4}
	var output []int

	for num := range result {
		output = append(output, num)
	}

	for i, v := range expected {
		if output[i] != v {
			t.Errorf("Expected %d, got %d", v, output[i])
		}
	}
}

func TestSleepSortEmptyInput(t *testing.T) {
	var input []int
	result := sleepSort(input)

	select {
	case _, ok := <-result:
		if ok {
			t.Errorf("Expected closed channel for empty input")
		}
	case <-time.After(1 * time.Second):
		t.Errorf("Test timed out")
	}
}
