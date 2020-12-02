package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var report []int
	reportLine, err := ioutil.ReadFile("adv01.txt")
	if err != nil {
		panic(err)
	}
	reportLines := strings.Split(string(reportLine), "\n")
	for i := range reportLines {
		if value, err := strconv.Atoi(reportLines[i]); err == nil {
			report = append(report, value)
		}
	}
	fmt.Println(report)

	for i := range report {
		for j := i + 1; j < len(report); j++ {
			if expectedSum == report[i]+report[j] {
				fmt.Println("Part1 answer is", report[i]*report[j])
			}
			for k := j + 1; k < len(report); k++ {
				if expectedSum == report[i]+report[j]+report[k] {
					fmt.Println("Part2 answer is", report[i]*report[j]*report[k])
				}
			}
		}
	}
}

const expectedSum = 2020
