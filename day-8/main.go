package main

import (
	"fmt"
	"strings"

	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	fmt.Println("Part 2")
	fmt.Println(Part2(data))

}

func checkLeft(mat [][]int, i, j int) bool {
	for k := j - 1; k >= 0; k-- {
		if mat[i][k] >= mat[i][j] {
			return false
		}
	}
	return true
}

func checkRight(mat [][]int, i, j int) bool {
	for k := j + 1; k < len(mat[0]); k++ {
		if mat[i][k] >= mat[i][j] {
			return false
		}
	}
	return true
}

func checkUp(mat [][]int, i, j int) bool {
	for k := i - 1; k >= 0; k-- {
		if mat[k][j] >= mat[i][j] {
			return false
		}
	}
	return true
}

func checkDown(mat [][]int, i, j int) bool {
	for k := i + 1; k < len(mat[0]); k++ {
		if mat[k][j] >= mat[i][j] {
			return false
		}
	}
	return true
}

func Part1(raw string) int {
	mat := parse(raw)
	visibleTrees := (len(mat)-2)*2 + len(mat[0])*2

	for i, row := range mat {
		if i == 0 || i == len(mat)-1 {
			continue
		}

		for j, _ := range row {
			if j == 0 || j == len(row)-1 {
				continue
			}
			if checkLeft(mat, i, j) {
				visibleTrees++
				continue
			}
			if checkRight(mat, i, j) {
				visibleTrees++
				continue
			}
			if checkUp(mat, i, j) {
				visibleTrees++
				continue
			}
			if checkDown(mat, i, j) {
				visibleTrees++
				continue
			}

		}
	}

	return visibleTrees
}

func calcLeft(mat [][]int, i, j int) int {
	vis := 0

	for k := j - 1; k >= 0; k-- {

		if mat[i][k] >= mat[i][j] {
			return vis + 1
		}
		vis++
	}
	return vis
}

func calcRight(mat [][]int, i, j int) int {
	vis := 0

	for k := j + 1; k < len(mat[0]); k++ {
		if mat[i][k] >= mat[i][j] {
			return vis + 1
		}
		vis++
	}
	return vis
}

func calcUp(mat [][]int, i, j int) int {
	vis := 0

	for k := i - 1; k >= 0; k-- {
		if mat[k][j] >= mat[i][j] {
			return vis + 1
		}
		vis++
	}
	return vis
}

func calcDown(mat [][]int, i, j int) int {
	vis := 0
	for k := i + 1; k < len(mat[0]); k++ {
		if mat[k][j] >= mat[i][j] {
			return vis + 1
		}
		vis++
	}
	return vis
}

func Part2(raw string) int {
	mat := parse(raw)
	max := 0
	for i, row := range mat {
		for j, _ := range row {
			score := calcLeft(mat, i, j) * calcRight(mat, i, j) * calcUp(mat, i, j) * calcDown(mat, i, j)
			if score > max {
				max = score
			}
		}
	}

	return max
}

func parse(raw string) [][]int {
	rows := strings.Split(string(raw), "\n")
	mat := make([][]int, len(rows))
	for i, row := range rows {
		mat[i] = make([]int, len(row))
	}

	for i, row := range rows {
		for j, c := range row {
			mat[i][j] = util.ParseInt(string(c))
		}
	}

	return mat
}
