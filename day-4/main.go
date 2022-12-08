package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	fmt.Println("Part 2")
	fmt.Println(Part2(data))

}

func Part1(raw string) int {

	var assignments = parse(string(raw))
	count := 0
	for _, pairAssignmentRange := range assignments {
		st1 := pairAssignmentRange[0]
		et1 := pairAssignmentRange[1]
		st2 := pairAssignmentRange[2]
		et2 := pairAssignmentRange[3]
		if (st1 >= st2 && et1 <= et2) || (st2 >= st1 && et2 <= et1) {
			count++
		}

	}

	return count

}

func Part2(raw string) int {
	var assignments = parse(string(raw))
	count := 0
	for _, pairAssignmentRange := range assignments {
		st1 := pairAssignmentRange[0]
		et1 := pairAssignmentRange[1]
		st2 := pairAssignmentRange[2]
		et2 := pairAssignmentRange[3]
		if (et1 >= st2 && et1 <= et2) || (et2 >= st1 && et2 <= et1) {
			count++
		}

	}

	return count
}

func parse(raw string) [][]int {
	lines := strings.Split(string(raw), "\n")
	pairs := [][]int{}
	r := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

	for _, l := range lines {
		match := r.FindStringSubmatch(l)
		pair := []int{}
		for _, m := range match[1:] {
			num, _ := strconv.Atoi(m)
			pair = append(pair, int(num))
		}
		pairs = append(pairs, pair)
	}

	return pairs
}
