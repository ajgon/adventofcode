package day11

import (
	"reflect"
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day11.input", "nosplit")

func getMap(data string) [][]byte {
	rows := strings.Split(data, "\n")
	results := make([][]byte, 0)

	for _, r := range rows {
		row := make([]byte, 0)

		for i := 0; i < len(r); i++ {
			row = append(row, r[i])
		}

		results = append(results, row)
	}

	return results
}

func inDimensions(data [][]byte, x, y int) bool {
	return x >= 0 && x < len(data[0]) && y >= 0 && y < len(data)
}

func getNearby(data [][]byte, x, y, xDirection, yDirection int) byte {
	if !inDimensions(data, x+xDirection, y+yDirection) {
		return '.'
	}

	return data[y+yDirection][x+xDirection]
}

func getRay(data [][]byte, x, y, xDirection, yDirection int) byte {
	for {
		x += xDirection
		y += yDirection
		if !inDimensions(data, x, y) {
			return '.'
		}
		if data[y][x] != '.' {
			return data[y][x]
		}
	}
}

func getOccupiedSeats(data [][]byte, x, y int, nearby bool) int {
	directions := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	seats := 0

	for _, direction := range directions {
		var seat byte
		if nearby {
			seat = getNearby(data, x, y, direction[0], direction[1])
			if seat == '#' {
				seats += 1
			}
		} else {
			seat = getRay(data, x, y, direction[0], direction[1])
			if seat == '#' {
				seats += 1
			}
		}
	}

	return seats
}

func calculateSeats(data [][]byte) int {
	seats := 0

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {
			if data[y][x] == '#' {
				seats += 1
			}
		}
	}

	return seats
}

func regorganizePeople(data [][]byte, nearby bool, limit int) [][]byte {
	newData := make([][]byte, 0)
	for i := 0; i < len(data); i++ {
		newData = append(newData, make([]byte, len(data[i])))
	}

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			occupiedSeats := getOccupiedSeats(data, x, y, nearby)

			if data[y][x] != '.' && occupiedSeats >= limit {
				newData[y][x] = 'L'
				continue
			}

			if data[y][x] != '.' && occupiedSeats == 0 {
				newData[y][x] = '#'
				continue
			}

			newData[y][x] = data[y][x]
		}
	}

	return newData
}

func SimpleSolution() string {
	data := getMap(input[0])

	for {
		newData := regorganizePeople(data, true, 4)
		if reflect.DeepEqual(data, newData) {
			break
		}
		data = newData
	}

	seats := calculateSeats(data)

	return strconv.Itoa(seats)
}

func AdvancedSolution() string {
	data := getMap(input[0])

	for {
		newData := regorganizePeople(data, false, 5)
		if reflect.DeepEqual(data, newData) {
			break
		}
		data = newData
	}

	seats := calculateSeats(data)

	return strconv.Itoa(seats)
}
