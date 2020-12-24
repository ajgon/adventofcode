package day24

import (
	"math"
	"regexp"
	"strconv"
	"year2020/helpers"
)

var directionsMap = map[string][]int{
	"se": []int{0, 1},
	"sw": []int{-1, 1},
	"nw": []int{0, -1},
	"ne": []int{1, -1},
	"e":  []int{1, 0},
	"w":  []int{-1, 0},
}

var input, _ = helpers.ProcessInput("day24.input")

func extractDirections(line string) []string {
	re := regexp.MustCompile("(se|sw|nw|ne|e|w)")

	return re.FindAllString(line, -1)
}

func processLine(grid *map[int]map[int]bool, line string) {
	directions := extractDirections(line)
	pointer := []int{0, 0}

	for _, direction := range directions {
		pointer[0] += directionsMap[direction][0]
		pointer[1] += directionsMap[direction][1]

		if len((*grid)[pointer[0]]) == 0 {
			(*grid)[pointer[0]] = make(map[int]bool)
		}

	}
	(*grid)[pointer[0]][pointer[1]] = !(*grid)[pointer[0]][pointer[1]]
}

func calculateBlackTiles(grid map[int]map[int]bool) int {
	blackTiles := 0
	for _, sub := range grid {
		for _, tile := range sub {
			if tile {
				blackTiles += 1
			}
		}
	}

	return blackTiles
}

func processLineSet(data []string) map[int]map[int]bool {
	grid := make(map[int]map[int]bool)

	for _, line := range data {
		processLine(&grid, line)
	}

	return grid
}

func countNeighbours(grid map[int]map[int]bool, x, y int) int {
	neighbours := 0
	for _, dir := range directionsMap {
		if grid[x+dir[0]][y+dir[1]] {
			neighbours += 1
		}
	}

	return neighbours
}

func evolveTiles(grid map[int]map[int]bool) map[int]map[int]bool {
	newGrid := make(map[int]map[int]bool)
	var maxX, maxY float64 = 0, 0

	for x, sub := range grid {
		for y, _ := range sub {
			absX := math.Abs(float64(x))
			absY := math.Abs(float64(y))
			if absX > maxX {
				maxX = absX
			}
			if absY > maxY {
				maxY = absY
			}
		}
	}
	maxX += 2
	maxY += 2

	for x := -int(maxX); x < int(maxX); x++ {
		if len(newGrid[x]) == 0 {
			newGrid[x] = make(map[int]bool)
		}
		if len(grid[x]) == 0 {
			grid[x] = make(map[int]bool)
		}

		for y := -int(maxY); y < int(maxY); y++ {
			tile := grid[x][y]
			neighbours := countNeighbours(grid, x, y)
			newGrid[x][y] = tile
			if tile && (neighbours == 0 || neighbours > 2) {
				newGrid[x][y] = false
			}
			if !tile && neighbours == 2 {
				newGrid[x][y] = true
			}
		}
	}

	return newGrid
}

func SimpleSolution() string {
	return strconv.Itoa(calculateBlackTiles(processLineSet(input)))
}

func AdvancedSolution() string {
	grid := processLineSet(input)

	for i := 0; i < 100; i++ {
		grid = evolveTiles(grid)
	}

	return strconv.Itoa(calculateBlackTiles(grid))
}
