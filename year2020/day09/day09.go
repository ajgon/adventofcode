package day09

import (
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day09.input")

func containsSum(stack []int, needle int) bool {
	for i := 0; i < len(stack); i++ {
		for j := i + 1; j < len(stack); j++ {
			if stack[i]+stack[j] == needle {
				return true
			}
		}
	}

	return false
}

func findMissingNumber(data []string, preamble int) int {
	numbers := helpers.ConvertStringsToNumbers(data)

	for i := preamble; i < len(data); i++ {
		chunk := numbers[(i - preamble):i]
		if !containsSum(chunk, numbers[i]) {
			return numbers[i]
		}
	}

	return -1
}

func findEncryptionWeakness(data []string, needle int) int {
	numbers := helpers.ConvertStringsToNumbers(data)

	for size := 2; size < len(numbers); size++ {
		for i := 0; i < len(numbers)-size; i++ {
			chunk := numbers[i : size+i]
			sum := 0

			for s := 0; s < len(chunk); s++ {
				sum += chunk[s]
			}

			if sum == needle {
				min, max := chunk[0], chunk[0]
				for m := 0; m < len(chunk); m++ {
					if min > chunk[m] {
						min = chunk[m]
					}
					if max < chunk[m] {
						max = chunk[m]
					}
				}
				return min + max
			}
		}
	}

	return -1
}

func SimpleSolution() string {
	return strconv.Itoa(findMissingNumber(input, 25))
}

func AdvancedSolution() string {
	needle := findMissingNumber(input, 25)

	return strconv.Itoa(findEncryptionWeakness(input, needle))
}

//func AdvancedSolution() string {
//}
