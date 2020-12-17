package day16

import (
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day16.input")

func extractTicketsData(data []string) (rules map[string][]int, ticket []int, nearby [][]int) {
	step := 0
	rules = make(map[string][]int)
	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			step += 1
			continue
		}

		switch step {
		case 0: // rules
			nameValues := strings.Split(data[i], ": ")
			rangeNames := strings.Split(nameValues[1], " or ")
			rulesNumbers := make([]int, 0)
			for _, rn := range rangeNames {
				upNameDownName := strings.Split(rn, "-")
				up, _ := strconv.Atoi(upNameDownName[0])
				down, _ := strconv.Atoi(upNameDownName[1])

				for j := up; j <= down; j++ {
					rulesNumbers = append(rulesNumbers, j)
				}
			}
			rules[nameValues[0]] = rulesNumbers
		case 1: // ticket
			ticket = helpers.ConvertStringsToNumbers(strings.Split(data[i], ","))
		case 2: // nearby
			parsedData := helpers.ConvertStringsToNumbers(strings.Split(data[i], ","))
			if len(parsedData) > 1 {
				nearby = append(nearby, parsedData)
			}
		}
	}

	return
}

func rulesInclude(rules map[string][]int, number int) bool {
	for _, ruleNumbers := range rules {
		if helpers.ArrayContainsInt(ruleNumbers, number) {
			return true
		}
	}

	return false
}

func findInvalidTicketValues(rules map[string][]int, nearby [][]int) []int {
	invalidValues := make([]int, 0)

	for _, row := range nearby {
		for _, value := range row {
			if !rulesInclude(rules, value) {
				invalidValues = append(invalidValues, value)
			}
		}
	}

	return invalidValues
}

func findValidNearbyTickets(rules map[string][]int, nearby [][]int) [][]int {
	validNearbyTickets := make([][]int, 0)
	var valid bool

	for _, row := range nearby {
		valid = true
		for _, value := range row {
			valid = valid && rulesInclude(rules, value)
		}

		if valid {
			validNearbyTickets = append(validNearbyTickets, row)
		}
	}

	return validNearbyTickets
}

func transpose(matrix [][]int) (transposed [][]int) {
	for x := 0; x < len(matrix[0]); x++ {
		transposed = append(transposed, make([]int, len(matrix)))
		for y := 0; y < len(matrix); y++ {
			transposed[x][y] = matrix[y][x]
		}
	}
	return
}

func determineApplicableNamesFromValues(rules map[string][]int, values []int) (applicableNames []string) {
	var valid bool

	for name, ruleValues := range rules {
		valid = true
		for v := 0; v < len(values); v++ {
			valid = valid && helpers.ArrayContainsInt(ruleValues, values[v])
		}

		if valid {
			applicableNames = append(applicableNames, name)
		}
	}

	return applicableNames
}

func removeFromArrayStr(stack []string, needle string) []string {
	cleanArray := make([]string, 0)

	for i := 0; i < len(stack); i++ {
		if stack[i] != needle {
			cleanArray = append(cleanArray, stack[i])
		}
	}

	return cleanArray
}

func narrowPossibilities(possibilities [][]string) []string {
	results := make([]string, len(possibilities))
	var empty bool

	for {
		empty = true
		for p := 0; p < len(possibilities); p++ {
			if len(possibilities[p]) == 1 {
				item := possibilities[p][0]
				results[p] = item
				for pp := 0; pp < len(possibilities); pp++ {
					possibilities[pp] = removeFromArrayStr(possibilities[pp], item)
				}
				empty = false
			}
		}

		if empty {
			break
		}
	}

	return results
}

func SimpleSolution() string {
	rules, _, nearby := extractTicketsData(input)
	data := findInvalidTicketValues(rules, nearby)
	result := 0

	for i := 0; i < len(data); i++ {
		result += data[i]
	}

	return strconv.Itoa(result)
}

func AdvancedSolution() string {
	rules, ticket, nearby := extractTicketsData(input)
	validNearby := findValidNearbyTickets(rules, nearby)
	transposedNearby := transpose(validNearby)
	applicableNames := make([][]string, 0)

	for i := 0; i < len(transposedNearby); i++ {
		applicableNames = append(applicableNames, determineApplicableNamesFromValues(rules, transposedNearby[i]))
	}

	finalOrder := narrowPossibilities(applicableNames)
	result := 1

	for i := 0; i < len(finalOrder); i++ {
		if len(finalOrder[i]) >= 9 && finalOrder[i][0:9] == "departure" {
			result *= ticket[i]
		}
	}

	return strconv.Itoa(result)
}
