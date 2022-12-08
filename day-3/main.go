package main

import (
	"fmt"
	"strings"

	"github.com/galElmalah/aoc-2022/ds/set"
	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	fmt.Println("Part 1")
	fmt.Println(Part1())

	fmt.Println("Part 2")
	fmt.Println(Part2())

}

func Part1() int {
	groups := parse()
	sum := 0
	for _, group := range groups {
		s1 := group[0]
		s2 := group[1]
		for _, k := range s1.Values() {
			if s2.Has(k) {
				sum += calcPriority(k)
			}
		}
	}

	return sum
}

func Part2() int {
	groups := parse2()

	sum := 0
	for _, group := range groups {
		s1 := group[0]
		s2 := group[1]
		s3 := group[2]
		for _, k := range s1.Values() {
			if s2.Has(k) && s3.Has(k) {
				sum += calcPriority(k)
			}
		}
	}

	return sum
}

func parse() [][]set.Set[rune] {
	data := util.ReadFile("./input.txt")
	lines := strings.Split(string(data), "\n")

	rucksacks := [][]set.Set[rune]{}
	for _, line := range lines {
		c1 := makeSet(line[len(line)/2:])
		c2 := makeSet(line[:len(line)/2])
		w := []set.Set[rune]{c1, c2}
		rucksacks = append(rucksacks, w)
	}

	return rucksacks
}

func parse2() [][]set.Set[rune] {
	data := util.ReadFile("./input.txt")
	rows := strings.Split(string(data), "\n")
	groups := chunkInto(rows, 3)
	rucksacks := [][]set.Set[rune]{}
	for _, chunk := range groups {
		w := []set.Set[rune]{}
		for _, c := range chunk {
			w = append(w, makeSet(c))
		}
		rucksacks = append(rucksacks, w)
	}
	return rucksacks
}

func calcPriority(c rune) int {
	val := int(c)
	if c >= 'a' && c <= 'z' {
		return val - 'a' + 1
	} else {
		return val - 'A' + 27
	}
}

func chunkInto(s []string, size int) [][]string {
	results := [][]string{}
	for i := 0; i < len(s); i += size {
		results = append(results, s[i:i+size])
	}
	return results
}

func makeSet(str string) set.Set[rune] {
	set := set.NewSet[rune]()
	for _, c := range str {
		set.Add(c)
	}
	return *set
}
