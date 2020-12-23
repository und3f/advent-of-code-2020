package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	// "regexp"
	//"sort"
	"strings"
)

const MOVES = 100
const PickupLen = 3

func printCups(cups []uint, pointer int) {
	for i, card := range cups {
		format := "%d "
		if i == pointer {
			format = "(%d) "
		}
		fmt.Printf(format, card)
	}
	fmt.Println()
}

func main() {
	reportLine, err := ioutil.ReadFile("adv23.txt")
	if err != nil {
		panic(err)
	}
	var cups []uint
	for _, cupChar := range strings.Split(
		strings.TrimSpace(string(reportLine)),
		"\n",
	)[0] {
		cups = append(cups, uint(cupChar-'0'))
	}
	lowestCard := cups[0]
	highestCard := cups[0]
	for _, card := range cups {
		if card < lowestCard {
			lowestCard = card
		} else if card > highestCard {
			highestCard = card
		}
	}

	pointer := 0

	for i := 0; i < MOVES; i++ {
		fmt.Println("Move", i+1)

		printCups(cups, pointer)

		destination := cups[pointer] - 1
		pickup := make([]uint, PickupLen)
		copy(pickup, cups[pointer+1:pointer+1+PickupLen])

		fmt.Println("Pickup", pickup)
		fmt.Println("Destination", destination)

		for isInPicked := true; isInPicked; {
			isInPicked = false
			if destination < lowestCard {
				destination = highestCard
			}

			for _, picked := range pickup {
				if picked == destination {
					isInPicked = true
					destination--
					break
				}
			}
		}

		destinationPointer := 0
		for i, card := range cups {
			if card == destination {
				destinationPointer = i
				break
			}
		}

		cups = append(cups[0:pointer+1], cups[pointer+4:]...)
		if pointer < destinationPointer {
			destinationPointer -= len(pickup)
		}
		printCups(cups, pointer)
		fmt.Println("Pickup", pickup)
		fmt.Println("Destination", destination)
		cups = append(cups[0:destinationPointer+1], append(pickup, cups[destinationPointer+1:]...)...)
		if destinationPointer < pointer {
			pointer += len(pickup)
		}
		pointer++

		if pointer == len(cups) {
			pointer = 0
		} else if pointer+PickupLen >= len(cups) {
			pointer -= PickupLen
			cups = append(cups[PickupLen:], cups[0:PickupLen]...)
		}
		fmt.Println(pointer+PickupLen, len(cups))
		printCups(cups, pointer)

		fmt.Println()
		fmt.Println()
	}

	var firstCup int
	for i, cup := range cups {
		if cup == 1 {
			firstCup = i
			break
		}
	}
	p1cups := make([]uint, 0)
	if firstCup+1 < len(cups) {
		p1cups = cups[firstCup+1:]
	}
	p1cups = append(p1cups, cups[0:firstCup]...)

	PartOne := ""
	for _, cup := range p1cups {
		PartOne += string(cup + '0')
	}

	fmt.Println("Part One:", PartOne)
}
