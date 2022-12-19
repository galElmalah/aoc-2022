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
	yl = 0
	xl = 7

	for i := 0; i < 2022; i++ {
		shapes = getShapes(yl)
		cs := shapes.Dequeue()
		cs.move(point.D)
		var move point.Direction
		if move = point.R; moves[i] == ">" {
			move = point.L
		}
		cs.move()
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

func (s *Shape) LowestY() int {
	lowest := s.points[0].Y
	for _, p := range s.points {
		lowest = util.Max(lowest, p.Y)
	}
	return lowest
}

func getShapes(y int) queue.Queue[*Shape] {
	Q := queue.Queue[*Shape]{}
	s1 := Shape{points: []*point.Point{point.NewPoint(2, y-3), point.NewPoint(3, y-3), point.NewPoint(4, y-3), point.NewPoint(5, y-3)}}
	s2 := Shape{points: []*point.Point{point.NewPoint(2, y-3), point.NewPoint(3, y-3), point.NewPoint(3, y-31), point.NewPoint(3, y-3), point.NewPoint(4, y-3)}}
	s3 := Shape{points: []*point.Point{point.NewPoint(4, y-3), point.NewPoint(4, y-3), point.NewPoint(4, y-3), point.NewPoint(3, y-3), point.NewPoint(2, y-3)}}
	s4 := Shape{points: []*point.Point{point.NewPoint(0, y-3), point.NewPoint(0, y-3), point.NewPoint(0, y-3), point.NewPoint(0, y-3)}}
	s5 := Shape{points: []*point.Point{point.NewPoint(2, y-3), point.NewPoint(3, y-3), point.NewPoint(2, y-3), point.NewPoint(3, y-3)}}
	for _, s := range []Shape{s1, s2, s3, s4, s5} {
		Q.Enqueue(&s)
	}
	return Q
}

func parse(raw string) (shapes queue.Queue[*Shape], moves []string) {
	shapes = getShapes(0)
	moves = strings.Split(string(raw), "")

	return shapes, moves
}
