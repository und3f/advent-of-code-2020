package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	"strings"
)

import "github.com/und3f/advent-of-code-2020/grid3d"

func newGridFromString(lines []string) *grid3d.Grid3D {
	grid := grid3d.NewGrid3D()
	offset := (len(lines) - 1) / 2

	for y, row := range lines {
		for x, cell := range row {
			grid.Set(grid3d.Coord{0, y - offset, x - offset}, cell == '#')
		}
	}

	return grid
}

const Cycles = 6

func main() {
	reportLine, err := ioutil.ReadFile("adv17.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(reportLine)), "\n")

	gridOrigin := newGridFromString(lines)

	var partOne, partTwo int

	grid := gridOrigin
	for i := 0; i < Cycles; i++ {
		grid = grid.Cycle()
	}
	grid.Print()
	partOne = grid.CalculateActive()

	fmt.Println("Part One:", partOne)
	fmt.Println("Part Two:", partTwo)
}
