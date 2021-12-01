package day01

import (
	"strconv"

	"year2021/helpers"
)

func FindIncreasingEntriesCount(numbers []int, slide int) (uint16, error) {
	var count uint16

	for i := slide; i < len(numbers); i++ {
		var previousSum, currentSum int

		for j := i - 1; j > i-slide-1; j-- {
			previousSum += numbers[j]
		}

		for j := i; j > i-slide; j-- {
			currentSum += numbers[j]
		}

		if currentSum > previousSum {
			count++
		}
	}

	return count, nil
}

func SimpleSolution() string {
	input, _ := helpers.ProcessInput("day01.input")

	slideWindow := 1
	numbers := helpers.ConvertStringsToNumbers(input)
	solution, _ := FindIncreasingEntriesCount(numbers, slideWindow)

	return strconv.Itoa(int(solution))
}

func AdvancedSolution() string {
	input, _ := helpers.ProcessInput("day01.input")

	slideWindow := 3
	numbers := helpers.ConvertStringsToNumbers(input)
	solution, _ := FindIncreasingEntriesCount(numbers, slideWindow)

	return strconv.Itoa(int(solution))
}
