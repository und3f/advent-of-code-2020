package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	"bytes"
	"strings"
)

import "github.com/und3f/advent-of-code-2020/grid3d"

type Cube struct {
	dimension [][][]bool
	wx, wy    int
	zOffset   int
}

func newCube(lines []string) *Cube {
	cube := new(Cube)

	cube.wy = len(lines)
	cube.wx = len(lines[0])

	cube.zOffset = 0
	cube.dimension = make([][][]bool, 1)
	cube.dimension[0] = make([][]bool, cube.wy)
	for y, row := range lines {
		cube.dimension[0][y] = make([]bool, cube.wx)
		for x, cell := range row {
			cube.dimension[0][y][x] = cell == '#'
		}
	}

	return cube
}

func (cube *Cube) Print() {
	for z, grid := range cube.dimension {
		fmt.Printf("z=%d\n", z+cube.zOffset)
		var buffer bytes.Buffer
		for _, row := range grid {
			for _, cell := range row {
				symbol := '.'
				if cell {
					symbol = '#'
				}
				buffer.WriteRune(symbol)
			}
			buffer.WriteRune('\n')
		}
		fmt.Println(buffer.String())
	}
}

func (cube *Cube) Cycle() {
	// zOffset := cube.zOffset + 1
	dimension := make([][][]bool, len(cube.dimension)+2)

	for z := range dimension {
		dimension[z] = make([][]bool, cube.wy)
		for y := 0; y < cube.wy; y++ {
			for x := 0; x < cube.wx; x++ {
				dimension[z][y] = make([]bool, cube.wx)
				zOld := z - 1
				activeNeighbors := 0
				for dz := -1; dz <= 1; dz++ {
					for dy := -1; dy <= 1; dy++ {
						for dx := -1; dx <= 1; dx++ {
							z2 := zOld + dz
							x2 := x + dx
							y2 := y + dy
							if (z2 < 0 || z2 > len(cube.dimension)-1) || (y2 < 0 || y2 > cube.wy-1) || (x2 < 0 || x2 > cube.wx-1) {
								continue
							}
							fmt.Println(z, y, x, z2, y2, x2)
							if cube.dimension[z2][y2][x2] {
								activeNeighbors++
							}
						}
					}
				}

				isCellActive := false

				if zOld > 0 && zOld < len(cube.dimension) {
					isCellActive = cube.dimension[zOld][y][x]
				}

				newCellState := false
				if isCellActive {
					if activeNeighbors == 2 || activeNeighbors == 3 {
						newCellState = true
					}
				} else {
					if activeNeighbors == 3 {
						newCellState = true
					}
				}

				dimension[z][y][x] = newCellState
			}
		}
	}

	cube.dimension = dimension
}

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

	grid := newGridFromString(lines)

	var partOne, partTwo int
	grid.Print()
	fmt.Println(grid.CalcNeighbors(grid3d.Coord{-1, 0, 0}))

	for i := 0; i < Cycles; i++ {
		grid = grid.Cycle()
	}
	grid.Print()
	partOne = grid.CalculateActive()

	fmt.Println("Part One:", partOne)
	fmt.Println("Part Two:", partTwo)
}
