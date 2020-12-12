package main

import (
	"fmt"
	"io/ioutil"
	//"sort"
	"strings"
	//"regexp"
	"strconv"
)

type Position struct {
	east, north int
}

var directions = []Position{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func moveShipWaypoint(instructions []string) int {
	waypoint := Position{10, 1}
	pos := Position{0, 0}

	for _, instr := range instructions {
		command := instr[0]
		value, err := strconv.Atoi(instr[1:])
		if err != nil {
			panic(err)
		}

		switch command {
		case 'N':
			waypoint.north += directions[0].north * value
		case 'S':
			waypoint.north += directions[2].north * value
		case 'E':
			waypoint.east += directions[1].east * value
		case 'W':
			waypoint.east += directions[3].east * value
		case 'L':
			steps := value / 90
			for j := 0; j < steps; j++ {
				n := waypoint.north
				waypoint.north = waypoint.east
				waypoint.east = -n
			}
		case 'R':
			steps := value / 90
			for j := 0; j < steps; j++ {
				n := waypoint.north
				waypoint.north = -waypoint.east
				waypoint.east = n
			}
		case 'F':
			pos.east += waypoint.east * value
			pos.north += waypoint.north * value
		default:
			panic("Unknown instruction")
		}

		fmt.Println(instr, pos)
	}

	return Abs(pos.east) + Abs(pos.north)
}

func moveShip(instructions []string) int {
	pos := Position{0, 0}
	direction := 1

	for _, instr := range instructions {
		command := instr[0]
		value, err := strconv.Atoi(instr[1:])
		if err != nil {
			panic(err)
		}

		switch command {
		case 'N':
			pos.north += directions[0].north * value
		case 'S':
			pos.north += directions[2].north * value
		case 'E':
			pos.east += directions[1].east * value
		case 'W':
			pos.east += directions[3].east * value
		case 'L':
			direction = ((direction - value/90) + len(directions)) % len(directions)
		case 'R':
			direction = ((direction + value/90) + len(directions)) % len(directions)
		case 'F':
			d := directions[direction]
			pos.east += d.east * value
			pos.north += d.north * value
		default:
			panic("Unknown instruction")
		}

		fmt.Println(instr, pos)
	}

	return Abs(pos.east) + Abs(pos.north)
}

func main() {
	part1 := 0
	part2 := 0

	input, err := ioutil.ReadFile("adv12.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Println(lines)
	part1 = moveShip(lines)
	part2 = moveShipWaypoint(lines)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
