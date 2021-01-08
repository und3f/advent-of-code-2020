package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	// "regexp"
	//"sort"
	"strings"
)

func prepareHandString(str string) []uint {
	var hand []uint
	for _, cardName := range strings.Split(str, "\n")[1:] {
		if card, err := strconv.ParseUint(cardName, 10, 32); err != nil {
			panic(err)
		} else {
			hand = append(hand, uint(card))
		}
	}

	return hand
}

func (round *Round) drawCards() Deck {
	table := Deck{(*round.Player(0))[0], (*round.Player(1))[0]}

	for i := 0; i < len(*round); i++ {
		(*round)[i] = (*round)[i][1:]
	}

	return table
}

func (deck *Deck) winner() uint {
	winner := uint(0)
	max := (*deck)[0]
	for i, v := range (*deck)[1:] {
		if v > max {
			winner = uint(i + 1)
			max = v
		}
	}

	return winner
}

func (deck *Deck) append(deckB *Deck) {
	*deck = append(*deck, (*deckB)...)
}

func (round *Round) play() (stop bool) {
	table := round.drawCards()
	winner := table.winner()

	(*round)[winner] = append((*round)[winner], table[winner], table[(1^winner)])

	stop = len((*round)[0]) == 0 || len((*round)[1]) == 0
	return stop
}

type Deck []uint

func (deck *Deck) equal(deckB Deck) bool {
	// fmt.Println("equal", deck, deckB)

	if len(*deck) != len(deckB) {
		return false
	}

	for i, a := range *deck {
		b := deckB[i]

		if a != b {
			return false
		}
	}

	return true
}

func (deck *Deck) Score() (result uint) {
	multiplier := uint(1)
	for i := len(*deck) - 1; i >= 0; i-- {
		result += (*deck)[i] * multiplier
		multiplier++
	}
	return result
}

type Round []Deck

func (round *Round) equal(roundB Round) bool {
	// fmt.Println("equal", round, roundB)

	if len(*round) != len(roundB) {
		return false
	}
	for i, a := range *round {
		b := roundB[i]

		if !a.equal(b) {
			return false
		}
	}

	return true
}

func (round *Round) Player(player uint) *Deck {
	return &(*round)[player]
}

func (round *Round) String() (res string) {
	for i, deck := range *round {
		res += fmt.Sprintf("Player's %d's deck: ", i+1)
		res += strconv.Itoa(int(deck[0]))
		for _, v := range deck[1:] {
			res += fmt.Sprintf(", %d", v)
		}
		res += "\n"
	}

	res = res[:len(res)-1]

	return
}

func (round *Round) Clone(cards Deck) *Round {
	roundClone := make(Round, len(cards))

	for i, v := range cards {
		roundClone[i] = make(Deck, v)
		copy(roundClone[i], (*round)[i])
	}

	return &roundClone
}

type PreviousRounds []Round

func (previousRounds *PreviousRounds) isContaining(round Round) bool {
	for _, deckPrevious := range *previousRounds {
		if deckPrevious.equal(round) {
			return true
		}
	}
	return false
}

func (previousRounds *PreviousRounds) append(_round Round) {
	round := make(Round, len(_round))
	copy(round, _round)
	*previousRounds = append(*previousRounds, round)
}

var GlobalGameN = 1

func RecursiveCombat(round Round) (result uint, winner uint) {
	winner = 0

	/*
		roundN := 0
		gameN := GlobalGameN
		GlobalGameN++
	*/

	for previousRounds := new(PreviousRounds); !previousRounds.isContaining(round); {
		/*
			roundN++
			fmt.Printf("-- Round %d (Game %d) --\n", roundN, gameN)
			fmt.Println(round.String())
		*/

		previousRounds.append(round)

		// fmt.Println("Loop")
		// fmt.Println("Round before", round)
		table := round.drawCards()
		// fmt.Println("Table", table)
		haveEnought := true
		for i, value := range table {
			if !(uint(len(*round.Player(uint(i)))) >= value) {
				haveEnought = false
				break
			}
		}

		var localWinner uint
		if haveEnought {
			_, localWinner = RecursiveCombat(*(round.Clone(table)))
		} else {
			localWinner = table.winner()
		}
		// fmt.Printf("Player %d wins round %d of game %d!\n", localWinner+1, roundN, gameN)

		if localWinner == 1 {
			table[0], table[1] = table[1], table[0]
		}
		round.Player(localWinner).append(&table)

		if len(round[0]) == 0 {
			winner = 1
			break
		} else if len(round[1]) == 0 {
			break
		}

		// fmt.Println("Round after", round)
		// fmt.Println()
	}

	return round.Player(winner).Score(), winner
}

func main() {
	reportLine, err := ioutil.ReadFile("adv22.txt")
	if err != nil {
		panic(err)
	}
	decksLines := strings.Split(strings.TrimSpace(string(reportLine)), "\n\n")
	round1 := &Round{
		prepareHandString(decksLines[0]),
		prepareHandString(decksLines[1]),
	}
	//fmt.Println(round1)

	for stop := false; !stop; {
		stop = round1.play()
	}
	winner := round1.Player(0)
	if len(*winner) == 0 {
		winner = round1.Player(1)
	}

	var partOne uint
	partOne = winner.Score()

	fmt.Println("Part one:", partOne)

	round := Round{
		prepareHandString(decksLines[0]),
		prepareHandString(decksLines[1]),
	}
	partTwo, _ := RecursiveCombat(round)
	fmt.Println("Part two:", partTwo)
}
