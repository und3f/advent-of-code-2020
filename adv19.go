package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	// "regexp"
	"strings"
)

func checkMessages(messages []string, rules []string) int {
	var result int

	for _, message := range messages {
		for _, rule := range rules {
			if message == rule {
				result++
				break
			}
		}
	}

	return result
}

type RuleDescription struct {
	index  int
	demand [][]int
}

func multiplySets(setA, setB []string) []string {
	if len(setA) == 0 {
		return setB
	} else if len(setB) == 0 {
		return setA
	}

	result := make([]string, len(setA)*len(setB))
	i := 0
	for _, a := range setA {
		for _, b := range setB {
			result[i] = a + b
			i++
		}
	}

	return result
}

func expandRule(rules [][]string, ruleDesc RuleDescription) []string {
	var rule []string
	for _, demand := range ruleDesc.demand {
		var rulePart []string
		for _, demandRuleI := range demand {
			demandedRules := rules[demandRuleI]
			if len(demandedRules) == 0 {
				return nil
			}
			rulePart = multiplySets(rulePart, demandedRules)
		}
		rule = append(rule, rulePart...)
	}

	// fmt.Println("return", rule)
	return rule
}

func buildRules(lines []string, patchRules []RuleDescription) [][]string {
	// var rules [][]string
	// fmt.Println(lines)
	rules := make([][]string, len(lines))

	var restRules []RuleDescription
	for _, line := range lines {
		ruleStr := strings.Split(line, ": ")

		ruleI, err := strconv.Atoi(ruleStr[0])
		if err != nil {
			panic(err)
		}
		if ruleStr[1][0] == '"' {
			rules[ruleI] = []string{ruleStr[1][1 : len(ruleStr[1])-1]}
		} else {
			ruleDemand := strings.Split(ruleStr[1], " ")
			var rule RuleDescription
			rule.index = ruleI
			rule.demand = [][]int{[]int{}}
			for _, v := range ruleDemand {
				if v == "|" {
					rule.demand = append([][]int{[]int{}}, rule.demand...)
				} else {
					n, _ := strconv.Atoi(v)
					rule.demand[0] = append(rule.demand[0], n)
				}
			}
			restRules = append(restRules, rule)
		}
	}

	for len(restRules) > 0 {
		var newRest []RuleDescription
		for _, v := range restRules {
			rule := expandRule(rules, v)
			if rule == nil {
				newRest = append(newRest, v)
				continue
			}

			rules[v.index] = rule
		}
		restRules = newRest
	}

	return rules
}

func main() {
	reportLine, err := ioutil.ReadFile("adv19.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(reportLine)), "\n")
	delimeterI := -1
	for i, l := range lines {
		if len(l) == 0 {
			delimeterI = i
			break
		}
	}

	messages := lines[delimeterI+1:]

	var partOne, partTwo int

	// fmt.Println(delimeterI)
	rules := buildRules(lines[:delimeterI], nil)
	partOne = checkMessages(messages, rules[0])

	rules = buildRules(lines[:delimeterI], []RuleDescription{
		RuleDescription{
			index:  8,
			demand: [][]int{[]int{42}, []int{42, 8}},
		},
		RuleDescription{
			index:  11,
			demand: [][]int{[]int{42, 31}, []int{42, 11, 31}},
		},
	})

	fmt.Println("Part One:", partOne)
	fmt.Println("Part Two:", partTwo)
}
