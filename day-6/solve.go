package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./input.txt")
	check(err)

	fmt.Println("Part 1")
	fmt.Println(part1(data, 4))

	fmt.Println("Part 2")
	fmt.Println(part1(data, 14))

}

func part1(raw []byte, offset int) int {
	chars := parse(string(raw))
	for i := 0; i < len(chars); i++ {
		duplicates := false
		seen := map[string]bool{}
		for j := i; j < i+offset; j++ {
			if seen[chars[j]] {
				duplicates = true
				break
			}
			seen[chars[j]] = true
		}

		if !duplicates {
			return i + offset
		}
	}
	return -1
}

func parse(raw string) []string {
	lines := strings.Split(string(raw), "")
	return lines
}
