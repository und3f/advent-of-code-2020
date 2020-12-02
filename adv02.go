package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	reportLine, err := ioutil.ReadFile("adv02.txt")
	if err != nil {
		panic(err)
	}
	passwordLines := strings.Split(string(reportLine), "\n")

	correct := 0
	correctPart2 := 0

	for i := range passwordLines {
		line := passwordLines[i]
		if len(line) == 0 {
			continue
		}

		data := strings.Split(line, ": ")
		password := data[1]
		policy := strings.Split(data[0], " ")
		letter := policy[1][0]
		minmax := strings.Split(policy[0], "-")
		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])

		occurances := 0
		for j := range password {
			if password[j] == letter {
				occurances++
			}
		}
		fmt.Println(min, max, letter, password, occurances)
		if occurances >= min && occurances <= max {
			correct++
		}

		if (password[min-1] == letter) != (password[max-1] == letter) {
			correctPart2++
		}

	}
	fmt.Println("Part1:", correct)
	fmt.Println("Part2:", correctPart2)
}
