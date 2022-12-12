package main

import (
	"testing"

	"github.com/galElmalah/aoc-2022/util"
)

func TestPart1(t *testing.T) {
	t.Run("Run Part 1 solution with example.txt as input", func(t *testing.T) {
		data := util.ReadFile("./example.txt")

		got := Part1(data)
		want := 31

		if got != want {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

}

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
