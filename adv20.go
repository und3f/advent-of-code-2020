package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Tile struct {
	id    int
	image [][]bool
}

func (tile *Tile) Rotate() *Tile {
	newTile := new(Tile)
	height := len(tile.image)
	width := len(tile.image[0])
	if height != width {
		panic("Unable to rotate")
	}

	newTile.image = make([][]bool, height)
	for i := range tile.image {
		newTile.image[i] = make([]bool, width)
	}

	for j := range tile.image[0] {
		for i := height - 1; i >= 0; i-- {
			newTile.image[j][height-1-i] = tile.image[i][j]
		}
	}

	newTile.id = tile.id
	return newTile
}

func (tile *Tile) FlipY() *Tile {
	newTile := new(Tile)
	width := len(tile.image[0])

	newTile.image = make([][]bool, width)
	for i := range tile.image {
		newTile.image[i] = make([]bool, width)
	}

	for j := 0; j < width; j++ {
		for i, _ := range tile.image {
			newTile.image[width-1-i][j] = tile.image[i][j]
		}
	}

	newTile.id = tile.id

	return newTile
}

func (tile *Tile) FlipX() *Tile {
	newTile := new(Tile)
	width := len(tile.image[0])

	newTile.image = make([][]bool, width)
	for i := range tile.image {
		newTile.image[i] = make([]bool, width)
	}

	for i, row := range tile.image {
		for j, cell := range row {
			newTile.image[i][width-1-j] = cell
		}
	}

	newTile.id = tile.id
	return newTile
}

func (tile *Tile) ImageString() (out string) {
	for _, row := range tile.image {
		for _, cell := range row {
			symbol := '.'
			if cell {
				symbol = '#'
			}

			out += string(symbol)
		}
		out += "\n"
	}

	return
}

func (tile *Tile) String() string {
	out := fmt.Sprintf("Tile %d:\n", tile.id)
	out += tile.ImageString()

	return out
}

func (tile *Tile) Borders() []uint {
	height := len(tile.image)
	width := len(tile.image[0])

	var borderUp, borderDown uint
	for i, _ := range tile.image[0] {
		up := tile.image[0][i]
		down := tile.image[height-1][i]
		a := uint(0)
		if up {
			a = 1
		}
		b := uint(0)
		if down {
			b = 1
		}

		borderUp = borderUp<<1 + a
		borderDown = borderDown<<1 + b
	}

	var borderL, borderR uint
	for _, row := range tile.image {
		l, r := uint(0), uint(0)
		if row[0] {
			l = 1
		}
		if row[width-1] {
			r = 1
		}
		borderL = borderL<<1 + l
		borderR = borderR<<1 + r
	}

	return []uint{borderUp, borderR, borderDown, borderL}
}

func parseTiles(reportLine string) []Tile {

	lines := strings.Split(strings.TrimSuffix(string(reportLine), "\n"), "\n")
	tiles := make([]Tile, 0)

	var currentTile *Tile
	for _, line := range lines {
		if len(line) == 0 {
			continue
		} else if line[:4] == "Tile" {
			if currentTile != nil {
				tiles = append(tiles, *currentTile)
			}
			currentTile = new(Tile)
			numberEnd := strings.IndexRune(line, ':')
			id, err := strconv.Atoi(line[5:numberEnd])
			if err != nil {
				panic(err)
			}
			currentTile.id = id
		} else {
			row := make([]bool, len(line))
			for i, v := range line {
				row[i] = v == '#'
			}
			currentTile.image = append(currentTile.image, row)
		}
	}

	if currentTile != nil {
		tiles = append(tiles, *currentTile)
	}

	return tiles
}

type Image [][]*Tile

const EmptyTile = `??????????
??????????
??????????
??????????
??????????
??????????
??????????
??????????
??????????
??????????`

func (image *Image) String() (result string) {
	for i, row := range *image {
		tilesStringLines := make([][]string, len(row))
		for i, tile := range row {
			str := EmptyTile
			if tile != nil {
				str = tile.ImageString()
			}
			tilesStringLines[i] = strings.Split(str, "\n")
		}

		// fmt.Println("Tile string", tilesStringLines)
		for i := 0; i < len(tilesStringLines[0])-1; i++ {
			for j := range tilesStringLines {
				result += tilesStringLines[j][i]
				if j < len(tilesStringLines)-1 {
					result += "|"
				}
			}
			result += "\n"
		}

		if i < len(*image)-1 {
			result += "----------+----------+----------\n"
		}
	}

	return
}

func (image *Image) ComposeTile() Tile {
	result := "Tile 1:\n"

	for _, row := range *image {
		tilesStringLines := make([][]string, len(row))
		for i, tile := range row {
			str := EmptyTile
			if tile != nil {
				str = tile.ImageString()
			}
			tilesStringLines[i] = strings.Split(str, "\n")
		}

		// fmt.Println("Tile string", tilesStringLines)
		for i := 1; i < len(tilesStringLines[0])-2; i++ {
			for j := range tilesStringLines {
				line := tilesStringLines[j][i]
				result += line[1 : len(line)-1]
			}
			result += "\n"
		}

	}

	return parseTiles(result)[0]
}

func removeBorder(image string) string {
	lines := strings.Split(image, "\n")
	lines = lines[1 : len(lines)-2]

	for i, line := range lines {
		lines[i] = line[1 : len(line)-1]
	}

	return strings.Join(lines, "\n")
}

func (image *Image) Result() int {
	height := len(*image)

	return (*image)[0][0].id * (*image)[0][height-1].id * (*image)[height-1][0].id * (*image)[height-1][height-1].id

}

func (image *Image) StringVerbose() (result string) {
	for _, row := range *image {
		for _, tile := range row {
			if tile != nil {
				result += fmt.Sprintf("%d    ", tile.id)
			} else {
				result += "        "
			}
		}

		result += "\n"
	}
	return
}

func (tile1 *Tile) connectsRight(tile2 *Tile) bool {
	if tile1 == nil || tile2 == nil {
		return true
	}

	width := len(tile1.image[0])
	for i := range tile2.image {
		if tile1.image[i][width-1] != tile2.image[i][0] {
			return false
		}
	}

	return true
}

func (tile1 *Tile) connectsDown(tile2 *Tile) bool {
	if tile1 == nil || tile2 == nil {
		return true
	}

	height := len(tile1.image)
	for i, v := range tile2.image[0] {
		if tile1.image[height-1][i] != v {
			return false
		}
	}

	return true
}

func (tile1 *Tile) bitmaskAtPosition(offsetY, offsetX int, bitmask *Tile) bool {
	for i, bitmaskRow := range bitmask.image {
		for j, v := range bitmaskRow {
			if v {
				if !tile1.image[i+offsetY][j+offsetX] {
					return false
				}
			}
		}
	}

	return true
}

func (tile *Tile) ExtractBitmask(offsetY, offsetX int, bitmask *Tile) {
	for i, bitmaskRow := range bitmask.image {
		for j, v := range bitmaskRow {
			if v {
				tile.image[i+offsetY][j+offsetX] = false
			}
		}
	}
}

func (tile *Tile) Score() (score int) {
	for _, row := range tile.image {
		for _, cell := range row {
			if cell {
				score++
			}
		}
	}

	return score
}

func (tile1 *Tile) FindBitmask(bitmask *Tile) int {
	bitmaskHeight := len(bitmask.image)
	bitmaskWidth := len(bitmask.image[0])
	var found int

	for i := 0; i < len(tile1.image)-bitmaskHeight; i++ {
		for j := 0; j < len(tile1.image[i])-bitmaskWidth; j++ {
			if tile1.bitmaskAtPosition(i, j, bitmask) {
				found++
				tile1.ExtractBitmask(i, j, bitmask)
			}
		}
	}

	return found
}

func isCorrect(image *Image) bool {
	for i := 0; i < len(*image); i++ {
		for j := 0; j < len((*image)[i])-1; j++ {
			if !(*image)[i][j].connectsRight((*image)[i][j+1]) {
				return false
			}
		}
	}

	for i := 0; i < len(*image)-1; i++ {
		for j := 0; j < len(*image); j++ {
			if !(*image)[i][j].connectsDown((*image)[i+1][j]) {
				return false
			}
		}
	}

	return true
}

func allFlips(tile *Tile) (res [4]*Tile) {
	res[0] = tile
	res[1] = tile.FlipX()
	res[2] = tile.FlipY()
	res[3] = res[1].FlipY()

	return
}

func TryNextTile(tiles []Tile, image *Image, position int) bool {
	imageHeight := len(*image)
	if position >= len(*image)*len((*image)[0]) {
		/*
				fmt.Println("Found solution")
				fmt.Println(image.String())
				fmt.Println(image.StringVerbose())
			fmt.Println((*image)[0][0].id * (*image)[0][2].id * (*image)[2][0].id * (*image)[2][2].id)
		*/
		return true
	}

	// fmt.Println("Trying next tile", position)
	for i := 0; i < len(tiles); i++ {
		tile := tiles[i]

		restTiles := make([]Tile, len(tiles)-1)
		copy(restTiles, tiles[:i])
		copy(restTiles[i:], tiles[i+1:])

		/*
			fmt.Println("Trying tile", tile.id, i)

			fmt.Print("Tiles")
			for _, tile := range tiles {
				fmt.Print(tile.id, " ")
			}
			fmt.Println()

			fmt.Print("Rest tiles:")
			for _, tile := range restTiles {
				fmt.Print(tile.id, " ")
			}
			fmt.Println()
		*/

		rotatedTile := &tile
		for rotation := 1; rotation < 4; rotation++ {
			for _, transformedTile := range allFlips(rotatedTile) {
				(*image)[position/imageHeight][position%imageHeight] = transformedTile
				if isCorrect(image) {
					// fmt.Println(image.StringVerbose())
					// fmt.Println(image.String())

					if TryNextTile(restTiles, image, position+1) {
						return true
					}
				}
			}
			rotatedTile = rotatedTile.Rotate()
		}
	}

	(*image)[position/imageHeight][position%imageHeight] = nil
	return false
}

func TryEveryTile(tiles []Tile) Image {
	height := int(math.Sqrt(float64(len(tiles))))
	image := make(Image, height)
	for i := range image {
		image[i] = make([]*Tile, height)
	}

	// fmt.Println(image)
	TryNextTile(tiles, &image, 0)
	return image
}

func FindMonsters(tile, monster *Tile) *Tile {
	rotatedTile := tile
	for rotation := 1; rotation < 4; rotation++ {
		for _, transformedTile := range allFlips(rotatedTile) {
			if transformedTile.FindBitmask(monster) > 0 {
				return transformedTile
			}
		}
		rotatedTile = rotatedTile.Rotate()
	}

	return nil
}

func main() {
	reportLine, err := ioutil.ReadFile("adv20.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Println(lines)
	tiles := parseTiles(string(reportLine))

	// fmt.Println(tiles)
	/*
		for _, tile := range tiles {
			fmt.Println(tile.String())
			fmt.Println(tile.Borders())
		}
	*/

	var partOne, partTwo int

	// fmt.Println()
	// fmt.Println(tiles[0].FitToEdge(tiles[1].Borders()[1], 3))

	/*
		var image Image
		image[0][0] = tiles[1].FlipY()
		image[0][1] = tiles[0].FlipY()
		image[0][2] = &tiles[8]
		image[1][0] = tiles[7].Rotate().Rotate().FlipX()
		image[1][1] = tiles[3].FlipY()
		//image[1][2] = *tiles[5].FlipX().Rotate()
		image[1][2] = tiles[5].Rotate().Rotate().Rotate().FlipX()
		image[2][0] = tiles[6].FlipY()
		image[2][1] = tiles[4].FlipY()
		image[2][2] = tiles[2].FlipX()

		fmt.Println(image.String())
	*/

	imageFound := TryEveryTile(tiles)
	// fmt.Println(imageFound.StringVerbose())
	partOne = imageFound.Result()
	fmt.Println("Part One:", partOne)

	// fmt.Println((imageFound.String()))
	completeImage := imageFound.ComposeTile()
	// completeImage = *(completeImage.Rotate().FlipX())
	// fmt.Println(completeImage.String())

	monster := parseTiles(`Tile 2:
                  # 
#    ##    ##    ###
 #  #  #  #  #  #   
`)[0]
	// fmt.Println(monster.String())

	monsterless := FindMonsters(&completeImage, &monster)
	// fmt.Println(monsterless)
	partTwo = monsterless.Score()
	fmt.Println("Part Two:", partTwo)
}
