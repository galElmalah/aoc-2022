package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/galElmalah/aoc-2022/ds/point"
	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	fmt.Println("Part 1 Ranges")
	fmt.Println(Part1Ranges(data))
	// fmt.Println("Part 2")
	// fmt.Println(Part2(data))

}

const ROW = 2000000

func Part1(raw string) int {
	sensors := parse(raw)
	c := 0
	cc := 0
	// 0 sensor, 1 beacon, 2 cover area
	grid := map[string]int{}
	for _, s := range sensors {
		grid[s.position.Id()] = 0
		grid[s.beacon.Id()] = 1
		dist := s.DistanceFromBeacon()
		if s.position.Y < ROW && s.position.Y+dist > ROW {
			cc += ((dist - (util.Abs(ROW - s.position.Y))) * 2) - 1

			for i := 0; i <= dist-(util.Abs(ROW-s.position.Y)); i++ {
				right, left := point.NewPoint(s.position.X+i, ROW), point.NewPoint(s.position.X-i, ROW)
				if _, ok := grid[right.Id()]; !ok {
					grid[right.Id()] = 2
					c++
				}
				if _, ok := grid[left.Id()]; !ok {
					grid[left.Id()] = 2
					c++
				}

			}
		}

		if s.position.Y > ROW && s.position.Y-dist < ROW {
			for i := 0; i <= dist-(util.Abs(ROW-s.position.Y)); i++ {
				right, left := point.NewPoint(s.position.X+i, ROW), point.NewPoint(s.position.X-i, ROW)
				if _, ok := grid[right.Id()]; !ok {
					grid[right.Id()] = 2
					c++
				}
				if _, ok := grid[left.Id()]; !ok {
					grid[left.Id()] = 2
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

		span := s.DistanceFromBeacon() - util.Abs(ROW-s.position.Y)
		from, to := s.position.X-span, s.position.X+span

		// Check if point X is in range of the sensor according to beacon location
		if s.IsInRange(s.position.X, ROW) {
			if s.beacon.Y == ROW {
				// account for beacons on the edge of the range and make sure to exclude them
				if s.beacon.X > s.position.X {
					ranges = append(ranges, []int{from, to - 1})
				} else {
					ranges = append(ranges, []int{from + 1, to})
				}
			} else {
				ranges = append(ranges, []int{from, to})

			}

		}

	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	currentRange := []int{ranges[0][0], ranges[0][1]}
	merged := [][]int{}
	ranges = ranges[1:]
	for i := 0; i < len(ranges); i++ {

		_, cex, nsx, nex := currentRange[0], currentRange[1], ranges[i][0], ranges[i][1]

		// This means we need to keep on merging!
		if cex >= nsx {

			currentRange[1] = util.Max(cex, nex)

			// else we can't merge so we push the range and continue
		} else {
			merged = append(merged, currentRange)
			currentRange = ranges[i]
		}

		if i == len(ranges)-1 {
			merged = append(merged, currentRange)
			break
		}
	}

	sum := 0

	for _, v := range merged {
		// + 1 to make it inclusive for each starting point
		d := util.Abs(v[0]-v[1]) + 1
		sum += d
	}

	return sum
}

func Part2(raw string) int {
	sensors := parse(raw)
	c := 0

	// 0 sensor, 1 beacon, 2 cover area
	grid := map[string]int{}
	for _, s := range sensors {
		grid[s.position.Id()] = 0
		grid[s.beacon.Id()] = 1
		dist := s.DistanceFromBeacon()
		if s.position.Y < ROW && s.position.Y+dist > ROW {
			for i := 0; i <= dist-(util.Abs(ROW-s.position.Y)); i++ {
				right, left := point.NewPoint(s.position.X+i, ROW), point.NewPoint(s.position.X-i, ROW)
				if _, ok := grid[right.Id()]; !ok {
					grid[right.Id()] = 2
					c++
				}
				if _, ok := grid[left.Id()]; !ok {
					grid[left.Id()] = 2
					c++
				}

			}
		}

		if s.position.Y > ROW && s.position.Y-dist < ROW {
			for i := 0; i <= dist-(util.Abs(ROW-s.position.Y)); i++ {
				right, left := point.NewPoint(s.position.X+i, ROW), point.NewPoint(s.position.X-i, ROW)
				if _, ok := grid[right.Id()]; !ok {
					grid[right.Id()] = 2
					c++
				}
				if _, ok := grid[left.Id()]; !ok {
					grid[left.Id()] = 2
					c++
				}

			}
		}
	}

	return c
}

type Sensor struct {
	position, beacon *point.Point
}

func newSensor(sx, sy, bx, by int) Sensor {
	return Sensor{position: point.NewPoint(sx, sy), beacon: point.NewPoint(bx, by)}
}

func manhattanDistance(p1, p2 *point.Point) int {
	return util.Abs(p1.X-p2.X) + util.Abs(p1.Y-p2.Y)
}

func (s *Sensor) DistanceFromBeacon() int {
	return manhattanDistance(s.position, s.beacon)
}

func (s *Sensor) IsInRange(x, y int) bool {
	bm := manhattanDistance(s.position, s.beacon)
	pm := manhattanDistance(s.position, point.NewPoint(x, y))
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
