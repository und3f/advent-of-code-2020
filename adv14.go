package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const InstEnding = "] = "
const InstBegin = "mem["

func runProgram(instructions []string) uint64 {
	memory := make(map[int]uint64)
	var mask, maskOr uint64

	for _, inst := range instructions {
		if strings.Index(inst, "mask = ") == 0 {
			var err error
			mask, err = strconv.ParseUint(strings.ReplaceAll(inst[7:], "X", "1"), 2, 64)
			if err != nil {
				panic(err)
			}

			maskOr, err = strconv.ParseUint(strings.ReplaceAll(inst[7:], "X", "0"), 2, 64)
		} else if strings.Index(inst, InstBegin) == 0 {
			addressEnd := strings.Index(inst, InstEnding)
			address, err := strconv.Atoi(inst[len(InstBegin):addressEnd])
			if err != nil {
				panic(err)
			}

			number, err := strconv.ParseUint(inst[addressEnd+len(InstEnding):], 10, 64)
			if err != nil {
				panic(err)
			}

			memory[address] = (number & mask) | maskOr
		}
	}

	var sum uint64
	for _, v := range memory {
		sum += v
	}

	return sum
}

func runProgramV2(instructions []string) uint64 {
	memory := make(map[uint64]uint64)
	var mask, maskOr uint64
	var maskVariants []uint64

	for _, inst := range instructions {
		if strings.Index(inst, "mask = ") == 0 {
			var err error
			maskStr := inst[7:]

			mask, err = strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(maskStr, "1", "0"), "X", "1"), 2, 64)
			if err != nil {
				panic(err)
			}

			maskOr, err = strconv.ParseUint(strings.ReplaceAll(maskStr, "X", "0"), 2, 64)

			/*
				fmt.Println("Setting mask", maskStr)
				fmt.Println("Mask", strconv.FormatUint(mask, 2), "Mask or", strconv.FormatUint(maskOr, 2))
			*/

			var maskBits []uint64
			maskStrLen := len(maskStr)

			for i := 0; i < maskStrLen; i++ {
				if maskStr[maskStrLen-i-1] == 'X' {
					maskBits = append(maskBits, uint64(i))
				}
			}

			// fmt.Println("maskBits", maskBits)

			maxNumber := 1 << len(maskBits)
			maskVariants = make([]uint64, maxNumber)
			for i := 0; i < maxNumber; i++ {
				var variant uint64
				number := i
				for _, bitIndex := range maskBits {
					bit := number & 1
					number = number >> 1
					if bit > 0 {
						variant += 1 << bitIndex
						// fmt.Println("BitIndex", bitIndex, "added", 1<<bitIndex)
					}
				}
				// fmt.Println("Variant", variant)
				maskVariants[i] = variant
			}
			// fmt.Println("maskVariants", maskVariants)
		} else if strings.Index(inst, InstBegin) == 0 {
			addressEnd := strings.Index(inst, InstEnding)
			address, err := strconv.ParseUint(inst[len(InstBegin):addressEnd], 10, 64)
			if err != nil {
				panic(err)
			}

			number, err := strconv.ParseUint(inst[addressEnd+len(InstEnding):], 10, 64)
			if err != nil {
				panic(err)
			}

			for _, maskVariant := range maskVariants {
				// fmt.Println("address", strconv.FormatUint(address, 2), "maskVariant", strconv.FormatUint(maskVariant, 2))
				lAddress := (address & (^mask)) | maskOr | maskVariant
				memory[lAddress] = number
				// fmt.Println("Resulted address", strconv.FormatUint(lAddress, 2), "number", number)
			}
		}
	}

	var sum uint64
	for _, v := range memory {
		sum += v
	}

	return sum
}

func main() {
	reportLine, err := ioutil.ReadFile("adv14.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(reportLine), "\n")
	var partOne, partTwo uint64

	partOne = runProgram(lines)
	partTwo = runProgramV2(lines)

	fmt.Println("Part One:", partOne)
	fmt.Println("Part Two:", partTwo)
}
