package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	// "regexp"
	//"sort"
	"strings"
)

const MOVES = 10
const PickupLen = 3

const DeckSizeP2 = 1000000
const MovesP2 = 10000000

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

type Deck struct {
	cups                    []uint
	pointer                 int
	lowestCard, highestCard uint
}

func newDeck(cups []uint) *Deck {
	deck := new(Deck)

	deck.cups = cups
	deck.pointer = 0
	deck.lowestCard = cups[0]
	deck.highestCard = cups[0]
	for _, card := range deck.cups {
		if card < deck.lowestCard {
			deck.lowestCard = card
		} else if card > deck.highestCard {
			deck.highestCard = card
		}
	}

	return deck
}

func (deck *Deck) Move() {
	// printCups(deck.cups, deck.pointer)
	// fmt.Println(deck)

	destination := deck.cups[deck.pointer] - 1
	pickup := make([]uint, PickupLen)
	copy(pickup, deck.cups[deck.pointer+1:deck.pointer+1+PickupLen])

	/*
		fmt.Println("Pickup", pickup)
		fmt.Println("Destination", destination)
	*/

	for isInPicked := true; isInPicked; {
		isInPicked = false
		if destination < deck.lowestCard {
			destination = deck.highestCard
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
	for i, card := range deck.cups {
		if card == destination {
			destinationPointer = i
			break
		}
	}

	deck.cups = append(deck.cups[0:deck.pointer+1], deck.cups[deck.pointer+4:]...)
	if deck.pointer < destinationPointer {
		destinationPointer -= len(pickup)
	}
	// printCups(deck.cups, deck.pointer)
	// fmt.Println("Pickup", pickup)
	// fmt.Println("Destination", destination)
	deck.cups = append(deck.cups[0:destinationPointer+1], append(pickup, deck.cups[destinationPointer+1:]...)...)
	if destinationPointer < deck.pointer {
		deck.pointer += len(pickup)
	}
	deck.pointer++

	if deck.pointer == len(deck.cups) {
		deck.pointer = 0
	} else if deck.pointer+PickupLen >= len(deck.cups) {
		deck.pointer -= PickupLen
		deck.cups = append(deck.cups[PickupLen:], deck.cups[0:PickupLen]...)
	}
	/*
		fmt.Println(deck.pointer+PickupLen, len(deck.cups))
		printCups(deck.cups, deck.pointer)

		fmt.Println()
		fmt.Println()
	*/
	//return deck
}

func (deck *Deck) LabelsAfterOne() uint64 {
	labels := uint(0)
	var firstCup int
	for i, cup := range deck.cups {
		if cup == 1 {
			firstCup = i
			break
		}
	}
	p1cups := make([]uint, 0)
	if firstCup+1 < len(deck.cups) {
		p1cups = deck.cups[firstCup+1:]
	}
	p1cups = append(p1cups, deck.cups[0:firstCup]...)

	for _, v := range p1cups {
		labels = labels*10 + v
	}

	return uint64(labels)
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

	deck := newDeck(cups)
	for i := 0; i < MOVES; i++ {
		fmt.Println("Move", i+1)
		deck.Move()
	}
	PartOne := strconv.FormatUint(deck.LabelsAfterOne(), 10)

	fmt.Println("Part One:", PartOne)

	nextCard := deck.highestCard + 1
	cupsMillion := make([]uint, DeckSizeP2)
	copy(cupsMillion, cups)
	for i := len(cups); i < len(cupsMillion); i++ {
		cupsMillion[i] = nextCard
		nextCard++
	}

	deck = newDeck(cupsMillion)
	for i := 0; i < MovesP2; i++ {
		if i%100 == 0 {
			fmt.Printf("Move %d (%f%%)\n", i+1, float64(i+1)/float64(MovesP2)*100.0)
		}
		deck.Move()
	}
	fmt.Println("Part Two:", strconv.FormatUint(deck.LabelsAfterOne(), 10))
}
