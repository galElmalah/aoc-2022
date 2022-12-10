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

func Part1(raw string) string {

	stacks, instructions := parse(string(raw))

	// mutating the stacks
	for _, instruction := range instructions {
		from := instruction.from
		to := instruction.to
		// another approach is to create a slice of size amount from `moveFrom`
		// reverse that slice and push it to `moveTo` but the approach here is much simpler to reason about
		for i := 0; i < instruction.amount; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	answer := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			answer += stack[len(stack)-1]
		}
	}

	return answer

}

func Part2(raw string) string {
	stacks, instructions := parse(string(raw))
	// mutating the stacks
	for _, instruction := range instructions {
		from := instruction.from
		to := instruction.to
		amount := instruction.amount
		takeRange := len(stacks[from]) - amount
		// take items from `takeRange` until the end of the slice and append them to the target stack
		stacks[to] = append(stacks[to], stacks[from][takeRange:]...)
		// remove items that come after the `takeRange` from our source crate
		stacks[from] = stacks[from][:takeRange]
	}

	answer := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			answer += stack[len(stack)-1]
		}
	}

	return answer

}

func chunkBy(items []string, chunkSize int) (chunks [][]string) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}

func parse(raw string) ([][]string, []Instruction) {
	chunks := strings.Split(string(raw), "\n\n")
	rawCrates := strings.Split(chunks[0], "\n")
	rawInstructions := strings.Split(chunks[1], "\n")

	return parseStacks(rawCrates), parseInstructions(rawInstructions)
}

func parseStacks(crates []string) [][]string {
	stacks := make([][]string, 9)
	for _, row := range crates {
		rowOfCrates := chunkBy(strings.Split(row, ""), 4)
		for crateNo, crateCandidate := range rowOfCrates {
			for _, char := range crateCandidate {
				if char >= "A" && char <= "Z" {
					// pre appending an element to array
					stacks[crateNo] = append([]string{char}, stacks[crateNo]...)
				}
			}
		}
	}
	return stacks
}

type Instruction struct {
	amount int
	from   int
	to     int
}

func toInts(fromStrings []string) (result []int) {
	for _, n := range fromStrings {
		num, _ := strconv.Atoi(n)
		result = append(result, num)
	}
	return result
}

func parseInstructions(rawInstructions []string) (instructions []Instruction) {
	matcher := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for _, line := range rawInstructions {
		match := toInts(matcher.FindStringSubmatch(line)[1:])
		instructions = append(instructions, Instruction{
			amount: match[0],
			from:   match[1] - 1,
			to:     match[2] - 1,
		})
	}

	return instructions
}
