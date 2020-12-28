package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	"strings"
)

import "github.com/und3f/advent-of-code-2020/grid3d"

func newGridFromString(lines []string) (*grid3d.Grid3D, *grid3d.HyperCube) {
	grid := grid3d.NewGrid3D()
	cube := grid3d.NewHyperCube()

	offset := (len(lines) - 1) / 2

	for y, row := range lines {
		for x, cell := range row {
			if cell == '#' {
				grid.Set(grid3d.Coord{0, y - offset, x - offset}, true)
				cube.Set(grid3d.HyperCoord{0, 0, y - offset, x - offset}, true)
			}
		}
	}

	return grid, cube
}

const Cycles = 6

func main() {
	reportLine, err := ioutil.ReadFile("adv17.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(reportLine)), "\n")

	gridOrigin, cubeOrigin := newGridFromString(lines)

	var partOne, partTwo int

	grid := gridOrigin
	for i := 0; i < Cycles; i++ {
		grid = grid.Cycle()
	}
	grid.Print()
	partOne = grid.CalculateActive()

	cube := cubeOrigin
	for i := 0; i < Cycles; i++ {
		cube = cube.Cycle()
	}
	cube.Print()
	partTwo = cube.CalculateActive()

	fmt.Println("Part One:", partOne)
	fmt.Println("Part Two:", partTwo)
}
