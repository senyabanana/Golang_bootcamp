package ex00

import (
	"reflect"
	"testing"
)

func TestMinCoins(t *testing.T) {
	tests := []struct {
		name     string
		val      int
		coins    []int
		expected []int
	}{
		{
			name:     "Basic example with exact coins",
			val:      13,
			coins:    []int{1, 5, 10},
			expected: []int{10, 1, 1, 1},
		},
		{
			name:     "Exact match with large coin",
			val:      15,
			coins:    []int{1, 5, 10},
			expected: []int{10, 5},
		},
		{
			name:     "Unsorted coins",
			val:      13,
			coins:    []int{10, 1, 5},
			expected: []int{10, 1, 1, 1},
		},
		{
			name:     "Single coin type",
			val:      3,
			coins:    []int{1},
			expected: []int{1, 1, 1},
		},
		{
			name:     "Empty coin slice",
			val:      7,
			coins:    []int{},
			expected: []int{},
		},
		{
			name:     "Val zero",
			val:      0,
			coins:    []int{1, 5, 10},
			expected: []int{},
		},
		{
			name:     "Coins with duplicates",
			val:      7,
			coins:    []int{1, 1, 5, 10, 10},
			expected: []int{5, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minCoins(tt.val, tt.coins)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("minCoins(%d, %v) = %v; want %v", tt.val, tt.coins, result, tt.expected)
			}
		})
	}
}
