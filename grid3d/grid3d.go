package grid3d

import "fmt"
import "bytes"

type Coord [3]int

type GridSlice struct {
	slice [][]bool
}

type Grid3D struct {
	grid   []GridSlice
	offset int
}

type HyperCoord [4]int
type HyperCube struct {
	cube   []*Grid3D
	offset int
}

func NewEmptySlice() GridSlice {
	return GridSlice{slice: [][]bool{[]bool{false}}}
}

func NewGrid3D() *Grid3D {
	grid := new(Grid3D)
	grid.grid = []GridSlice{GridSlice{slice: [][]bool{[]bool{false}}}}

	return grid
}

func NewHyperCube() *HyperCube {
	cube := new(HyperCube)
	cube.cube = []*Grid3D{NewGrid3D()}
	return cube
}

func (slice *GridSlice) CoordToIndex(coord Coord) (int, int) {
	yOffset := (len(slice.slice) - 1) / 2
	xOffset := (len(slice.slice[0]) - 1) / 2
	y := coord[1] + yOffset
	x := coord[2] + xOffset
	return y, x
}

func (slice *GridSlice) Get(coord Coord) bool {
	y, x := slice.CoordToIndex(coord)
	if y < 0 || x < 0 || y >= len(slice.slice) || x >= len(slice.slice[y]) {
		return false
	}
	// fmt.Println("GET", y, x)

	return slice.slice[y][x]
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func (slice *GridSlice) Resize(newSize int) {
	newSlice := make([][]bool, newSize)
	for i := range newSlice {
		newSlice[i] = make([]bool, newSize)
	}

	inserted := (newSize - len(slice.slice)) / 2
	// fmt.Println("INserted", inserted)
	for i, row := range slice.slice {
		copy(newSlice[inserted+i][inserted:], row)
	}
	slice.slice = newSlice
}

func (slice *GridSlice) Set(coord Coord, value bool) {

	ySize := Abs(coord[1]*2) + 1
	xSize := Abs(coord[2]*2) + 1
	maxSize := ySize
	if xSize > maxSize {
		maxSize = xSize
	}

	// fmt.Println("MAx Size", ySize, xSize)

	if maxSize > len(slice.slice) {
		// fmt.Println("RESIZE", maxSize, len(slice.slice))
		slice.Resize(maxSize)
	}

	y, x := slice.CoordToIndex(coord)
	slice.slice[y][x] = value
}

func (slice *GridSlice) Print() {
	var buffer bytes.Buffer
	for _, row := range slice.slice {
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

func (slice *GridSlice) CalculateActive() int {
	var sum int
	for _, row := range slice.slice {
		for _, cell := range row {
			if cell {
				sum++
			}
		}
	}

	return sum
}

func (grid *Grid3D) Get(coord Coord) bool {
	z := coord[0] + grid.offset
	if z < 0 || z >= len(grid.grid) {
		return false
	}

	return grid.grid[z].Get(coord)
}

func (grid *Grid3D) Set(coord Coord, value bool) {
	z := coord[0] + grid.offset

	if z < 0 {
		insert := -z
		fmt.Println("PREPEND", coord)
		for i := 0; i < insert; i++ {
			grid.grid = append([]GridSlice{NewEmptySlice()}, grid.grid...)
		}
		grid.offset += insert
		z += insert
	} else if z >= len(grid.grid) {
		fmt.Println("APPEND", coord)
		insert := z - len(grid.grid) + 1
		for i := 0; i < insert; i++ {
			grid.grid = append(grid.grid, NewEmptySlice())
		}
	}

	grid.grid[z].Set(coord, value)
}

func (grid *Grid3D) CalculateActive() int {
	var sum int
	for _, slice := range grid.grid {
		sum += slice.CalculateActive()
	}

	return sum
}

func (grid *Grid3D) CalcNeighbors(coord Coord) int {
	sum := 0
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}

				checking := Coord{coord[0] + z, coord[1] + y, coord[2] + x}
				if grid.Get(checking) {
					sum++

					if sum > 3 {
						return sum
					}
				}
			}
		}
	}
	return sum
}

func (grid *Grid3D) Cycle() *Grid3D {

	newGrid := NewGrid3D()
	minZ := -grid.offset
	maxZ := len(grid.grid) - 1 - grid.offset

	var maxSlice int
	for _, slice := range grid.grid {
		sliceLen := len(slice.slice)
		if sliceLen > maxSlice {
			maxSlice = sliceLen
		}
	}
	sliceOffset := (maxSlice - 1) / 2

	for z := minZ - 1; z <= maxZ+1; z++ {
		for y := -sliceOffset - 1; y <= sliceOffset+1; y++ {
			for x := -sliceOffset - 1; x <= sliceOffset+1; x++ {
				coord := Coord{z, y, x}
				// fmt.Println(coord)
				cell := grid.Get(coord)
				// fmt.Println("Neigh")
				neighbors := grid.CalcNeighbors(coord)
				if cell {
					if neighbors == 2 || neighbors == 3 {
						newGrid.Set(coord, true)
						// fmt.Println(coord, true)
						// newGrid.Print()
					}
				} else {
					if neighbors == 3 {
						newGrid.Set(coord, true)
						// fmt.Println(coord, true)
						// newGrid.Print()
					}
				}
			}
		}
	}

	return newGrid
}

func (grid *Grid3D) Print() {
	for z, slice := range grid.grid {
		fmt.Printf("z=%d\n", z-grid.offset)
		slice.Print()
	}
}

func (cube *HyperCube) Print() {
	for w, grid := range cube.cube {
		for z, slice := range grid.grid {
			fmt.Printf("z=%d, w=%d\n", z-grid.offset, w-cube.offset)
			slice.Print()
		}
	}
}

func (cube *HyperCube) Get(coord HyperCoord) bool {
	w := coord[0] + cube.offset
	if w < 0 || w >= len(cube.cube) {
		return false
	}

	return cube.cube[w].Get(Coord{coord[1], coord[2], coord[3]})
}

func (cube *HyperCube) Set(coord HyperCoord, value bool) {
	w := coord[0] + cube.offset

	if w < 0 {
		insert := -w
		fmt.Println("PREPEND", coord)
		for i := 0; i < insert; i++ {
			cube.cube = append([]*Grid3D{NewGrid3D()}, cube.cube...)
		}
		cube.offset += insert
		w += insert
	} else if w >= len(cube.cube) {
		fmt.Println("APPEND", coord)
		insert := w - len(cube.cube) + 1
		for i := 0; i < insert; i++ {
			cube.cube = append(cube.cube, NewGrid3D())
		}
	}

	cube.cube[w].Set(Coord{coord[1], coord[2], coord[3]}, value)
}

func (cube *HyperCube) CalculateActive() int {
	var sum int
	for _, grid := range cube.cube {
		sum += grid.CalculateActive()
	}

	return sum
}

func (cube *HyperCube) CalcNeighbors(coord HyperCoord) int {
	sum := 0
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}

					checking := HyperCoord{coord[0] + w, coord[1] + z, coord[2] + y, coord[3] + x}
					if cube.Get(checking) {
						sum++

						if sum > 3 {
							return sum
						}
					}
				}
			}
		}
	}
	return sum
}

func (cube *HyperCube) Cycle() *HyperCube {

	newCube := NewHyperCube()
	minW := -cube.offset
	maxW := len(cube.cube) - 1 - cube.offset

	minZ := int(^uint(0) >> 1)
	maxZ := 0

	var maxSlice int
	for _, grid := range cube.cube {

		localMinZ := -grid.offset
		localMaxZ := len(grid.grid) - 1 - grid.offset
		if minZ > localMinZ {
			minZ = localMinZ
		}
		if maxZ < localMaxZ {
			maxZ = localMaxZ
		}

		for _, slice := range grid.grid {
			sliceLen := len(slice.slice)
			if sliceLen > maxSlice {
				maxSlice = sliceLen
			}
		}
	}
	sliceOffset := (maxSlice - 1) / 2

	for w := minW - 1; w <= maxW+1; w++ {
		for z := minZ - 1; z <= maxZ+1; z++ {
			for y := -sliceOffset - 1; y <= sliceOffset+1; y++ {
				for x := -sliceOffset - 1; x <= sliceOffset+1; x++ {
					coord := HyperCoord{w, z, y, x}
					// fmt.Println(coord)
					cell := cube.Get(coord)
					// fmt.Println("Neigh")
					neighbors := cube.CalcNeighbors(coord)
					if cell {
						if neighbors == 2 || neighbors == 3 {
							newCube.Set(coord, true)
							// fmt.Println(coord, true)
							// newGrid.Print()
						}
					} else {
						if neighbors == 3 {
							newCube.Set(coord, true)
							// fmt.Println(coord, true)
							// newGrid.Print()
						}
					}
				}
			}
		}
	}

	return newCube
}
