package day20

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day20.input")

type FormattedPicture struct {
	grid []string
}

func (fp FormattedPicture) Draw() {
	for y := 0; y < len(fp.grid); y++ {
		fmt.Println(fp.grid[y])
	}
}

func (fp FormattedPicture) Variants() (serials [][]string) {
	return [][]string{
		fp.grid, flipTopBottom(fp.grid), flipLeftRight(fp.grid),
		flipTopBottom(flipLeftRight(fp.grid)), // rotate(2)
		rotate(fp.grid, 1), rotate(fp.grid, 3),
		flipLeftRight(rotate(fp.grid, 1)), flipLeftRight(rotate(fp.grid, 3)),
	}
}

func (fp FormattedPicture) Serialize() (serial string) {
	for y := 0; y < len(fp.grid); y++ {
		serial += fp.grid[y]
	}

	return
}

type Picture struct {
	tiles [][]Tile
}

func NewPicture(size int) Picture {
	picture := Picture{}

	for y := 0; y < size; y++ {
		row := make([]Tile, 0)
		for x := 0; x < size; x++ {
			row = append(row, NewTile())
		}
		picture.tiles = append(picture.tiles, row)
	}

	return picture
}

func (p Picture) Draw() {
	for y := 0; y < len(p.tiles); y++ {
		for t := 0; t < 10; t++ {
			for x := 0; x < len(p.tiles[y]); x++ {
				fmt.Print(p.tiles[y][x].grid[t] + " ")
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}

func (p Picture) Format() FormattedPicture {
	fp := FormattedPicture{}

	for y := 0; y < len(p.tiles); y++ {
		for t := 1; t < 9; t++ {
			row := ""
			for x := 0; x < len(p.tiles[y]); x++ {
				row += p.tiles[y][x].grid[t][1:9]
			}
			fp.grid = append(fp.grid, row)
		}
	}

	return fp
}

type Tile struct {
	number int
	grid   []string
}

func NewTile() Tile {
	tile := Tile{}

	for i := 0; i < 10; i++ {
		tile.grid = append(tile.grid, "          ")
	}

	return tile
}

func (t Tile) edges() (edges []string) {
	edges = append(edges, t.grid[0])

	edge := ""

	for y := 0; y < len(t.grid); y++ {
		edge += string(t.grid[y][9])
	}
	edges = append(edges, edge)

	edges = append(edges, t.grid[9])

	edge = ""

	for y := 0; y < len(t.grid); y++ {
		edge += string(t.grid[y][0])
	}
	edges = append(edges, edge)

	return
}

func (t Tile) possibleEdges() (edges []string) {
	edges = t.edges()

	for e := 0; e < 4; e++ {
		edges = append(edges, helpers.StringReverse(edges[e]))
	}

	return
}

func (t Tile) topBorder() string {
	return t.grid[0]
}

func (t Tile) rightBorder() string {
	border := ""
	for y := 0; y < len(t.grid); y++ {
		border += string(t.grid[y][len(t.grid)-1])
	}
	return border
}

func (t Tile) bottomBorder() string {
	return t.grid[len(t.grid)-1]
}

func (t Tile) leftBorder() string {
	border := ""
	for y := 0; y < len(t.grid); y++ {
		border += string(t.grid[y][0])
	}
	return border
}

func (t Tile) flipTopBottom() Tile {
	return Tile{number: t.number, grid: flipTopBottom(t.grid)}
}

func (t Tile) flipLeftRight() Tile {
	return Tile{number: t.number, grid: flipLeftRight(t.grid)}
}

func (t Tile) rotate(times int) Tile {
	return Tile{number: t.number, grid: rotate(t.grid, times)}
}

func flipTopBottom(grid []string) []string {
	newGrid := []string{}

	for y := len(grid) - 1; y >= 0; y-- {
		newGrid = append(newGrid, grid[y])
	}

	return newGrid
}

func flipLeftRight(grid []string) []string {
	newGrid := []string{}

	for y := 0; y < len(grid); y++ {
		newGrid = append(newGrid, helpers.StringReverse(grid[y]))
	}

	return newGrid
}

func rotate(grid []string, times int) []string {
	for n := 0; n < times; n++ {
		newGrid := make([]string, 0)
		for x := 0; x < len(grid[0]); x++ {
			row := ""
			for y := len(grid) - 1; y >= 0; y-- {
				row += string(grid[y][x])
			}
			newGrid = append(newGrid, row)
		}
		grid = newGrid
	}

	return grid
}

func extractTiles(data []string) (results []Tile) {
	tile := Tile{}

	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			results = append(results, tile)
			tile = Tile{}
			continue
		}

		if data[i][0:4] == "Tile" {
			tile.number, _ = strconv.Atoi(data[i][5:9])
			continue
		}

		tile.grid = append(tile.grid, data[i])
	}

	results = append(results, tile)

	return
}

func findUniqueEdges(tiles []Tile, full bool) (uniques []string) {
	edgesCount := make(map[string]int)

	for t := 0; t < len(tiles); t++ {
		possibleEdges := tiles[t].possibleEdges()

		for p := 0; p < len(possibleEdges); p++ {
			edgesCount[possibleEdges[p]] += 1
		}
	}

	for edge, count := range edgesCount {
		if count == 1 {
			for t := 0; t < len(tiles); t++ {
				if full || helpers.ArrayContainsStr(tiles[t].edges(), edge) {
					uniques = append(uniques, edge)
					break
				}
			}
		}
	}

	return
}

func findCornerTiles(tiles []Tile) (corners []int) {
	uniqueEdges := findUniqueEdges(tiles, false)
	tileEdgesCount := make(map[int]int)

	for t := 0; t < len(tiles); t++ {
		tileEdges := tiles[t].edges()

		for e := 0; e < len(tileEdges); e++ {
			if helpers.ArrayContainsStr(uniqueEdges, tileEdges[e]) {
				tileEdgesCount[tiles[t].number] += 1
			}
		}
	}

	for id, count := range tileEdgesCount {
		if count == 2 {
			corners = append(corners, id)
		}
	}

	return
}

func getTile(tiles []Tile, numbet int) Tile {
	for i := 0; i < len(tiles); i++ {
		if tiles[i].number == numbet {
			return tiles[i]
		}
	}

	panic("getTile")
}

func tileVariants(tile Tile) []Tile {
	return []Tile{
		tile, tile.flipTopBottom(), tile.flipLeftRight(),
		tile.flipTopBottom().flipLeftRight(), // rotate(2)
		tile.rotate(1), tile.rotate(3),
		tile.rotate(1).flipLeftRight(), tile.rotate(3).flipLeftRight(),
	}
}

func orientateFirstTile(tiles []Tile, uniqueEdges []string) Tile {
	cornerTiles := findCornerTiles(tiles)

	firstTile := getTile(tiles, cornerTiles[0])
	firstTileVariants := tileVariants(firstTile)

	for _, tile := range firstTileVariants {
		if helpers.ArrayContainsStr(uniqueEdges, tile.topBorder()) && helpers.ArrayContainsStr(uniqueEdges, tile.leftBorder()) {
			return tile
		}
	}

	panic("orientateFirstTile")
}

func findMatchingTileAgainstLeft(tiles []Tile, uniqueEdges []string, leftBorder string) Tile {
	for _, tile := range tiles {
		tileVariants := tileVariants(tile)
		for _, variant := range tileVariants {
			if variant.leftBorder() == leftBorder && helpers.ArrayContainsStr(uniqueEdges, variant.topBorder()) {
				return variant
			}
		}
	}

	panic("findMatchingTileAgainstLeft")
}

func findMatchingTileAgainstTop(tiles []Tile, uniqueEdges []string, topBorder string) Tile {
	for _, tile := range tiles {
		tileVariants := tileVariants(tile)
		for _, variant := range tileVariants {
			if variant.topBorder() == topBorder && helpers.ArrayContainsStr(uniqueEdges, variant.leftBorder()) {
				return variant
			}
		}
	}

	panic("findMatchingTileAgainstTop")
}

func findMatchingTileAgainstTopLeft(tiles []Tile, topBorder, leftBorder string) Tile {
	for _, tile := range tiles {
		tileVariants := tileVariants(tile)
		for _, variant := range tileVariants {
			if variant.topBorder() == topBorder && variant.leftBorder() == leftBorder {
				return variant
			}
		}
	}

	panic("findMatchingTileAgainstTopLeft")
}

func removeTileFromPack(tiles []Tile, tile Tile) []Tile {
	newTiles := make([]Tile, 0)

	for _, t := range tiles {
		if t.number != tile.number {
			newTiles = append(newTiles, t)
		}
	}

	return newTiles
}

func arrangeTiles(tiles []Tile) Picture {
	picture := NewPicture(int(math.Sqrt(float64(len(tiles)))))
	uniqueEdges := findUniqueEdges(tiles, true)

	picture.tiles[0][0] = orientateFirstTile(tiles, uniqueEdges)
	tiles = removeTileFromPack(tiles, picture.tiles[0][0])

	for p := 1; p < len(picture.tiles[0]); p++ {
		picture.tiles[0][p] = findMatchingTileAgainstLeft(tiles, uniqueEdges, picture.tiles[0][p-1].rightBorder())
		tiles = removeTileFromPack(tiles, picture.tiles[0][p])
	}

	for y := 1; y < len(picture.tiles); y++ {
		picture.tiles[y][0] = findMatchingTileAgainstTop(tiles, uniqueEdges, picture.tiles[y-1][0].bottomBorder())
		tiles = removeTileFromPack(tiles, picture.tiles[y][0])
		for x := 1; x < len(picture.tiles[y]); x++ {
			picture.tiles[y][x] = findMatchingTileAgainstTopLeft(tiles, picture.tiles[y-1][x].bottomBorder(), picture.tiles[y][x-1].rightBorder())
			tiles = removeTileFromPack(tiles, picture.tiles[y][x])
		}
	}

	return picture
}

func countNessies(photo []string) (count int) {
	nessieRegex := regexp.MustCompile(strings.Join([]string{
		"..................#.",
		"#....##....##....###",
		".#..#..#..#..#..#...",
	}, ""))

	// nessie size 20x3
	for y := 0; y < len(photo)-3; y++ {
		for x := 0; x < len(photo[y])-20; x++ {
			scan := photo[y][x:(x+20)] + photo[y+1][x:(x+20)] + photo[y+2][x:(x+20)]
			if nessieRegex.MatchString(scan) {
				count += 1
			}
		}
	}

	return count
}

func SimpleSolution() string {
	result := 1
	tileNumbers := findCornerTiles(extractTiles(input))

	for t := 0; t < len(tileNumbers); t++ {
		result *= tileNumbers[t]
	}

	return strconv.Itoa(result)
}

func AdvancedSolution() string {
	tiles := extractTiles(input)
	formattedPicture := arrangeTiles(tiles).Format()
	serializedPicture := formattedPicture.Serialize()

	nessiesCount := 0
	hashesCount := 0

	for c := 0; c < len(serializedPicture); c++ {
		if serializedPicture[c] == '#' {
			hashesCount += 1
		}
	}

	for _, v := range formattedPicture.Variants() {
		nessiesCount = countNessies(v)

		if nessiesCount > 0 {
			break
		}
	}

	return strconv.Itoa(hashesCount - nessiesCount*15)
}
