package main

import (
	"fmt"
	"testing"
)

func InMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinBasic(t *testing.T) {
	ans := InMin(2, -2)
	if ans != -2 {
		t.Errorf("InMin(2, -2) = %d; want -2", ans)
	}
}

func TestMinTableDriven(t *testing.T) {

	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{-2, 2, -2},
		{0, 0, 0},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := InMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InMin(1, 2)
	}
}
