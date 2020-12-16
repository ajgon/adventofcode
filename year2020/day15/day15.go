package day15

import (
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day15.input", ",")

func calculateNthIteration(data []int, n uint32) int {
	var seen [][]uint32 = make([][]uint32, n)
	var spoken uint32 = 0
	var dataLen = uint32(len(data))

	for i := uint32(0); i < uint32(dataLen); i++ {
		seen[uint32(data[i])] = make([]uint32, 1)
		seen[uint32(data[i])] = append(seen[uint32(data[i])], i)
	}

	for i := uint32(0); i < n-dataLen; i++ {
		lastSeen := seen[spoken]
		if len(lastSeen) == 2 {
			spoken = lastSeen[1] - lastSeen[0]
		} else {
			spoken = 0
		}

		if len(seen[spoken]) == 0 {
			seen[spoken] = []uint32{dataLen + i}
		} else {
			seen[spoken] = []uint32{seen[spoken][len(seen[spoken])-1], dataLen + i}
		}
	}

	return int(spoken)
}

func SimpleSolution() string {
	data := helpers.ConvertStringsToNumbers(input)

	return strconv.Itoa(calculateNthIteration(data, 2020))
}

func AdvancedSolution() string {
	data := helpers.ConvertStringsToNumbers(input)

	return strconv.Itoa(calculateNthIteration(data, 30000000))
}
