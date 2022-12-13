package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./example.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	// fmt.Println("Part 2")
	// fmt.Println(Part2(data))

}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func compare(left, right []float64) {
	for i, a := range left {
		if typeof(a) == "int" && typeof(right[i]) == "int" {
			println(a)
		}
	}
}

func Part1(raw string) int {
	pairs := parse(raw)
	for _, p := range pairs {
		fmt.Println(p)
		compare(p[0], p[1])
	}

	return -1
}

func Part2(raw string) int {
	input := parse(raw)
	fmt.Println(input)

	return -1
}

func parse(raw string) [][]any {
	var pairs [][]any
	lines := strings.Split(string(raw), "\n\n")
	for _, l := range lines {
		var left []any
		var right []any
		parts := strings.Split(l, "\n")
		json.Unmarshal([]byte(parts[0]), &left)
		json.Unmarshal([]byte(parts[1]), &right)

		pairs = append(pairs, []any{left, right})
	}
	return pairs
}
