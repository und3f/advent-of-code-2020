package main

import (
	"fmt"
	"io/ioutil"
	//"sort"
	"strings"
	//"regexp"
	"github.com/deanveloper/modmath/bigmod"
	"math"
	"math/big"
	"strconv"
)

func calcDeparture(time int, busses []int) int {
	minTime := int(^uint(0) >> 1)
	var busid int
	for i := range busses {
		time := int(math.Ceil(float64(time)/float64(busses[i]))) * busses[i]
		if minTime > time {
			minTime = time
			busid = busses[i]
		}
	}
	return busid * (minTime - time)
}

func calcTimestamp(aSchedule []int) *big.Int {
	var crts []bigmod.CrtEntry

	for i, v := range aSchedule {
		if v != 0 {
			a := int64(v) - int64(i)
			crts = append(crts, bigmod.CrtEntry{A: big.NewInt(int64(a)), N: big.NewInt(int64(v))})
		}
	}

	return bigmod.SolveCrtMany(crts)
}

func main() {
	input, err := ioutil.ReadFile("adv13.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	time, _ := strconv.Atoi(lines[0])
	busses := strings.Split(lines[1], ",")
	var bussesIds []int
	var schedule []int
	for i := range busses {
		id, err := strconv.Atoi(busses[i])
		if err == nil {
			bussesIds = append(bussesIds, id)
			schedule = append(schedule, id)
		} else {
			schedule = append(schedule, 0)
		}
	}
	part1 := calcDeparture(time, bussesIds)
	part2 := calcTimestamp(schedule)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
