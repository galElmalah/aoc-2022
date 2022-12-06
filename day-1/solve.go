package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	chunks := parse()

	result := []int{}
	
	for _, chunk := range chunks {
		result = append(result, sumChunk(chunk))
	}

	sort.Ints(result)

	fmt.Println("Part 1")
	fmt.Println(result[len(result)-1])

	fmt.Println("Part 2")
	fmt.Println(result[len(result)-1] + result[len(result)-2] + result[len(result)-3])

}

func sumChunk(chunk string) int {
	sum := 0
	for _, num := range strings.Split(chunk, "\n") {
		v, _ := strconv.Atoi(num)
		sum += v
	}
	return sum
}

func parse() []string {
	data, _ := os.ReadFile("./input.txt")
	chunks := strings.Split(string(data), "\n\n")
	return chunks
}
