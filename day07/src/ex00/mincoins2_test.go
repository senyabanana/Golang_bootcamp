package ex00

import (
	"reflect"
	"testing"
)

func TestMinCoins2(t *testing.T) {
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
			name:     "Unsorted coins with duplicates",
			val:      15,
			coins:    []int{10, 5, 1, 1, 5},
			expected: []int{10, 5},
		},
		{
			name:     "Empty coin slice",
			val:      7,
			coins:    []int{},
			expected: []int{},
		},
		{
			name:     "Single coin type",
			val:      4,
			coins:    []int{1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "Val zero",
			val:      0,
			coins:    []int{1, 2, 5},
			expected: []int{},
		},
		{
			name:     "Coins not sorted",
			val:      6,
			coins:    []int{3, 1, 4, 2},
			expected: []int{4, 2},
		},
		{
			name:     "Large value with mixed coins",
			val:      23,
			coins:    []int{1, 3, 7, 10},
			expected: []int{10, 10, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minCoins2(tt.val, tt.coins)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("minCoins2(%d, %v) = %v; want %v", tt.val, tt.coins, result, tt.expected)
			}
		})
	}
}
