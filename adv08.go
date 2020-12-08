package main

import (
	"fmt"
	"io/ioutil"
	//"sort"
	"strings"
	//"regexp"
	"strconv"
)

type Instruction struct {
	op    string
	value int
}

func execute(instructions []Instruction) (int, bool) {
	executed := make([]bool, len(instructions))
	value := 0
	ip := 0
	for ip < len(instructions) && executed[ip] != true {
		executed[ip] = true
		op := instructions[ip]
		// fmt.Printf("%d: %s %d (%d)\n", ip, op.op, op.value, value)
		switch op.op {
		case "nop":
			ip++
		case "acc":
			value += op.value
			ip++
		case "jmp":
			ip += op.value
		}
	}

	return value, ip >= len(instructions)

}

func main() {
	part1 := 0
	part2 := 0

	input, err := ioutil.ReadFile("adv08.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	instructions := make([]Instruction, len(lines)-1)
	for i := range lines {
		if len(lines[i]) == 0 {
			continue
		}

		arg := strings.Split(lines[i], " ")
		arg1 := arg[1]
		value, err := strconv.Atoi(arg1)
		if err != nil {
			panic("Wrong value " + arg1)
		}

		instructions[i] = Instruction{
			arg[0], value,
		}
	}

	fmt.Println(instructions)
	value, _ := execute(instructions)

	for i := range instructions {
		minstr := make([]Instruction, len(instructions))
		copy(minstr, instructions)

		switch instructions[i].op {
		case "nop":
			minstr[i].op = "jmp"
		case "jmp":
			minstr[i].op = "nop"
		default:
			continue
		}
		value, finished := execute(minstr)
		if finished {
			part2 = value
			break
		}
	}

	part1 = value
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
