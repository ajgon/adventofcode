package day01

import (
	"errors"
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day01.input")

func findNEntriesWithSum(n uint8, expectedSum uint16) ([]uint16, error) {
	var numbers []uint16

	for _, item := range input {
		number, _ := strconv.ParseInt(item, 10, 64)
		numbers = append(numbers, uint16(number))
	}

	var group []uint16
	var sum uint16
	limit := len(input)

	positions := make([]uint8, n-1)

	for alter := false; !alter; {
		alter = true
		for p := 0; p < len(positions); p++ {
			for i := 0; i < len(numbers); i++ {
				group = make([]uint16, 0)
				group = append(group, numbers[i])
				sum = 0

				for pp := 0; pp < len(positions); pp++ {
					group = append(group, numbers[positions[pp]])
				}

				for _, term := range group {
					sum += term
				}

				if sum == expectedSum {
					return group, nil
				}
			}

			if alter {
				positions[p] += 1
				alter = positions[p] == uint8(limit)
				if alter {
					positions[p] = 0
				}
			}
		}
	}

	group = make([]uint16, 0)

	return group, errors.New("Sum not found")
}

func SimpleSolution() string {
	var solution uint = 1
	group, _ := findNEntriesWithSum(2, 2020)

	for _, term := range group {
		solution *= uint(term)
	}

	return strconv.Itoa(int(solution))
}

func AdvancedSolution() string {
	var solution uint = 1
	group, _ := findNEntriesWithSum(3, 2020)

	for _, term := range group {
		solution *= uint(term)
	}

	return strconv.Itoa(int(solution))
}
