package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./input.txt")
	exampleData, _ := os.ReadFile("./example.txt")
	_ = exampleData
	check(err)

	fmt.Println("Part 1")

	fmt.Println(part1(data))

	fmt.Println("Part 2")
	fmt.Println(part2(data))

}

func part1(raw []byte) string {

	stacks, instructions := parse(string(raw))

	// mutating the stacks
	for _, instruction := range instructions {
		moveFrom := stacks[instruction.from]
		moveTo := stacks[instruction.to]
		for i := 0; i < instruction.amount; i++ {
			moveTo = append(moveTo, moveFrom[len(moveFrom)-1])
			moveFrom = moveFrom[:len(moveFrom)-1]
		}
		stacks[instruction.from] = moveFrom
		stacks[instruction.to] = moveTo
	}

	answer := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			answer += stack[len(stack)-1]
		}
	}

	return answer

}

func part2(raw []byte) string {
	stacks, instructions := parse(string(raw))
	// mutating the stacks
	for _, instruction := range instructions {
		takeRange := len(stacks[instruction.from]) - instruction.amount

		stacks[instruction.to] = append(stacks[instruction.to], stacks[instruction.from][takeRange:]...)
		stacks[instruction.from] = stacks[instruction.from][:takeRange]
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

func parse(raw string) ([][]string, []instruction) {
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

type instruction struct {
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

func parseInstructions(rawInstructions []string) (instructions []instruction) {
	matcher := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for _, line := range rawInstructions {
		match := toInts(matcher.FindStringSubmatch(line)[1:])
		instructions = append(instructions, instruction{
			amount: match[0],
			from:   match[1] - 1,
			to:     match[2] - 1,
		})
	}

	return instructions
}
