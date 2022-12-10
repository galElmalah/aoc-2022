package main

import (
	"testing"

	"github.com/galElmalah/aoc-2022/util"
)

func TestPart1(t *testing.T) {
	t.Run("Run Part 1 solution on example.txt input", func(t *testing.T) {
		data := util.ReadFile("./example.txt")

		got := Part1(data, 4)
		want := 7

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Run Part 2 solution on example.txt input", func(t *testing.T) {
		data := util.ReadFile("./example.txt")

		got := Part1(data, 14)
		want := 19

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	input := util.ReadFile("./input.txt")
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Part1(input, 4)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.ReadFile("./input.txt")
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Part1(input, 14)
	}
}
