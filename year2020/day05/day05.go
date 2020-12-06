package day05

import (
	"sort"
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day05.input")

func calculateID(seat string) int {
	var rules []string = strings.Split(seat, "")
	down, up, left, right := 0, 127, 0, 7

	for _, c := range rules {
		switch c {
		case "F":
			up -= (up - down + 1) / 2
		case "B":
			down += (up - down + 1) / 2
		case "L":
			right -= (right - left + 1) / 2
		case "R":
			left += (right - left + 1) / 2
		}
	}

	return down*8 + left
}

func SimpleSolution() string {
	highestID := 0

	for _, data := range input {
		newID := calculateID(data)
		if newID > highestID {
			highestID = newID
		}
	}

	return strconv.Itoa(highestID)
}

func AdvancedSolution() string {
	var IDs []int

	for _, data := range input {
		IDs = append(IDs, calculateID(data))
	}

	sort.Ints(IDs)

	for i := 1; i < len(IDs)-1; i++ {
		if !(IDs[i-1] == IDs[i]-1 && IDs[i+1] == IDs[i]+1) {
			return strconv.Itoa(IDs[i] + 1)
		}
	}

	return ""
}
