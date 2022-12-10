package main

import (
	"fmt"

	"github.com/galElmalah/aoc-2022/day-7/fileSystem"
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
	fs := parse(raw)
	sum := 0
	fs.Walk(func(t *fileSystem.FileSystemNode) {
		if t.Size <= 100000 {
			sum += t.Size
		}
	})

	return sum
}

func Part2(raw string) int {
	fs := parse(raw)

	const OS_MEM = 70000000
	const THRESHOLD = 30000000

	unusedSpace := OS_MEM - fs.Size()
	min := OS_MEM
	fs.Walk(func(node *fileSystem.FileSystemNode) {
		if unusedSpace+node.Size > THRESHOLD {
			if min > node.Size {
				min = node.Size
			}
		}
	})

	return min
}

func parse(raw string) *fileSystem.FileSystem {
	return fileSystem.NewFileSystem(token.Tokenize(raw))
}
