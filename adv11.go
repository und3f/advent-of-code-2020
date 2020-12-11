package main

import (
	"fmt"
	"io/ioutil"
	//"sort"
	"strings"
	//"regexp"
	//"strconv"
)

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func round(seats [][]rune) ([][]rune, bool) {

	result := make([][]rune, len(seats))
	for i := range seats {
		result[i] = make([]rune, len(seats[i]))
		for j := range seats[i] {
			occupied := 0
			for k := range directions {
				y := i + directions[k][0]
				x := j + directions[k][1]
				if y >= 0 && y < len(seats) && x >= 0 && x < len(seats[y]) {
					/// fmt.Println(i, j, y, x)
					if seats[y][x] == '#' {
						occupied++
					}
				}
			}

			result[i][j] = seats[i][j]
			if seats[i][j] == 'L' {
				if occupied == 0 {
					result[i][j] = '#'
				}
			} else if seats[i][j] == '#' {
				if occupied >= 4 {
					result[i][j] = 'L'
				}
			}
		}
	}

	for i := range result {
		for j := range result[i] {
			if seats[i][j] != result[i][j] {
				return result, false
			}
		}
	}

	return result, true
}

func round2(seats [][]rune) ([][]rune, bool) {

	result := make([][]rune, len(seats))
	//fmt.Println("here")
	for i := range seats {
		result[i] = make([]rune, len(seats[i]))
		for j := range seats[i] {
			occupied := 0
			for k := range directions {
				y := i + directions[k][0]
				x := j + directions[k][1]
				for y >= 0 && y < len(seats) && x >= 0 && x < len(seats[y]) {
					//fmt.Println(i, j, y, x, string(seats[y][x]))
					if seats[y][x] == '.' {
						y = y + directions[k][0]
						x = x + directions[k][1]
						continue
					}

					if seats[y][x] == '#' {
						occupied++
					}
					break
				}
			}

			result[i][j] = seats[i][j]
			if seats[i][j] == 'L' {
				if occupied == 0 {
					result[i][j] = '#'
				}
			} else if seats[i][j] == '#' {
				if occupied >= 5 {
					result[i][j] = 'L'
				}
			}
		}
	}

	for i := range result {
		for j := range result[i] {
			if seats[i][j] != result[i][j] {
				return result, false
			}
		}
	}

	return result, true
}

func printSeats(seats [][]rune) {
	for i := range seats {
		for j := range seats[i] {
			fmt.Printf("%s", string(seats[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	part1 := 0
	part2 := 0

	input, err := ioutil.ReadFile("adv11.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var seats, seats2 [][]rune
	for i := range lines {
		seats = append(seats, []rune(lines[i]))
		seats2 = append(seats2, []rune(lines[i]))
	}
	stabilized := false

	for !stabilized {
		seats, stabilized = round(seats)
		//printSeats(seats)
	}

	for i := range seats {
		for j := range seats[i] {
			if seats[i][j] == '#' {
				part1++
			}
		}
	}

	stabilized = false
	for !stabilized {
		seats2, stabilized = round2(seats2)
		//printSeats(seats2)
	}

	for i := range seats2 {
		for j := range seats2[i] {
			if seats2[i][j] == '#' {
				part2++
			}
		}
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
