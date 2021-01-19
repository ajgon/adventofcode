package day18

import (
	"regexp"
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day18.input")

func calculateSumsOnly(equation string) string {
	re := regexp.MustCompile(`(?:[0-9]+ \+ )+[0-9]+`)

	return re.ReplaceAllStringFunc(equation, calculateSimple)
}

func calculateSimple(equation string) string {
	var sign byte = '+'
	result := 0
	parsedValue := -1
	equation = equation + " "

	for i := 0; i < len(equation); i++ {
		switch equation[i] {
		case ' ':
			if parsedValue == -1 {
				continue
			}
			if sign == '+' {
				result += parsedValue
			} else {
				result *= parsedValue
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if parsedValue == -1 {
				parsedValue = int(equation[i] - 48)
			} else {
				parsedValue = 10*parsedValue + int(equation[i]-48)
			}
		case '+':
			parsedValue = -1
			sign = '+'
		case '*':
			parsedValue = -1
			sign = '*'
		}
	}

	return strconv.Itoa(result)
}

func calculateAdvanced(equation string) string {
	return calculateSimple(calculateSumsOnly(equation))
}

func calculateInBrackets(equation string, calculator func(string) string) (result int) {
	re := regexp.MustCompile(`\([^()]+\)`)

	for {
		if !strings.Contains(equation, "(") {
			break
		}
		equation = re.ReplaceAllStringFunc(equation, calculator)
	}

	equation = calculator(equation)
	result, _ = strconv.Atoi(equation)

	return
}

func SimpleSolution() string {
	var result int
	for i := 0; i < len(input); i++ {
		result += calculateInBrackets(input[i], calculateSimple)
	}

	return strconv.Itoa(result)
}

func AdvancedSolution() string {
	var result int
	for i := 0; i < len(input); i++ {
		result += calculateInBrackets(input[i], calculateAdvanced)
	}

	return strconv.Itoa(result)
}
