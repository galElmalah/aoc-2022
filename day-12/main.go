package main

import (
	"fmt"
	"strings"

	"github.com/galElmalah/aoc-2022/ds/queue"
	"github.com/galElmalah/aoc-2022/ds/set"
	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	fmt.Println("Part 2 - Naive approach")
	fmt.Println(Part2NaiveApproach(data))

	fmt.Println("Part 2 - Multi-source BFS")
	fmt.Println(Part2MultiSourceBfs(data))
}

func Part1(raw string) int {
	graph, start, dest := parse(raw)
	steps := BFS(graph, start, dest)
	return steps
}

func Part2NaiveApproach(raw string) int {
	graph, start, dest := parse2(raw)
	steps := []int{}
	for _, s := range start {
		res := BFS(graph, s, dest)
		if res > 0 {
			steps = append(steps, res)
		}
	}

	return getMinValue(steps)
}

func Part2MultiSourceBfs(raw string) int {
	graph, start, dest := parse2(raw)
	steps := MultiSourceBFS(graph, start, dest)

	return steps
}

func BFS(graph [][]*Point, s *Point, destination *Point) int {
	Q := queue.Queue[*Point]{}
	Q.Enqueue(s)
	backTrack := map[string]string{}
	seen := set.NewSimpleSet[string]()

	seen.Add(s.id())
	for !Q.IsEmpty() {
		currentNode := Q.Dequeue()
		if currentNode == destination {
			break
		}
		neighbors := getNeighbors(graph, currentNode)

		for _, v := range neighbors {
			if !seen.Has(v.id()) {
				seen.Add(v.id())
				backTrack[v.id()] = currentNode.id()
				Q.Enqueue(v)
			}
		}
	}
	return count(backTrack, destination.id())
}

func getNeighbors(graph [][]*Point, sink *Point) (neighbors []*Point) {
	moves := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, move := range moves {
		di, dj := move[0]+sink.i, move[1]+sink.j
		if di >= 0 && di < len(graph) && dj >= 0 && dj < len(graph[0]) {
			delta := graph[di][dj].v - sink.v
			if delta <= 1 {
				neighbors = append(neighbors, graph[di][dj])
			}
		}
	}
	return neighbors
}

func count(backTrack map[string]string, id string) int {
	v, ok := backTrack[id]
	if ok {
		return 1 + count(backTrack, v)
	}
	return 0
}

func parse(raw string) (matrix [][]*Point, start *Point, dest *Point) {
	lines := strings.Split(string(raw), "\n")
	matrix = make([][]*Point, len(lines))
	for i := range matrix {
		matrix[i] = make([]*Point, len(lines[0]))
		rows := []*Point{}
		for j, c := range lines[i] {
			if c == 'S' {
				start = createPoint(i, j, c)
			}
			if c == 'E' {
				dest = createPoint(i, j, c)
			}
			rows = append(rows, createPoint(i, j, c))
		}
		matrix[i] = rows
	}

	return matrix, start, dest
}

func MultiSourceBFS(graph [][]*Point, s []*Point, destination *Point) int {
	Q := queue.Queue[*Point]{}
	seen := set.NewSimpleSet[string]()
	for _, v := range s {
		Q.Enqueue(v)
		seen.Add(v.id())
	}

	backTrack := map[string]string{}

	for !Q.IsEmpty() {
		currentNode := Q.Dequeue()

		if currentNode == destination {
			break
		}
		neighbors := getNeighbors(graph, currentNode)

		for _, v := range neighbors {
			if !seen.Has(v.id()) {
				seen.Add(v.id())
				backTrack[v.id()] = currentNode.id()
				Q.Enqueue(v)
			}
		}
	}
	return count(backTrack, destination.id())
}

func getMinValue(arr []int) (m int) {
	for i, e := range arr {
		if i == 0 || e < m && e != 0 {
			m = e
		}
	}
	return m
}

func parse2(raw string) (matrix [][]*Point, start []*Point, dest *Point) {
	lines := strings.Split(string(raw), "\n")
	matrix = make([][]*Point, len(lines))
	for i := range matrix {
		matrix[i] = make([]*Point, len(lines[0]))
		rows := []*Point{}
		for j, c := range lines[i] {
			if c == 'S' || c == 'a' {
				start = append(start, createPoint(i, j, c))
			}
			if c == 'E' {
				dest = createPoint(i, j, c)
			}
			rows = append(rows, createPoint(i, j, c))
		}
		matrix[i] = rows
	}

	return matrix, start, dest
}

type Point struct {
	i, j, v int
}

func (p *Point) id() string {
	return fmt.Sprintf("(%d,%d)", p.i, p.j)
}

func newPoint(i, j, v int) *Point {
	return &Point{i: i, j: j, v: v}
}

func createPoint(i, j int, r rune) *Point {
	switch r {
	case 'S':
		return newPoint(i, j, 0)
	case 'E':
		return newPoint(i, j, int('z')-int('a'))
	default:
		return newPoint(i, j, int(r)-int('a'))
	}
}
