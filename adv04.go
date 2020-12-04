package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var validHCL = regexp.MustCompile(`\#[a-f0-9]{6,6}`)

func main() {
	input, err := ioutil.ReadFile("adv04.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")

	var recordFields []string
	var validRecordFields []string
	part1 := 0
	part2 := 0

	for i := range lines {
		line := lines[i]
		if len(line) == 0 {
			if len(recordFields) >= 7 {
				part1++
			}
			if len(validRecordFields) >= 7 {
				part2++
			}
			recordFields = make([]string, 0)
			validRecordFields = make([]string, 0)
			continue
		}
		fields := strings.Split(line, " ")
		for j := range fields {
			kv := strings.Split(fields[j], ":")
			if kv[0] != "cid" {
				recordFields = append(recordFields, kv[0])

				switch kv[0] {
				case "byr":
					if v, err := strconv.Atoi(kv[1]); err == nil && v >= 1920 && v <= 2002 {
						validRecordFields = append(validRecordFields, kv[0])
					}
				case "iyr":
					if v, err := strconv.Atoi(kv[1]); err == nil && v >= 2010 && v <= 2020 {
						validRecordFields = append(validRecordFields, kv[0])
					}
				case "eyr":
					if v, err := strconv.Atoi(kv[1]); err == nil && v >= 2020 && v <= 2030 {
						validRecordFields = append(validRecordFields, kv[0])
					}
				case "hgt":
					m := kv[1][len(kv[1])-2:]
					if v, err := strconv.Atoi(kv[1][:len(kv[1])-2]); err == nil {
						if (m == "cm" && v >= 150 && v <= 193) || (m == "in" && v >= 59 && v <= 76) {
							validRecordFields = append(validRecordFields, kv[0])
						}
					}
				case "hcl":
					if validHCL.MatchString(kv[1]) {
						validRecordFields = append(validRecordFields, kv[0])
					}
				case "ecl":
					if kv[1] == "amb" || kv[1] == "blu" || kv[1] == "brn" || kv[1] == "gry" || kv[1] == "grn" || kv[1] == "hzl" || kv[1] == "oth" {
						validRecordFields = append(validRecordFields, kv[0])
					}
				case "pid":
					if _, err := strconv.Atoi(kv[1]); len(kv[1]) == 9 && err == nil {
						validRecordFields = append(validRecordFields, kv[0])
					}
				}
			}
		}
	}
	fmt.Println("Part1", part1)
	fmt.Println("Part2", part2)
}
