package main

import (
	"fmt"
	"strings"

	"github.com/galElmalah/aoc-2022/ds/point"
	"github.com/galElmalah/aoc-2022/ds/queue"
	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./example.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	// fmt.Println("Part 2")
	// fmt.Println(Part2(data))

}

func Part1(raw string) int {
	shapes, moves := parse(raw)
	// grid := map[string]int
	xl := 7

	for i := 0; i < 2022; i++ {

	}
	return -1
}

// func Part2(raw string) int {
// 	input := parse(raw)
// 	fmt.Println(input)

// 	return -1
// }

type Shape struct {
	points []*point.Point
}

func (s *Shape) move(direction point.Direction) {
	for _, p := range s.points {
		p.MoveMutate(direction)
	}
}

// func (s *Shape) bounds() (top, right, down, left *point.Point) {

// 	for _, p := range s.points {
// 		if top.Y > p.Y {
// 			top = p
// 		}
// 		if right.X > p.Y {
// 			top = p
// 		}
// 	}†
// }

func getShapes() queue.Queue[*Shape] {
	Q := queue.Queue[*Shape]{}
	s1 := Shape{points: []*point.Point{point.NewPoint(2, 0), point.NewPoint(3, 0), point.NewPoint(4, 0), point.NewPoint(5, 0)}}
	s2 := Shape{points: []*point.Point{point.NewPoint(2, 0), point.NewPoint(3, 0), point.NewPoint(3, -1), point.NewPoint(3, 1), point.NewPoint(4, 0)}}
	s3 := Shape{points: []*point.Point{point.NewPoint(4, 0), point.NewPoint(4, 1), point.NewPoint(4, 2), point.NewPoint(3, 2), point.NewPoint(2, 2)}}
	s4 := Shape{points: []*point.Point{point.NewPoint(0, 0), point.NewPoint(0, 1), point.NewPoint(0, 2), point.NewPoint(0, 3)}}
	s5 := Shape{points: []*point.Point{point.NewPoint(2, 0), point.NewPoint(3, 0), point.NewPoint(2, 1), point.NewPoint(3, 1)}}
	for _, s := range []Shape{s1, s2, s3, s4, s5} {
		Q.Enqueue(&s)
	}
	return Q
}

func parse(raw string) (shapes []Shape, moves []string) {
	shapes = getShapes()
	moves = strings.Split(string(raw), "")
	return shapes, moves
}
