package day03

import (
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day03.input")

func countTrees(data []string, right int, down int) uint16 {
	var count uint16
	var dataWidth = len(data[0])
	var dataHeight = len(data)

	for x, y := 0, 0; y < dataHeight; x, y = x+right, y+down {
		if data[y][x%dataWidth] == '#' {
			count += 1
		}
	}

	return count
}

func SimpleSolution() string {
	return strconv.Itoa(int(countTrees(input, 3, 1)))
}

func AdvancedSolution() string {
	count := 1
	var rules = [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	for i := 0; i < len(rules); i++ {
		count *= int(countTrees(input, rules[i][0], rules[i][1]))
	}

	return strconv.Itoa(int(count))
}
