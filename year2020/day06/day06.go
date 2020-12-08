package day06

import (
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day06.input", "\n\n")

func simpleCountAnswers(answers string) int {
	var lines []string = strings.Split(answers, "\n")
	var rules []string

	for _, line := range lines {
		letters := strings.Split(line, "")

		for _, letter := range letters {
			if !helpers.ArrayContainsStr(rules, letter) {
				rules = append(rules, letter)
			}
		}
	}

	return len(rules)
}

func advancedCountAnswers(answers string) int {
	var lines []string = strings.Split(answers, "\n")
	var rules []string = strings.Split(lines[0], "")

	for _, line := range lines {
		var newRules []string
		letters := strings.Split(line, "")

		for _, letter := range letters {
			if helpers.ArrayContainsStr(rules, letter) {
				newRules = append(newRules, letter)
			}
		}

		rules = newRules
	}

	return len(rules)
}

func SimpleSolution() string {
	count := 0

	for _, data := range input {
		count += simpleCountAnswers(data)
	}

	return strconv.Itoa(count)
}

func AdvancedSolution() string {
	count := 0

	for _, data := range input {
		count += advancedCountAnswers(data)
	}

	return strconv.Itoa(count)
}
