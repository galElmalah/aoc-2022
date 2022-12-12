package main

import (
	"fmt"
	"strings"

	"github.com/galElmalah/aoc-2022/ds/set"
	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	fmt.Println("Part 2")
	fmt.Println(Part2(data))

}

func (p *Point) adjust(leadingPoint *Point) {
	dx := util.Abs(leadingPoint.x - p.x)
	dy := util.Abs(leadingPoint.y - p.y)

	if dx >= 2 || dy >= 2 {
		if leadingPoint.x > p.x {
			p.x++
		} else if leadingPoint.x < p.x {
			p.x--
		}

		if leadingPoint.y > p.y {
			p.y++
		} else if leadingPoint.y < p.y {
			p.y--
		}
	}
}

type Point struct {
	x, y int
}

func (p *Point) move(direction string) {
	switch direction {
	case "L":
		p.x--
	case "R":
		p.x++
	case "U":
		p.y--
	case "D":
		p.y++
	}
}

func (p *Point) id() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func newPoint(x, y int) *Point {
	return &Point{x, y}
}

func Part1(raw string) int {
	instructions := parse(raw)
	head, tail := newPoint(0, 0), newPoint(0, 0)

	visited := set.NewSimpleSet[string]()
	for _, ci := range instructions {
		for i := 0; i < ci.steps; i++ {
			visited.Add(tail.id())
			head.move(ci.direction)
			tail.adjust(head)
		}

	}

	return visited.Size()
}

func Part2(raw string) int {
	instructions := parse(raw)
	knots := make([]*Point, 10)

	for i, _ := range knots {
		knots[i] = newPoint(0, 0)
	}

	visited := set.NewSimpleSet[string]()

	for _, ci := range instructions {
		for i := 0; i < ci.steps; i++ {
			knots[0].move(ci.direction)
			for j := 0; j < len(knots)-1; j++ {
				head, tail := knots[j], knots[j+1]
				tail.adjust(head)
			}
			visited.Add(knots[len(knots)-1].id())
		}
	}

	return visited.Size()
}

type Instruction struct {
	direction string
	steps     int
}

func parse(raw string) (instructions []Instruction) {
	lines := strings.Split(string(raw), "\n")
	for _, l := range lines {
		parts := strings.Split(l, " ")
		instructions = append(instructions, Instruction{direction: parts[0], steps: util.ParseInt(parts[1])})
	}
	return instructions
}

/*
..##..
...##.
.####.
....#.
####..
*/
