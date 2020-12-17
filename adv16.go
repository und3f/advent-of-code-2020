package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Rules map[string][]int
type Ticket []int

func parseTicket(input string) Ticket {
	fields := strings.Split(input, ",")
	ticket := make(Ticket, len(fields))
	for i, v := range fields {
		value, _ := strconv.Atoi(v)
		ticket[i] = value
	}

	return ticket
}

func parseInput(input []string) (Rules, []Ticket) {
	rules := make(Rules)
	tickets := make([]Ticket, 1)
	i := 0
	for ; len(input[i]) != 0; i++ {
		fields := strings.Split(input[i], ": ")
		name := fields[0]
		ranges := strings.Split(fields[1], " or ")
		rangeArr := make([]int, len(ranges)*2)
		for i, ruleRange := range ranges {
			minMax := strings.Split(ruleRange, "-")
			min, _ := strconv.Atoi(minMax[0])
			max, _ := strconv.Atoi(minMax[1])
			rangeArr[i*2] = min
			rangeArr[i*2+1] = max
		}
		rules[name] = rangeArr
	}

	i++
	i++
	tickets[0] = parseTicket(input[i])
	i++
	i++
	i++
	for ; i < len(input); i++ {
		tickets = append(tickets, parseTicket(input[i]))
	}

	return rules, tickets
}

type Fields []string

type PossibleFields [][][]string

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func recognizeFields(rules Rules, tickets []Ticket) Fields {
	fields := make(Fields, len(tickets[0]))

	possibleFields := make(PossibleFields, len(tickets[0]))

	for _, ticket := range tickets {
		for i, field := range ticket {
			var possibleRules []string
			for ruleName, rule := range rules {
				for i := 0; i < len(rule); i = i + 2 {
					if field >= rule[i] && field <= rule[i+1] {
						possibleRules = append(possibleRules, ruleName)
						break
					}
				}
			}

			possibleFields[i] = append(possibleFields[i], possibleRules)
		}
	}

	possibleFieldsUnrec := make([][]string, len(tickets[0]))
	for fieldI, possibilities := range possibleFields {
		possibleFields := possibilities[0]
		for _, fieldsB := range possibilities[1:] {
			var newFields []string
			for _, field := range possibleFields {
				if Contains(fieldsB, field) {
					newFields = append(newFields, field)
				}
			}
			possibleFields = newFields
		}

		var newFields []string
		for _, field := range possibleFields {
			if !Contains(fields, field) {
				newFields = append(newFields, field)
			}
		}
		possibleFields = newFields

		possibleFieldsUnrec[fieldI] = possibleFields
		if len(possibleFields) == 1 {
			fields[fieldI] = possibleFields[0]
		}
	}

	fullyRecognized := false
	for !fullyRecognized {
		fullyRecognized = true
		for fieldI, possibilities := range possibleFieldsUnrec {
			if len(possibilities) == 1 {
				continue
			}

			//fmt.Println(fieldI, possibilities)
			var newPossibilities []string
			for _, possibility := range possibilities {
				if !Contains(fields, possibility) {
					newPossibilities = append(newPossibilities, possibility)
				}
			}
			possibilities = newPossibilities
			if len(possibilities) == 1 {
				fields[fieldI] = possibilities[0]
			} else {
				fullyRecognized = false
			}
			possibleFieldsUnrec[fieldI] = possibilities
		}
	}

	//fmt.Println(possibleFieldsUnrec)
	//fmt.Println(fields)

	return fields
}

func part1(rules Rules, tickets []Ticket) (int, []Ticket) {
	var errorRate int
	var validTickets []Ticket

	for _, ticket := range tickets {
		validTicket := true
		for _, field := range ticket {
			validField := false
			for _, rule := range rules {
				for i := 0; i < len(rule); i = i + 2 {
					if field >= rule[i] && field <= rule[i+1] {
						validField = true
						break
					}
				}
				if validField {
					break
				}
			}
			if !validField {
				errorRate += field
				validTicket = false
				break
			}
		}
		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}

	return errorRate, validTickets
}

func main() {
	reportLine, err := ioutil.ReadFile("adv16.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(reportLine)), "\n")
	//fmt.Println(lines)
	rules, tickets := parseInput(lines)
	//fmt.Println(rules, tickets)

	partOne, validTickets := part1(rules, tickets[1:])

	fields := recognizeFields(rules, validTickets)
	myTicket := tickets[0]
	partTwo := 1
	for i, field := range fields {
		if strings.Index(field, "departure") != 0 {
			continue
		}
		partTwo = partTwo * myTicket[i]
	}

	fmt.Println("Part One:", partOne)
	fmt.Println("Part Two:", partTwo)
}
