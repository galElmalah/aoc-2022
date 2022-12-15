package main

import (
	"fmt"
	"strings"

	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	// fmt.Println("Part 2")
	// fmt.Println(Part2(data))

}

func Part1(raw string) int {
	input := parse(raw)
	fmt.Println(input)

	return -1
}

func Part2(raw string) int {
	input := parse(raw)
	fmt.Println(input)

	return -1
}

func parse(raw string) []string {
	lines := strings.Split(string(raw), "")
	return lines
}
