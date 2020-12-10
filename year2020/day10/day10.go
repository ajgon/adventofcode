package day10

import (
	"sort"
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day10.input")

func findDifferences(data []string) []int {
	numbers := helpers.ConvertStringsToNumbers(data)
	sort.Ints(numbers)

	differences := make([]int, 1)
	differences[0] = numbers[0]

	for i := 0; i < len(numbers); i++ {
		if i+1 == len(numbers) {
			differences = append(differences, 3) // last socket
			break
		}

		differences = append(differences, numbers[i+1]-numbers[i])
	}

	return differences
}

func calculatePermutations(data []string) int {
	differences := findDifferences(data)
	countOnes := 0
	result := 1

	permutations := make([]int, 0)

	for _, d := range differences {
		if d == 3 {
			switch countOnes {
			case 1: // 1
				permutations = append(permutations, 1)
			case 2: // 11 or 2
				permutations = append(permutations, 2)
			case 3: // 111 or 12 or 21 or 3
				permutations = append(permutations, 4)
			case 4: // 1111 or 112 or 121 or 211 or 13 or 31 or 22
				permutations = append(permutations, 7)
			}
			countOnes = 0
			continue
		}

		if d == 1 {
			countOnes += 1
		}
	}

	for _, p := range permutations {
		result *= p
	}

	return result
}

func SimpleSolution() string {
	differences := findDifferences(input)
	countOnes := 0
	countThrees := 0

	for _, d := range differences {
		if d == 1 {
			countOnes += 1
			continue
		}

		if d == 3 {
			countThrees += 1
		}
	}

	return strconv.Itoa(countOnes * countThrees)
}

func AdvancedSolution() string {
	return strconv.Itoa(calculatePermutations(input))
}
