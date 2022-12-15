package main

import (
	"testing"

	"github.com/galElmalah/aoc-2022/util"
)

func BenchmarkPart1(b *testing.B) {
	input := util.ReadFile("./input.txt")
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Part1(input)
	}
}

func BenchmarkPart2NaiveApproach(b *testing.B) {
	input := util.ReadFile("./input.txt")
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Part2NaiveApproach(input)
	}
}

func BenchmarkPart2MultiSourceBfs(b *testing.B) {
	input := util.ReadFile("./input.txt")
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Part2MultiSourceBfs(input)
	}
}
