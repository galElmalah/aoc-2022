package main

import (
	"fmt"

	"github.com/galElmalah/aoc-2022/day-7/fileTree"
	"github.com/galElmalah/aoc-2022/day-7/token"
	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	fmt.Println("Part 2")
	fmt.Println(Part2(data))

}

func Part1(raw string) int {
	tree := parse(raw)
	sum := 0
	count := 0
	tree.Walk(func(t *fileTree.FileTree) {
		count++
		if t.Size <= 100000 {
			sum += t.Size
		}
	})

	return sum
}

func Part2(raw string) int {
	tree := parse(raw)

	const OS_MEM = 70000000
	const THRESHOLD = 30000000

	unusedSpace := OS_MEM - tree.Size
	min := OS_MEM
	tree.Walk(func(t *fileTree.FileTree) {
		if unusedSpace+t.Size > THRESHOLD {
			if min > t.Size {
				min = t.Size
			}
		}
	})

	return min
}

func parse(raw string) *fileTree.FileTree {
	return fileTree.CreateFileTree(token.Tokenize(raw))
}
