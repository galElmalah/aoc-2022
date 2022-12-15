package main

import (
	"encoding/json"
	"fmt"
	"sort"
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

func compare(left, right any) int {
	leftArr, isLeftArray := left.([]any)
	rightArr, isRightArray := right.([]any)
	if !isLeftArray && !isRightArray {
		return int(left.(float64) - right.(float64))
	} else if !isLeftArray {
		return compare([]any{left}, right)
	} else if !isRightArray {
		return compare(left, []any{right})
	}

	for i := 0; i < len(leftArr) && i < len(rightArr); i++ {
		res := compare(leftArr[i], rightArr[i])
		if res != 0 {
			return res
		}
	}

	return len(leftArr) - len(rightArr)

}

func Part1(raw string) int {
	packets := parse(raw)
	sum := 0

	for i := 0; i < len(packets)-1; i += 2 {
		rs := compare(packets[i], packets[i+1])
		if rs <= 0 {
			sum += int(i/2) + 1
		}
	}

	return sum
}

func Part2(raw string) int {
	packets := parse(raw)
	div1, div2 := []any{[]any{float64(2)}}, []any{[]any{float64(6)}}
	packets = append(packets, div1, div2)
	sort.Slice(packets, func(i, j int) bool { return compare(packets[i], packets[j]) < 0 })
	res := 1
	for i, p := range packets {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			res *= i + 1
		}
	}

	return res
}

func parse(raw string) []any {
	var packets []any
	lines := strings.Split(string(raw), "\n\n")
	for _, l := range lines {
		var left, right any
		parts := strings.Split(l, "\n")
		json.Unmarshal([]byte(parts[0]), &left)
		json.Unmarshal([]byte(parts[1]), &right)

		packets = append(packets, left, right)
	}

	return packets
}
