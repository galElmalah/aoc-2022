package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	fmt.Println("Part 1 ranges")
	fmt.Println(Part1Ranges(data))
	// fmt.Println("Part 2")
	// fmt.Println(Part2(data))

}

const ROW = 10

func Part1(raw string) int {
	sensors := parse(raw)
	c := 0
	cc := 0
	// 0 sensor, 1 beacon, 2 cover area
	grid := map[string]int{}
	for _, s := range sensors {
		grid[s.position.id()] = 0
		grid[s.beacon.id()] = 1
		dist := s.DistanceFromBeacon()
		if s.position.y < ROW && s.position.y+dist > ROW {
			cc += ((dist - (util.Abs(ROW - s.position.y))) * 2) - 1

			for i := 0; i <= dist-(util.Abs(ROW-s.position.y)); i++ {
				right, left := newPoint(s.position.x+i, ROW), newPoint(s.position.x-i, ROW)
				if _, ok := grid[right.id()]; !ok {
					grid[right.id()] = 2
					c++
				}
				if _, ok := grid[left.id()]; !ok {
					grid[left.id()] = 2
					c++
				}

			}
		}

		if s.position.y > ROW && s.position.y-dist < ROW {
			for i := 0; i <= dist-(util.Abs(ROW-s.position.y)); i++ {
				right, left := newPoint(s.position.x+i, ROW), newPoint(s.position.x-i, ROW)
				if _, ok := grid[right.id()]; !ok {
					grid[right.id()] = 2
					c++
				}
				if _, ok := grid[left.id()]; !ok {
					grid[left.id()] = 2
					c++
				}

			}
		}
	}

	return c
}

func Part1Ranges(raw string) int {
	sensors := parse(raw)
	ranges := [][]int{}
	// 0 sensor, 1 beacon, 2 cover area
	for _, s := range sensors {

		dist := s.DistanceFromBeacon()
		span := dist - util.Abs(ROW-s.position.y)
		if s.position.y < ROW && s.position.y+dist > ROW {
			if s.beacon.y == ROW {
				if s.beacon.x > s.position.x {
					ranges = append(ranges, []int{s.position.x - span, s.position.x + span - 1})
				} else {
					ranges = append(ranges, []int{s.position.x - span + 1, s.position.x + span})
				}
			} else {
				ranges = append(ranges, []int{s.position.x - span, s.position.x + span})

			}

		}

		if s.position.y > ROW && s.position.y-dist < ROW {

			if s.beacon.y == ROW {
				if s.beacon.x > s.position.x {
					ranges = append(ranges, []int{s.position.x - span, s.position.x + span - 1})
				} else {
					ranges = append(ranges, []int{s.position.x - span + 1, s.position.x + span})
				}
			} else {
				ranges = append(ranges, []int{s.position.x - span, s.position.x + span})

			}

		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	curr := []int{ranges[0][0], ranges[0][1]}
	merged := [][]int{}
	wo := ranges[1:]
	for i := 0; i < len(wo); i++ {

		csx, cex, nsx, nex := curr[0], curr[1], wo[i][0], wo[i][1]

		if cex >= nsx {
			max := 0
			if max = cex; cex < nex {
				max = nex
			}
			curr = []int{csx, max}
		} else {
			merged = append(merged, curr)
			curr = wo[i]
		}

		if i == len(wo)-1 {
			merged = append(merged, curr)
			break
		}
	}
	s := 0

	for _, v := range merged {
		d := util.Abs(v[0]-v[1]) + 1
		s += d
	}

	return s
}

func Part2(raw string) int {
	sensors := parse(raw)
	c := 0

	// 0 sensor, 1 beacon, 2 cover area
	grid := map[string]int{}
	for _, s := range sensors {
		grid[s.position.id()] = 0
		grid[s.beacon.id()] = 1
		dist := s.DistanceFromBeacon()
		if s.position.y < ROW && s.position.y+dist > ROW {
			for i := 0; i <= dist-(util.Abs(ROW-s.position.y)); i++ {
				right, left := newPoint(s.position.x+i, ROW), newPoint(s.position.x-i, ROW)
				if _, ok := grid[right.id()]; !ok {
					grid[right.id()] = 2
					c++
				}
				if _, ok := grid[left.id()]; !ok {
					grid[left.id()] = 2
					c++
				}

			}
		}

		if s.position.y > ROW && s.position.y-dist < ROW {
			for i := 0; i <= dist-(util.Abs(ROW-s.position.y)); i++ {
				right, left := newPoint(s.position.x+i, ROW), newPoint(s.position.x-i, ROW)
				if _, ok := grid[right.id()]; !ok {
					grid[right.id()] = 2
					c++
				}
				if _, ok := grid[left.id()]; !ok {
					grid[left.id()] = 2
					c++
				}

			}
		}
	}

	return c
}

type Point struct {
	x, y int
}

func (p *Point) id() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func newPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

type Sensor struct {
	position, beacon *Point
}

func newSensor(sx, sy, bx, by int) Sensor {
	return Sensor{position: newPoint(sx, sy), beacon: newPoint(bx, by)}
}

func manhattanDistance(p1, p2 *Point) int {
	return util.Abs[int](p1.x-p2.x) + util.Abs[int](p1.y-p2.y)
}

func (s *Sensor) DistanceFromBeacon() int {
	return manhattanDistance(s.position, s.beacon)
}

func (s *Sensor) IsInRange(x, y int) bool {
	bm := manhattanDistance(s.position, s.beacon)
	pm := manhattanDistance(s.position, newPoint(x, y))
	return bm >= pm
}

func parse(raw string) (sensors []Sensor) {
	lines := strings.Split(string(raw), "\n")
	r := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	for _, l := range lines {
		m := util.ToInts(r.FindStringSubmatch(l)[1:])
		sx, sy, bx, by := m[0], m[1], m[2], m[3]
		sensors = append(sensors, newSensor(sx, sy, bx, by))
	}
	return sensors
}
