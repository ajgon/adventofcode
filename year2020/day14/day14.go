package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day14.input")

func applyMask(value int, mask string) int {
	replacer := strings.NewReplacer("1", "0", "X", "1")
	newMaskStr := replacer.Replace(mask)
	newMaskInt, _ := strconv.ParseInt(newMaskStr, 2, 64)
	value = value & int(newMaskInt)

	replacer = strings.NewReplacer("X", "0")
	newMaskStr = replacer.Replace(mask)
	newMaskInt, _ = strconv.ParseInt(newMaskStr, 2, 64)

	return value | int(newMaskInt)
}

func replaceFloating(combinations []string) []string {
	results := make([]string, 0)
	for _, combination := range combinations {
		parsedResults := make([]string, 0)
		if strings.Index(combination, "X") != -1 {
			parsedResults = append(parsedResults, strings.Replace(combination, "X", "0", 1))
			parsedResults = append(parsedResults, strings.Replace(combination, "X", "1", 1))

			newResults := replaceFloating(parsedResults)
			for n := 0; n < len(newResults); n++ {
				results = append(results, newResults[n])
			}
		} else {
			results = append(results, combination)
		}
	}

	return results
}

func applyFloatingMask(value int, mask string) []int {
	newValueArr := strings.Split(fmt.Sprintf("%036b", value), "")

	for v := 0; v < len(newValueArr); v++ {
		switch mask[v] {
		case '1':
			newValueArr[v] = "1"
		case 'X':
			newValueArr[v] = "X"
		}
	}

	combinations := make([]string, 0)
	combinations = append(combinations, strings.Join(newValueArr, ""))

	output := replaceFloating(combinations)
	results := make([]int, 0)

	for _, o := range output {
		number, _ := strconv.ParseInt(o, 2, 64)
		results = append(results, int(number))
	}

	return results
}

func processProgram(data []string, version int) map[int]int {
	var memory map[int]int = make(map[int]int, 0)
	memRegex := regexp.MustCompile("mem\\[([0-9]+)\\] = ([0-9]+)")
	mask := fmt.Sprintf("%036s", "0")

	for _, instruction := range data {
		if instruction[0:4] == "mask" {
			mask = instruction[7:]
		}

		if instruction[0:3] == "mem" {
			match := memRegex.FindStringSubmatch(instruction)
			address, _ := strconv.Atoi(match[1])
			value, _ := strconv.Atoi(match[2])

			if version == 1 {
				memory[address] = applyMask(value, mask)
				continue
			}

			addressPool := applyFloatingMask(address, mask)

			for _, addr := range addressPool {
				memory[addr] = value
			}
		}
	}

	return memory
}

func SimpleSolution() string {
	memory := processProgram(input, 1)
	result := 0

	for _, value := range memory {
		result += value
	}

	return strconv.Itoa(result)
}

func AdvancedSolution() string {
	memory := processProgram(input, 2)
	result := 0

	for _, value := range memory {
		result += value
	}

	return strconv.Itoa(result)
}
