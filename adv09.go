package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	//"regexp"
	"strconv"
)

const preambleLength = 25

func main() {
	part1 := 0
	part2 := 0

	input, err := ioutil.ReadFile("adv09.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]
	numbers := make([]int, len(lines))
	for i := range lines {
		v, err := strconv.Atoi(lines[i])
		if err != nil {
			panic("Wrong number")
		}
		numbers[i] = v
	}
	fmt.Println(lines)

	for i := preambleLength; i < len(lines); i++ {
		preamble := numbers[i-preambleLength : i]
		found := false
		for j := range preamble {
			for k := j + 1; k < len(preamble); k++ {
				if preamble[j]+preamble[k] == numbers[i] {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		fmt.Println(preamble, lines[i])
		if !found {
			part1 = numbers[i]
			break
		}
	}

	for i := range numbers {
		sum := numbers[i]
		j := i + 1
		for ; j < len(numbers); j++ {
			sum += numbers[j]
			if sum >= part1 {
				break
			}
		}
		if sum == part1 {
			contiguous := numbers[i : j+1]
			sort.Ints(contiguous)
			part2 = contiguous[0] + contiguous[len(contiguous)-1]
			break
		}
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
