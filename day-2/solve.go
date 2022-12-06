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

const WINNING_POINTS = 6
const TIE_POINTS = 3

func main() {
	data, err := os.ReadFile("./input.txt")
	check(err)

	fmt.Println("Part 1")
	fmt.Println(part1(data))

	fmt.Println("Part 2")
	fmt.Println(part2(data))

}

func part1(raw []byte) int {

	pairs := parse(string(raw))

	// X Rock, Y Paper, Z Scissors
	scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	win := map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}

	tie := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	score := 0
	for _, pair := range pairs {
		his := pair[0]
		my := pair[1]
		score += scores[my]
		if win[my] == his {
			score += WINNING_POINTS
		}

		if tie[my] == his {
			score += TIE_POINTS
		}

	}
	return score

}

func part2(raw []byte) int {

	var pairs = parse(string(raw))

	// X Lose, Y Tie, Z Win
	scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	win := map[string]string{
		"C": "X",
		"A": "Y",
		"B": "Z",
	}

	tie := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	lose := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}
	score := 0

	for _, pair := range pairs {
		hisMove := pair[0]
		myMove := pair[1]

		// we lose
		if myMove == "X" {
			score += scores[lose[hisMove]]
		}
		// we end in a tie
		if myMove == "Y" {
			score += TIE_POINTS
			score += scores[tie[hisMove]]
		}
		// we win
		if myMove == "Z" {
			score += WINNING_POINTS
			score += scores[win[hisMove]]
		}
	}
	return score
}

func parse(raw string) [][]string {
	chunks := strings.Split(string(raw), "\n")
	pairs := make([][]string, len(chunks))
	for i := range pairs {
		pairs[i] = strings.Split(chunks[i], " ")
	}

	return pairs
}
