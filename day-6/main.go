package main

import (
	"fmt"

	"github.com/galElmalah/aoc-2022/ds/set"
	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data, 4))

	fmt.Println("Part 2")
	fmt.Println(Part1(data, 14))

}

func Part1(raw string, offset int) int {
	for i := range raw {
		set := set.NewSimpleSet[rune]()
		for _, c := range raw[i : i+offset] {
			if set.Has(c) {
				break
			}
			set.Add(c)
		}

		if set.Size() == offset {
			return i + offset
		}
	}
	return -1
}
