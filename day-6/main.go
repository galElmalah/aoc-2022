package main

import (
	"fmt"
	"strings"

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
	chars := parse(raw)
	for i := range chars {
		set := map[string]bool{}
		for _, c := range chars[i : i+offset] {
			set[c] = true
		}

		if len(set) == offset {
			return i + offset
		}
	}
	return -1
}

func parse(raw string) []string {
	lines := strings.Split(string(raw), "")
	return lines
}
