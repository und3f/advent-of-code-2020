package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	// "regexp"
	"strings"
)

func checkMessages(messages []string, rule Rule) int {
	var result int

	for _, message := range messages {
		if rule.match(message) {
			result++
		}
	}

	return result
}

func composePartTwoRule(rule42, rule31 *StringRule, maxLength int) []string {
	rule0Str := "0: 8 11"
	rule8Str := "8: 42"
	rule11Str := "11: 42 31"

	minL42 := rule42.minRuleLength()
	minL31 := rule31.minRuleLength()

	minRulePartTwo := minL42 + minL31
	for i := 2; i*minL42+minRulePartTwo <= maxLength; i++ {
		rule8Str += " |"
		for j := 1; j <= i; j++ {
			rule8Str += " 42"
		}
	}

	for i := 2; minL42+i*minRulePartTwo <= maxLength; i++ {
		rule11Str += " |"
		ruleSufix := ""
		for j := 1; j <= i; j++ {
			rule11Str += " 42"
			ruleSufix += " 31"
		}
		rule11Str += ruleSufix

	}

	return []string{rule8Str, rule11Str, rule0Str}
}

func checkPartTwoMessages(messages []string, rule42, rule31 *StringRule) int {
	var result int

	for _, message := range messages {
		if found, count := matchRule42(message, rule42, rule31); found && count > 0 {
			result++
			// fmt.Println(message)
		}
	}

	return result
}

func matchRule42(message string, rule42, rule31 *StringRule) (bool, int) {
	for _, rule := range rule42.rules {
		if strings.HasPrefix(message, rule) {
			found, submatchcount := matchRule42(message[len(rule):], rule42, rule31)
			if found {
				return true, 1 + submatchcount
			}
		}

		submatch := matchRule31(message, rule31)
		if submatch > 0 {
			return true, -submatch
		}
	}

	return false, 0
}

func matchRule31(message string, rule31 *StringRule) int {
	if len(message) == 0 {
		return 0
	}

	for _, rule := range rule31.rules {
		if strings.HasPrefix(message, rule) {
			substr := matchRule31(message[len(rule):], rule31)
			if substr >= 0 {
				return 1 + substr
			}
		}
	}
	return -1
}

type RuleDescription struct {
	index  int
	demand [][]int
}

type Rule interface {
	match(string) bool
	matchPartially(string) (bool, string)
}

type StringRule struct {
	rules []string
}

func (this StringRule) match(message string) bool {
	for _, rule := range this.rules {
		if message == rule {
			return true
		}
	}
	return false
}

func (this StringRule) minRuleLength() int {
	minLength := len(this.rules[0])
	for _, rule := range this.rules {
		if len(rule) < minLength {
			minLength = len(rule)
		}
	}
	return minLength
}

func (this StringRule) matchPartially(message string) (bool, string) {
	for _, rule := range this.rules {
		if strings.HasPrefix(message, rule) {
			return true, message[len(rule):]
		}
	}
	return false, ""
}

func newStringRule(rules []string) *StringRule {
	rule := new(StringRule)
	rule.rules = rules
	return rule
}

type RuleSet map[int]Rule
type StringRuleSet map[int]*StringRule

type Rule8 struct{ rule42 *StringRule }

func (this Rule8) match(message string) bool {
	matched, rest := this.matchPartially(message)
	if matched {
		return len(rest) == 0
	}
	return false
}

func (this Rule8) matchPartially(message string) (bool, string) {
	rest := message
	for matched := true; matched; {
		matched, rest = this.rule42.matchPartially(rest)
		if len(rest) == 0 {
			return true, ""
		}
	}

	return len(rest) < len(message), rest
}

func newRule8(rules StringRuleSet) *Rule8 {
	rule := new(Rule8)
	rule.rule42 = rules[42]
	return rule
}

type Rule11 struct{ rule42, rule31 *StringRule }

func (Rule11) match(message string) bool {
	return false
}
func (Rule11) matchPartially(message string) (bool, string) {
	return false, ""
}

func newRule11(rules StringRuleSet) *Rule11 {
	rule := new(Rule11)
	rule.rule42 = rules[42]
	rule.rule31 = rules[31]
	return rule
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

func expandRule(rules map[int][]string, ruleDesc RuleDescription) []string {
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

func buildStringRules(lines []string) StringRuleSet {
	// var rules [][]string
	// fmt.Println(lines)
	rules := make(map[int][]string)

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

	prevRestRulesLen := len(restRules) + 1
	for len(restRules) > 0 && prevRestRulesLen > len(restRules) {
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

	rulesObj := make(StringRuleSet)
	for key, rule := range rules {
		rulesObj[key] = newStringRule(rule)
	}
	return rulesObj
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

	rulesLines := lines[:delimeterI]
	messages := lines[delimeterI+1:]

	var partOne, partTwo int

	// fmt.Println(delimeterI)
	rules := buildStringRules(rulesLines)
	partOne = checkMessages(messages, rules[0])
	fmt.Println("Part One:", partOne)

	// Part Two
	/*
		var modifiedRules []string
		for i, line := range rulesLines {
			if strings.HasPrefix(line, "0:") || strings.HasPrefix(line, "8:") || strings.HasPrefix(line, "11:") {
				modifiedRules = append(rulesLines[:i], rulesLines[i+1:]...)
				break
			}
		}

		maxLength := 0
		for _, line := range messages {
			if len(line) > maxLength {
				maxLength = len(line)
			}
		}

		fmt.Println(rules[0])
	*/

	partTwo = checkPartTwoMessages(messages, rules[42], rules[31])
	fmt.Println("Part Two:", partTwo)
}
