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

	fmt.Println("Part 2")
	Part2(data)

}

func Part1(raw string) int {
	instructions := parse(raw)
	x := 1
	result := 0
	ticks := 0
	for _, ci := range instructions {
		for j := 0; j < ci.cycles; j++ {
			ticks++
			if ticks%20 == 0 && ticks%40 != 0 {
				result += x * ticks
			}
		}
		x += ci.value
	}

	return result
}

func print(crt [][]string) {
	for i, r := range crt {
		fmt.Println(i, r)
	}
}

func makeCrt() [][]string {
	crt := make([][]string, 6)
	for i, _ := range crt {
		crt[i] = make([]string, 40)
	}
	return crt
}

func Part2(raw string) {
	instructions := parse(raw)
	crt := makeCrt()
	x := 1
	ticks := 0
	for _, ci := range instructions {
		for j := 0; j < ci.cycles; j++ {
			row := int(ticks / 40)
			col := ticks % 40
			d := util.Abs(col - x)
			if d < 2 {
				crt[row][col] = "#"
			} else {
				crt[row][col] = "."
			}
			ticks++
		}
		x += ci.value
	}

	print(crt)
}

type Instruction struct {
	cycles int
	value  int
}

func parse(raw string) (instructions []Instruction) {
	lines := strings.Split(string(raw), "\n")

	for _, l := range lines {
		if strings.Contains(l, "noop") {
			instructions = append(instructions, Instruction{cycles: 1, value: 0})
		} else {
			parts := strings.Split(l, " ")
			instructions = append(instructions, Instruction{cycles: 2, value: util.ParseInt(parts[1])})
		}
	}
	return instructions
}
