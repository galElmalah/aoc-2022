package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	// fmt.Println("Part 1")
	// fmt.Println(Part1(data))

	fmt.Println("Part 2")
	fmt.Println(Part2(data))

}

func id(x, y int) string {
	return fmt.Sprintf("(%d,%d)", x, y)
}

func Part1(raw string) int {
	grid, limit := parse(raw)
	passedThreshold := false
	c := 0

	for !passedThreshold {
		sx, sy := 500, 0
		for !grid[id(sx, sy)] {
			if sy > limit {
				grid[id(sx, sy)] = true
				passedThreshold = true
				break
			}

			if !grid[id(sx, sy+1)] {
				sy++
				continue
			}

			if grid[id(sx, sy+1)] && !grid[id(sx-1, sy+1)] {
				sy++
				sx--
				continue

			}

			if grid[id(sx, sy+1)] && !grid[id(sx+1, sy+1)] {
				sy++
				sx++
				continue
			}

			c++
			grid[id(sx, sy)] = true

		}
	}
	return c
}

func Part2(raw string) int {
	grid, limit := parse(raw)
	limit += 2
	passedThreshold := false
	c := 0
	fmt.Println(limit)
	for !passedThreshold {
		sx, sy := 500, 0
		for !grid[id(sx, sy)] {
			if sy+1 == limit {
				c++
				grid[id(sx, sy)] = true
				break
			}

			if !grid[id(sx, sy+1)] {
				sy++
				continue
			}

			if grid[id(sx, sy+1)] && !grid[id(sx-1, sy+1)] {
				sy++
				sx--
				continue

			}

			if grid[id(sx, sy+1)] && !grid[id(sx+1, sy+1)] {
				sy++
				sx++
				continue
			}
			if sx == 500 && sy == 0 {
				// fmt.Println("asdfdasf")
				c++
				return c
			}
			c++
			grid[id(sx, sy)] = true

		}
	}
	return c
}

func parse(raw string) (map[string]bool, int) {
	lines := strings.Split(string(raw), "\n")
	r := regexp.MustCompile(`(\d+),(\d+)`)
	pointGroups := [][][]int{}

	for _, l := range lines {
		m := r.FindAllString(l, -1)
		points := [][]int{}
		for _, s := range m {
			point := strings.Split(s, ",")
			x, y := point[0], point[1]

			points = append(points, []int{util.ParseInt(x), util.ParseInt(y)})
		}
		pointGroups = append(pointGroups, points)
	}

	grid := map[string]bool{}
	limit := 0
	for _, points := range pointGroups {
		for i := 0; i < len(points)-1; i++ {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[i+1][0], points[i+1][1]
			grid[id(x1, y1)] = true
			grid[id(x2, y2)] = true

			if y1 > limit {
				limit = y1
			}
			if x1 < x2 {
				for x1 < x2 {
					grid[id(x1, y1)] = true
					x1++
				}
			} else {
				for x2 < x1 {
					grid[id(x2, y1)] = true
					x2++
				}
			}

			if y1 < y2 {
				for y1 < y2 {
					grid[id(x1, y1)] = true
					y1++
				}
			} else {
				for y2 < y1 {
					grid[id(x1, y2)] = true
					y2++
				}
			}
		}
	}
	return grid, limit
}
