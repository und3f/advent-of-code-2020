package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	"strings"
)

var slopes = [][]int{
	[]int{1, 1},
	[]int{1, 3},
	[]int{1, 5},
	[]int{1, 7},
	[]int{2, 1},
}

func main() {
	reportLine, err := ioutil.ReadFile("adv03.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(reportLine), "\n")
	answers := make([]int, len(slopes))
	for i := range slopes {
		x := 0
		trees := 0
		for y := 0; y < len(lines); y += slopes[i][0] {
			if len(lines[y]) == 0 {
				break
			}
			//fmt.Println(y, x, string(lines[y][x]))
			if lines[y][x] == '#' {
				trees++
			}
			x = (x + slopes[i][1]) % len(lines[0])
		}
		answers[i] = trees
	}
	partTwo := 1
	for i := range answers {
		partTwo *= answers[i]
	}
	fmt.Println(answers)
	fmt.Println("Part One:", answers[1])
	fmt.Println("Part Two:", partTwo)
}
