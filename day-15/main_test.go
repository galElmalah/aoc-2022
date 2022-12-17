package main

import (
	"testing"

	"github.com/galElmalah/aoc-2022/util"
)

// func TestPart1(t *testing.T) {
// 	t.Run("Run Part 1 solution with example.txt as input", func(t *testing.T) {
// 		data := util.ReadFile("./example.txt")

// 		got := Part1(data)
// 		want := 7

// 		if got != want {
// 			t.Errorf("got %+v want %+v", got, want)
// 		}
// 	})

// }

// func TestPart2(t *testing.T) {
// 	t.Run("Run Part 2 solution with example.txt as input", func(t *testing.T) {
// 		data := util.ReadFile("./example.txt")

// 		got := Part1(data)
// 		want := 7

// 		if got != want {
// 			t.Errorf("got %+v want %+v", got, want)
// 		}
// 	})

// }

func BenchmarkPart1(b *testing.B) {
	input := util.ReadFile("./input.txt")
	for n := 0; n < b.N; n++ {
		Part1(input)
	}
}

func BenchmarkPart1Ranges(b *testing.B) {
	input := util.ReadFile("./input.txt")
	for n := 0; n < b.N; n++ {
		Part1Ranges(input)
	}
}
