package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	// "regexp"
	//"sort"
	"strings"
)

type Floor struct {
	tiles [][]bool
}

func (floor *Floor) incSize() {
	for i, row := range floor.tiles {
		floor.tiles[i] = append([]bool{true}, append(row, true)...)
	}
	newLength := len(floor.tiles) + 2
	emptyRow1 := make([]bool, newLength)
	emptyRow2 := make([]bool, newLength)
	for i := range emptyRow1 {
		emptyRow1[i] = true
	}
	copy(emptyRow2, emptyRow1)
	floor.tiles = append([][]bool{emptyRow1}, append(floor.tiles, emptyRow2)...)
}

func newFloor() *Floor {
	floor := new(Floor)
	floor.tiles = [][]bool{{true}}

	return floor
}

func (floor Floor) String() (out string) {
	for i, row := range floor.tiles {
		if i != 0 {
			out += "\n"
		}

		if i%2 != 1 {
			out += "  "
		}
		for j, state := range row {
			if j > 0 {
				out += "  "
			}
			if state {
				out += " ⬡"
			} else {
				out += " ⬢"
			}
		}
	}
	return out
}

func (floor *Floor) flip(y, x int) {
	center := (len(floor.tiles)) / 2
	for {
		if !(x < -center || x >= center+1 || y < -center || y >= center+1) {
			break
		}
		floor.incSize()
		center = (len(floor.tiles)) / 2
	}
	floor.tiles[y+center][x+center] = !floor.tiles[y+center][x+center]
}

func (floor Floor) getTile(y, x int) bool {
	center := (len(floor.tiles)) / 2
	if x < -center || x >= center+1 || y < -center || y >= center+1 {
		return true
	}

	return floor.tiles[y+center][x+center]
}

func (floor Floor) countBlack() uint {
	var count uint
	for _, row := range floor.tiles {
		for _, cell := range row {
			if !cell {
				count++
			}
		}
	}
	return count
}

func (floor Floor) countAdjancent(y, x int, color bool) uint {
	adjancent := uint(0)
	for k := 0; k <= 1; k++ {
		if color == floor.getTile(y-1, x+k) {
			adjancent += 1
		}
		if color == floor.getTile(y+1, x+k) {
			adjancent += 1
		}
		if color == floor.getTile(y+(k*2-1)*2, 0) {
			adjancent += 1
		}
		if color == floor.getTile(y, x+k*2-1) {
			adjancent += 1
		}
	}
	return adjancent
}

func (oldFloor Floor) flipLive() *Floor {
	floor := newFloor()
	center := (len(oldFloor.tiles)) / 2
	for y := -center - 1; y <= center+2; y++ {
		for x := -center - 1; x <= center+2; x++ {
			if oldFloor.getTile(y, x) {
				if oldFloor.countAdjancent(y, x, false) == 2 {
					floor.flip(y, x)
				}
			} else {
				count := oldFloor.countAdjancent(y, x, false)
				if count == 1 || count == 2 {
					floor.flip(y, x)
				}
			}
		}
	}
	return floor
}

type Walker struct {
	y, x  int
	floor *Floor
}

func newWalker(floor *Floor) *Walker {
	walker := new(Walker)
	walker.floor = floor

	return walker
}

func (walker *Walker) step(y, x int) {
	walker.y += y
	walker.x += x

	// walker.floor.flip(walker.y, walker.x)
}

func (walker *Walker) walk(directions string) {
	var y int

	for _, direction := range directions {
		switch direction {
		case 's':
			y = 1
		case 'n':
			y = -1
		case 'e':
			if y == 0 {
				walker.step(0, -2)
			} else {
				walker.step(y, -1)
				y = 0
			}

		case 'w':
			if y == 0 {
				walker.step(0, 2)
			} else {
				walker.step(y, 1)
				y = 0
			}
		default:
			panic("Unexpected input")
		}
	}
	walker.floor.flip(walker.y, walker.x)
}

func main() {
	input, err := ioutil.ReadFile("adv24.txt")
	if err != nil {
		panic(err)
	}

	floor := newFloor()
	for _, move := range strings.Split(
		strings.TrimSpace(string(input)),
		"\n",
	) {
		walker := newWalker(floor)
		walker.walk(move)
	}

	fmt.Println(floor)
	fmt.Println("Part1:", floor.countBlack())
	/*
		for i := 1; i <= 10; i++ {
			floor = floor.flipLive()
			fmt.Println(floor)
			fmt.Println("Day", i, floor.countBlack())
		}
	*/
}
