package main

import (
	"testing"

	"moneybag/ex00"
)

func BenchmarkMinCoins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex00.MinCoins(10000, []int{1, 5, 10, 25, 50, 100, 500, 1000})
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex00.MinCoins2(10000, []int{1, 5, 10, 25, 50, 100, 500, 1000})
	}
}
