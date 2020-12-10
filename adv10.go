package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	//"regexp"
	"strconv"
)

func calcVariants(start, offset int, numbers []int, cache []int) int {
	//fmt.Printf("%d -> %v\n", start, numbers)
	if len(numbers) == 0 {
		return 1
	}

	variants := 0
	for i := offset; i < offset+3 && i < len(numbers); i++ {
		if start+3 < numbers[i] {
			break
		}

		//fmt.Println(cache)
		if i+1 == len(numbers) {
			variants++
		} else {
			localV := 0

			if cached := cache[i+1]; cached != 0 {
				localV = cached
			} else {
				localV = calcVariants(numbers[i], i+1, numbers, cache)
			}
			//fmt.Println(cache, localV)
			variants += localV
		}
	}

	cache[offset] = variants

	return variants
}

func main() {
	part1 := 0
	part2 := 0

	input, err := ioutil.ReadFile("adv10.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]

	numbers := make([]int, len(lines))
	for i := range numbers {
		v, err := strconv.Atoi(lines[i])
		if err != nil {
			panic("Wrong value ")
		}
		numbers[i] = v
	}
	sort.Ints(numbers)

	diffs := make([]int, 3)
	adapter := 0

	for i := range numbers {
		diff := numbers[i] - adapter
		adapter = numbers[i]
		diffs[diff-1]++
	}
	diffs[2]++
	part1 = diffs[0] * diffs[2]

	cache := make([]int, len(numbers))
	part2 = calcVariants(0, 0, numbers, cache)

	fmt.Println(diffs)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
