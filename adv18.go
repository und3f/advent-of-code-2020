package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	// "regexp"
	"strings"
)

type EvalateExpression func([]int, []rune) int

func evalate(expression string, evalateExpression EvalateExpression) (int, string) {
	// fmt.Println("evalating", expression)
	pos := 0

	operators := []rune{}
	arguments := []int{0}

	for pos = 0; pos < len(expression); pos++ {
		char := rune(expression[pos])
		if char >= '0' && char <= '9' {
			arguments[len(arguments)-1] = arguments[len(arguments)-1]*10 + int(char-'0')
		} else {
			switch char {
			case '+', '*':
				arguments = append(arguments, 0)
				operators = append(operators, char)
			case '(':
				var remainStr string
				arguments[len(arguments)-1], remainStr = evalate(expression[pos+1:], evalateExpression)
				// fmt.Println(len(arguments))
				pos = -1
				expression = remainStr
				continue
			case ')':
				return evalateExpression(arguments, operators), expression[pos+1:]
			case ' ':
			default:
				fmt.Errorf("Unexpected char %c\n", char)
			}
		}
	}
	return evalateExpression(arguments, operators), ""
}

func evalateBasic(arguments []int, operators []rune) int {
	// fmt.Println("evalateBasic", arguments, operators)
	// fmt.Printf("%d", arguments[0])
	result := arguments[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case '+':
			result += arguments[i+1]
		case '*':
			result *= arguments[i+1]
		}
		// fmt.Printf(" %s %d", string(operators[i]), arguments[i+1])
	}
	// fmt.Println("=", result)
	return result
}

func evalateOrder(_arguments []int, _operators []rune) int {
	// fmt.Println("evalateOrder", _arguments, _operators)

	arguments := []int{_arguments[0]}
	operators := []rune{}
	for i := 0; i < len(_operators); i++ {
		if _operators[i] == '+' {
			arguments[len(arguments)-1] += _arguments[i+1]
			// fmt.Println(arguments)
		} else {
			operators = append(operators, _operators[i])
			arguments = append(arguments, _arguments[i+1])
		}
	}

	return evalateBasic(arguments, operators)
}

func main() {
	reportLine, err := ioutil.ReadFile("adv18.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(reportLine)), "\n")

	fmt.Println(lines)
	var partOne, partTwo int
	for _, line := range lines {
		res, _ := evalate(line, evalateBasic)
		partOne += res
		res2, _ := evalate(line, evalateOrder)
		partTwo += res2
	}

	fmt.Println("Part One:", partOne)
	fmt.Println("Part Two:", partTwo)
}
