package day17

import (
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day17.input")

func countNeighbours(dimension [][][][]string, posW, posZ, posY, posX int) int {
	count := 0

	for w := posW - 1; w <= posW+1; w++ {
		for z := posZ - 1; z <= posZ+1; z++ {
			for y := posY - 1; y <= posY+1; y++ {
				for x := posX - 1; x <= posX+1; x++ {
					if (x == posX && y == posY && z == posZ && w == posW) || x < 0 || y < 0 || z < 0 || w < 0 {
						continue
					}

					if w < len(dimension) && z < len(dimension[w]) && y < len(dimension[w][z]) && x < len(dimension[w][z][y]) && dimension[w][z][y][x] == "#" {
						count += 1
					}
				}
			}
		}
	}

	return count
}

func initializeGrid(data []string) (output [][][][]string) {
	output = emptyGrid(1, 1, len(data), len(data[0]))

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			output[0][0][y][x] = string(data[y][x])
		}
	}

	return
}

func emptyGrid(wMax, zMax, yMax, xMax int) (output [][][][]string) {
	for w := 0; w < wMax; w++ {
		wRow := make([][][]string, 0)
		for z := 0; z < zMax; z++ {
			zRow := make([][]string, 0)
			for y := 0; y < yMax; y++ {
				yRow := make([]string, 0)
				for x := 0; x < xMax; x++ {
					yRow = append(yRow, ".")
				}
				zRow = append(zRow, yRow)
			}
			wRow = append(wRow, zRow)
		}
		output = append(output, wRow)
	}

	return output
}

func normalizeGrid(dimension [][][][]string) (normalized [][][][]string) {
	normalized = emptyGrid(len(dimension)+2, len(dimension[0])+2, len(dimension[0][0])+2, len(dimension[0][0][0])+2)

	for w := 0; w < len(dimension)+2; w++ {
		for z := 0; z < len(dimension[0])+2; z++ {
			for y := 0; y < len(dimension[0][0])+2; y++ {
				for x := 0; x < len(dimension[0][0][0])+2; x++ {
					if x > 0 && y > 0 && z > 0 && w > 0 && x < len(dimension[0][0][0])+1 && y < len(dimension[0][0])+1 && z < len(dimension[0])+1 && w < len(dimension)+1 {
						normalized[w][z][y][x] = dimension[w-1][z-1][y-1][x-1]
					}
				}
			}
		}
	}

	return
}

func nextGeneration(dimension [][][][]string, include4D bool) (generation [][][][]string) {
	dimension = normalizeGrid(dimension)
	generation = emptyGrid(len(dimension), len(dimension[0]), len(dimension[0][0]), len(dimension[0][0][0]))

	for w := 0; w < len(dimension); w++ {
		for z := 0; z < len(dimension[0]); z++ {
			for y := 0; y < len(dimension[0][0]); y++ {
				for x := 0; x < len(dimension[0][0][0]); x++ {
					if w != len(dimension)/2 && !include4D {
						continue
					}
					count := countNeighbours(dimension, w, z, y, x)
					generation[w][z][y][x] = dimension[w][z][y][x]
					if dimension[w][z][y][x] == "#" && (count > 3 || count < 2) {
						generation[w][z][y][x] = "."
					} else if dimension[w][z][y][x] == "." && count == 3 {
						generation[w][z][y][x] = "#"
					}
				}
			}
		}
	}

	return
}

func bootProcessBlocksCount(generation [][][][]string, include4D bool) (count int) {
	for i := 0; i < 6; i++ {
		generation = nextGeneration(generation, include4D)
	}

	for w := 0; w < len(generation); w++ {
		for z := 0; z < len(generation[0]); z++ {
			for y := 0; y < len(generation[0][0]); y++ {
				for x := 0; x < len(generation[0][0][0]); x++ {
					if generation[w][z][y][x] == "#" {
						count += 1
					}
				}
			}
		}
	}

	return
}

func SimpleSolution() string {
	genesis := initializeGrid(input)
	return strconv.Itoa(bootProcessBlocksCount(genesis, false))
}

func AdvancedSolution() string {
	genesis := initializeGrid(input)
	return strconv.Itoa(bootProcessBlocksCount(genesis, true))
}
