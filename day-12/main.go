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

	// fmt.Println("Part 1")
	// fmt.Println(Part1(data))

	fmt.Println("Part 2")
	fmt.Println(Part2(data))

}

func Part1(raw string) int {
	graph, start, dest := parse(raw)
	steps := BFS(graph, start, dest)
	// for _, v := range graph {
	// 	for _, v2 := range v {
	// 		fmt.Println(v2)

	// 	}
	// }
	// fmt.Println(graph)

	return steps
}

type Point struct {
	i, j, v int
}

func (p *Point) id() string {
	return fmt.Sprintf("(%d,%d)", p.i, p.j)
}

func newPoint(i, j, v int) *Point {
	return &Point{i, j, v}
}

func getNs(graph [][]*Point, sink *Point) []*Point {
	result := []*Point{}
	moves := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, move := range moves {
		di, dj := move[0]+sink.i, move[1]+sink.j
		if di == sink.i && dj == sink.j {
			continue
		}
		if di >= 0 && di < len(graph) && dj >= 0 && dj < len(graph[0]) {
			delta := graph[di][dj].v - sink.v
			if delta > 1 {
				continue
			}
			// fmt.Printf("yo %s %s", sink.id(), graph[di][dj].id())

			result = append(result, graph[di][dj])
		}

	}
	return result

}

func print(backTrack map[string]string, id string) {
	// fmt.Println(id)
	v, ok := backTrack[id]
	if ok {
		print(backTrack, v)
	}
}

func count(backTrack map[string]string, id string) int {
	v, ok := backTrack[id]
	if ok {
		return 1 + count(backTrack, v)
	}
	return 0
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
			// fmt.Println("found it!!!", currentNode)
			break
		}
		neighbors := getNs(graph, currentNode)
		// fmt.Printf("%s", currentNode.id())

		for _, v := range neighbors {
			if !seen.Has(v.id()) {
				seen.Add(v.id())
				backTrack[v.id()] = currentNode.id()
				Q.Enqueue(v)
				// fmt.Printf(" -> %s", v.id())
			}

		}
		// fmt.Printf("\n")

	}

	// for _, v := range graph {
	// 	fmt.Println()
	// 	for _, k := range v {
	// 		fmt.Printf(" %v %v  |", k.id(), k.v)
	// 	}
	// }

	return count(backTrack, destination.id())

}

func parse(raw string) (graph [][]*Point, start *Point, dest *Point) {
	lines := strings.Split(string(raw), "\n")
	graph = make([][]*Point, len(lines))
	for i := range graph {
		graph[i] = make([]*Point, len(lines[0]))
		points := []*Point{}
		for j, c := range lines[i] {
			if c == 'S' {
				start = newPoint(i, j, 0)
				points = append(points, start)
				continue
			}
			if c == 'E' {
				dest = newPoint(i, j, int('z')-int('a'))
				points = append(points, dest)
			} else {
				points = append(points, newPoint(i, j, int(c)-int('a')))
			}
		}
		graph[i] = points
	}

	return graph, start, dest

}

func BFS2(graph [][]*Point, s []*Point, destination *Point) int {
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
			// fmt.Println("found it!!!", currentNode)
			break
		}
		neighbors := getNs(graph, currentNode)

		for _, v := range neighbors {
			if !seen.Has(v.id()) {
				seen.Add(v.id())
				backTrack[v.id()] = currentNode.id()
				Q.Enqueue(v)
				// fmt.Printf(" -> %s", v.id())
			}

		}
		// fmt.Printf("\n")

	}

	// for _, v := range graph {
	// 	fmt.Println()
	// 	for _, k := range v {
	// 		fmt.Printf(" %v %v  |", k.id(), k.v)
	// 	}
	// }

	return count(backTrack, destination.id())

}
func Part2(raw string) int {
	graph, start, dest := parse2(raw)
	
	steps := BFS2(graph, start, dest)

	return steps
}

func parse2(raw string) (graph [][]*Point, start []*Point, dest *Point) {
	lines := strings.Split(string(raw), "\n")
	graph = make([][]*Point, len(lines))
	for i := range graph {
		graph[i] = make([]*Point, len(lines[0]))
		points := []*Point{}
		for j, c := range lines[i] {
			if c == 'S' {
				start = append(start, newPoint(i, j, 0))
				points = append(points, newPoint(i, j, 0))

				continue
			}
			if c == 'S' || c == 'a' {
				start = append(start, newPoint(i, j, 0))
			}
			if c == 'E' {
				dest = newPoint(i, j, int('z')-int('a'))
				points = append(points, dest)
			} else {
				points = append(points, newPoint(i, j, int(c)-int('a')))
			}
		}
		graph[i] = points
	}

	return graph, start, dest

}
