package main

import (
	"fmt"
	"io/ioutil"
	//"sort"
	"strings"
	//"regexp"
	//"strconv"
)

func main() {
	input, err := ioutil.ReadFile("adv06.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")

	// fmt.Println(lines)

	part1 := 0
	part2 := 0

	group := make(map[rune]interface{})
	groupI := make(map[rune]int)
	groupPersons := 0
	for i := range lines {
		if len(lines[i]) == 0 {
			part1 += len(group)
			group = make(map[rune]interface{})

			for _, v := range groupI {
				if v == groupPersons {
					part2++
				}
			}
			groupI = make(map[rune]int)
			groupPersons = 0
			continue
		}
		groupPersons++

		for j := range lines[i] {
			char := rune(lines[i][j])
			group[char] = nil
			if v, exists := groupI[char]; exists {
				groupI[char] = v + 1
			} else {
				groupI[char] = 1
			}
		}
	}
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
