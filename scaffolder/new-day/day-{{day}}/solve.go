package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./example.txt")
	check(err)

	fmt.Println("Part 1")
	fmt.Println(part1(data))

	fmt.Println("Part 2")
	fmt.Println(part2(data))

}

func part1(raw []byte) int {

	var input = parse(string(raw))
	fmt.Println(input)
	return 1

}

func part2(raw []byte) int {
	var input = parse(string(raw))
	fmt.Println(input)

	return 1

}

func parse(raw string) []string {
	lines := strings.Split(string(raw), "\n")

	return lines
}
