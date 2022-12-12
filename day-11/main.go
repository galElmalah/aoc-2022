package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/galElmalah/aoc-2022/ds/set"
	"github.com/galElmalah/aoc-2022/util"
)

func main() {
	data := util.ReadFile("./input.txt")

	fmt.Println("Part 1")
	fmt.Println(Part1(data))

	// fmt.Println("Part 2")
	// fmt.Println(Part2(data))

}
func remove(s []int, toDelete *set.SimpleSet[int]) []int {
	newarr := []int{}
	for i, v := range s {
		if !toDelete.Has(i) {
			newarr = append(newarr, v)
		}
	}
	return newarr
}
func Part1(raw string) int {
	monkeys := parse(raw)
	count := map[int]int{}
	var base int = 1
	for _, m := range monkeys {
		base *= m.test.divisor
	}
	fmt.Println(base)

	for k := 0; k < 10000; k++ {

		for i, m := range monkeys {
			toDelete := set.NewSimpleSet[int]()

			for ii, item := range m.items {
				count[m.id]++
				newVal := item
				var r, l int
				if l = ParseInt(m.op.l); m.op.l == "old" {
					l = item
				}
				if r = ParseInt(m.op.r); m.op.r == "old" {

					r = item
				}
				// fmt.Println("before::", newVal)
				if m.op.operator == "+" {
					newVal = r + l
				} else {
					newVal = r * l
				}
				newVal = newVal % base
				// fmt.Println("after::", newVal)
				// fmt.Println("Module::", newVal, m.test.divisor, mod)
				if newVal%m.test.divisor == 0 {
					monkeys[m.test.t].items = append(monkeys[m.test.t].items, newVal)
					// fmt.Printf("test t:%+v\n%v\n", monkeys[0].items, newVal.Mod(newVal, m.test.divisor))

				} else {
					// fmt.Printf("test f:%+v\n%v\n", monkeys[0].items, newVal.Mod(newVal, m.test.divisor))

					monkeys[m.test.f].items = append(monkeys[m.test.f].items, newVal)
				}
				toDelete.Add(ii)
			}

			monkeys[i].items = remove(m.items, toDelete)

		}
		// fmt.Printf("%+v\n", monkeys[0].items)
		// fmt.Printf("%+v\n", monkeys[1].items)
		// fmt.Printf("%+v\n", monkeys[2].items)
	}

	arr := []int{}
	for _, v := range count {
		arr = append(arr, v)
	}
	sort.Ints(arr)
	fmt.Println(count)
	return arr[len(arr)-1] * arr[len(arr)-2]
}

func Part2(raw string) int {
	input := parse(raw)
	fmt.Println(input)

	return -1
}

type Op struct {
	operator string
	l        string
	r        string
}

type Test struct {
	divisor int
	t       int // if true throw to monkey NO
	f       int
}

type Monkey struct {
	items     []int
	op        Op
	test      Test
	id        int
	Inspected int
}

func ParseInt(s string) int {
	val, _ := strconv.Atoi(s)
	return int(val)
}

func parseItems(str string) (items []int) {
	for _, v := range strings.Split(str, ", ") {
		items = append(items, ParseInt(v))
	}
	return items
}

func parse(raw string) []Monkey {
	chunks := strings.Split(string(raw), "\n\n")
	monkeys := []Monkey{}
	r := regexp.MustCompile(`Starting items: (.*)
  Operation: new = (.*) (\*|\+) (.*)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)`)
	for i, m := range chunks {
		result := r.FindStringSubmatch(m)[1:]
		fmt.Println(result)
		monkeys = append(monkeys, Monkey{
			id:    i,
			items: parseItems(result[0]),
			op:    Op{l: result[1], operator: result[2], r: result[3]},
			test:  Test{divisor: ParseInt(result[4]), t: util.ParseInt(result[5]), f: util.ParseInt(result[6])},
		})
	}
	return monkeys
}
