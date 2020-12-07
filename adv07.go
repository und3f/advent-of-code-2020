package main

import (
	"fmt"
	"io/ioutil"
	//"sort"
	"strings"
	//"regexp"
	"strconv"
)

type BagAmount struct {
	Amount int
	Name   string
}

type Bag struct {
	Contain []BagAmount
}

type BagMap map[string]Bag

func DeepCalc(bagname string, bags BagMap) int {
	count := 1

	bag, exists := bags[bagname]
	if !exists {
		panic("Unknown bag " + bagname)
	}

	for i := range bag.Contain {
		count += bag.Contain[i].Amount * DeepCalc(bag.Contain[i].Name, bags)
	}

	return count
}

func main() {
	part1 := 0
	part2 := 0

	input, err := ioutil.ReadFile("adv07.txt")
	if err != nil {
		panic(err)
	}

	bagmap := make(BagMap)
	bags := make(BagMap)
	str := strings.ReplaceAll(string(input), " bags", "")
	str = strings.ReplaceAll(str, " bag", "")
	lines := strings.Split(str, ".\n")

	var declarations [][]string

	for i := range lines {
		if len(lines[i]) == 0 {
			continue
		}
		declaration := strings.Split(strings.Trim(lines[i], "\n"), " contain ")
		contain := strings.Split(declaration[1], ", ")
		declaration = append([]string{declaration[0]}, contain...)

		bag, exists := bags[declaration[0]]
		if !exists {
			bag = Bag{}
		}
		bags[declaration[0]] = bag
		if declaration[1] == "no other" {
		} else {
			fmt.Println(declaration[0])
			for i := 1; i < len(declaration); i++ {
				insideBag := strings.SplitN(declaration[i], " ", 2)

				amount, err := strconv.Atoi(insideBag[0])
				if err != nil {
					panic(err)
				}

				fmt.Printf("%q\n", insideBag)
				bag.Contain = append(bag.Contain, BagAmount{amount, insideBag[1]})
			}
			declarations = append(declarations, declaration)
		}
		fmt.Println(bag)
		bags[declaration[0]] = bag
	}

	bagmap["shiny gold"] = Bag{}
	for found := true; found; {
		found = false
		for j := range declarations {
			declaration := declarations[j]
			if len(declaration) == 0 {
				continue
			}

			for i := 1; i < len(declaration); i++ {
				insideBag := strings.SplitN(declaration[i], " ", 2)
				if _, exists := bagmap[insideBag[1]]; exists {
					part1++
					found = true
					declarations[j] = []string{}
					bagmap[declaration[0]] = Bag{}
					break
				}
			}
		}
	}

	fmt.Println(bags)
	part2 = DeepCalc("shiny gold", bags) - 1

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
