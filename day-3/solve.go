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
	fmt.Println("Part 1")
	fmt.Println(pt1())

	fmt.Println("Part 2")
	fmt.Println(pt2())
	// fmt.Println(result[len(result)-1] + result[len(result)-2] + result[len(result)-3])

}

func pt1() int {
	groups := parse()
	results := []int{}
	for _, group := range groups {
		s1 := group[0]
		s2 := group[1]
		for k, _ := range s1 {
			if s2[k] {
				results = append(results, calcPriority(k))
			}
		}
	}

	return sum(results)
}

func pt2() int {
	groups := parse2()

	results := []int{}
	for _, group := range groups {
		s1 := group[0]
		s2 := group[1]
		s3 := group[2]
		for k, _ := range s1 {
			if s2[k] && s3[k] {
				results = append(results, calcPriority(k))
			}
		}
	}

	return sum(results)
}

func makeSet(chars string) map[rune]bool {
	set := map[rune]bool{}
	for _, c := range chars {
		set[c] = true
	}
	return set
}

func parse() [][]map[rune]bool {
	data, _ := os.ReadFile("./input.txt")

	chunks := strings.Split(string(data), "\n")
	groups := [][]map[rune]bool{}
	for _, chunk := range chunks {
		c1 := makeSet(chunk[0 : len(chunk)/2])
		c2 := makeSet(chunk[len(chunk)/2:])
		w := []map[rune]bool{c1, c2}
		groups = append(groups, w)
	}

	return groups
}

func parse2() [][]map[rune]bool {
	data, _ := os.ReadFile("./input.txt")
	rows := strings.Split(string(data), "\n")
	chunks := chunkInto(rows, 3)
	groups := [][]map[rune]bool{}
	for _, chunk := range chunks {
		w := []map[rune]bool{}
		for _, c := range chunk {
			w = append(w, makeSet(c))
		}
		groups = append(groups, w)
	}
	return groups
}

func calcPriority(c rune) int {
	if c >= 97 && c <= 122 {
		return int(c) - 97 + 1
	} else {
		return int(c) - 65 + 27
	}
}

func chunkInto(s []string, size int) [][]string {
	results := [][]string{}
	for i := 0; i < len(s); i += size {
		results = append(results, s[i:i+size])
	}
	return results
}

func sum(arr []int) int {
	res := 0
	for _, num := range arr {
		res += num
	}
	return res
}
