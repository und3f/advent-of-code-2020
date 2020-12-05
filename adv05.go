package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	//"regexp"
	//"strconv"
)

func CalcBSP(bsp string, region int) int {
	low := 0
	high := region - 1

	// fmt.Println(bsp)
	for i := range bsp {
		letter := bsp[i]
		median := low + (high-low)/2
		// fmt.Println(low, high, median)
		if letter == 'F' || letter == 'L' {
			high = median
		} else if letter == 'B' || letter == 'R' {
			low = median + 1
		} else {
			panic("Error")
		}
	}

	if low != high {
		panic("Error")
	}
	return low
}

func main() {
	input, err := ioutil.ReadFile("adv05.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	highSeatId := 0
	var seats []int
	for i := range lines {
		bsp := lines[i]
		if len(bsp) < 8 {
			continue
		}
		seatId := CalcBSP(bsp[:7], 128)*8 + CalcBSP(bsp[7:], 8)
		if seatId > highSeatId {
			highSeatId = seatId
		}
		seats = append(seats, seatId)
	}

	fmt.Println("Part1:", highSeatId)

	sort.Ints(seats)
	for i := 0; i < len(seats)-1; i++ {
		if seats[i]+1 < seats[i+1] {
			fmt.Println("Part2:", seats[i]+1)
			break
		}
	}
}
