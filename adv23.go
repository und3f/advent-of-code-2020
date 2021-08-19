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

type Cup struct {
	value uint
	next  *Cup
}

type Deck2 struct {
	cups                    map[uint]*Cup
	pointer                 *Cup
	lowestCard, highestCard uint
}

func newDeck2(cups []uint) *Deck2 {
	deck := new(Deck2)
	deck.cups = make(map[uint]*Cup)
	var prevCup *Cup
	for _, value := range cups {
		cup := new(Cup)
		cup.value = value
		if prevCup != nil {
			prevCup.next = cup
		}
		deck.cups[value] = cup
		prevCup = cup
	}
	firstCup := deck.cups[cups[0]]
	prevCup.next = firstCup
	deck.pointer = firstCup

	deck.lowestCard = cups[0]
	deck.highestCard = cups[0]
	for _, card := range cups {
		if card < deck.lowestCard {
			deck.lowestCard = card
		} else if card > deck.highestCard {
			deck.highestCard = card
		}
	}

	return deck
}

func (deck *Deck2) Move() {
	destination := deck.pointer.value - 1

	for isInPicked := true; isInPicked; {
		isInPicked = false
		if destination < deck.lowestCard {
			destination = deck.highestCard
		}

		pickedPointer := deck.pointer.next
		for i := 1; i <= PickupLen; i++ {
			if pickedPointer.value == destination {
				isInPicked = true
				destination--
				break
			}
			pickedPointer = pickedPointer.next
		}
	}

	pickedLast := deck.pointer
	for i := 1; i <= PickupLen; i++ {
		pickedLast = pickedLast.next
	}

	pickedFirst := deck.pointer.next
	deck.pointer.next = pickedLast.next

	destinationPointer := deck.cups[destination]
	// fmt.Println(deck.pointer, destination, deck.cups[4])
	pickedLast.next = destinationPointer.next
	destinationPointer.next = pickedFirst

	deck.pointer = deck.pointer.next
}

func (deck *Deck2) PrintCups() {
	fmt.Printf("(%d) ", deck.pointer.value)
	for cup := deck.pointer.next; cup != deck.pointer; cup = cup.next {
		fmt.Printf("%d ", cup.value)
	}
	fmt.Println()
}

func (deck *Deck2) Part2() uint64 {
	cupOne := deck.cups[1]
	// fmt.Println(cupOne.next.value, cupOne.next.next.value)
	return uint64(cupOne.next.value) * uint64(cupOne.next.next.value)
}

func (deck *Deck2) LabelsAfterOne() uint64 {
	labels := uint(0)
	firstCup := deck.cups[1]

	for i := firstCup.next; i != firstCup; i = i.next {
		labels = labels*10 + i.value
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

	cupsMillion := make([]uint, DeckSizeP2)
	copy(cupsMillion, cups)

	deck := newDeck2(cups)
	for i := 0; i < MOVES; i++ {
		// fmt.Println("Move", i+1)
		deck.Move()
		// deck.PrintCups()
	}
	PartOne := strconv.FormatUint(deck.LabelsAfterOne(), 10)

	fmt.Println("Part One:", PartOne)

	nextCard := deck.highestCard + 1
	for i := len(cups); i < len(cupsMillion); i++ {
		cupsMillion[i] = nextCard
		nextCard++
	}

	deck2 := newDeck2(cupsMillion)
	for i := 0; i < MovesP2; i++ {
		// fmt.Println("Pointer", deck2.pointer)
		if i%100 == 0 {
			// fmt.Printf("Move %d (%f%%)\n", i+1, float64(i+1)/float64(MovesP2)*100.0)
		}
		deck2.Move()
	}
	fmt.Println("Part Two:", strconv.FormatUint(deck2.Part2(), 10))
}
