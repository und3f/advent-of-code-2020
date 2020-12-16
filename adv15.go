package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const maxTurns = 2020
const maxTurns2 = 30000000

type GameState struct {
	turn       uint64
	numbers    map[uint64][]uint64
	lastNumber uint64
}

func newGame(starting []uint64) *GameState {
	state := new(GameState)
	state.turn = 0
	state.numbers = make(map[uint64][]uint64)
	for _, number := range starting {
		state.Spoke(number)
	}
	return state
}

func (state *GameState) Spoke(number uint64) {
	state.turn++
	// fmt.Println("Spoke", state.turn, number)
	v, exists := state.numbers[number]
	if !exists {
		v = make([]uint64, 0)
	}

	v = append(v, state.turn)
	if len(v) > 2 {
		v = v[len(v)-2:]
	}
	state.numbers[number] = v
	state.lastNumber = number
}

func (state *GameState) NextNumber() uint64 {
	nextNumber := uint64(0)
	v, exists := state.numbers[state.lastNumber]
	if exists && len(v) > 1 {
		// a := v[len(v)-2 : len(v)-1]
		nextNumber = v[len(v)-1] - v[len(v)-2]
	}

	return nextNumber
}

func playGame(starting []uint64, turns uint64) uint64 {
	game := newGame(starting)

	for game.turn < turns {
		nextNumber := game.NextNumber()
		// fmt.Println(game.turn+1, nextNumber, game.numbers)
		game.Spoke(nextNumber)
	}

	// fmt.Println(game)
	return game.lastNumber
}

func main() {
	reportLine, err := ioutil.ReadFile("adv15.txt")
	if err != nil {
		panic(err)
	}
	line := strings.Split(string(reportLine), "\n")[0]
	var starting []uint64
	for _, str := range strings.Split(line, ",") {
		v, err := strconv.Atoi(str)
		if err != nil {
			panic("Wrong number")
		}
		starting = append(starting, uint64(v))
	}
	var partOne, partTwo uint64

	partOne = playGame(starting, maxTurns)
	fmt.Println("Part One:", partOne)

	partTwo = playGame(starting, maxTurns2)
	fmt.Println("Part Two:", partTwo)
}
