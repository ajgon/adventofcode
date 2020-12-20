package day20

import (
	"reflect"
	"sort"
	"strings"
	"testing"
	"year2020/helpers"
)

var testData = strings.Split(`Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`, "\n")

var testTile = Tile{
	number: 2311,
	grid: []string{
		"..##.#..#.",
		"##..#.....",
		"#...##..#.",
		"####.#...#",
		"##.##.###.",
		"##...#.###",
		".#.#.#..##",
		"..#....#..",
		"###...#.#.",
		"..###..###",
	},
}

func TestExtractTiles(t *testing.T) {
	tiles := extractTiles(testData)
	got1 := tiles[0]
	want1 := testTile
	got2 := tiles[1]
	want2 := Tile{
		number: 1951,
		grid: []string{
			"#.##...##.",
			"#.####...#",
			".....#..##",
			"#...######",
			".##.#....#",
			".###.#####",
			"###.##.##.",
			".###....#.",
			"..#.#..#.#",
			"#...##.#..",
		},
	}
	got3 := len(tiles)
	want3 := 9

	if !reflect.DeepEqual([]Tile{got1, got2}, []Tile{want1, want2}) {
		t.Errorf("got %+v and %+v, want %+v and %+v", got1, got2, want1, want2)
	}

	if got3 != want3 {
		t.Errorf("got %d, want %d", got3, want3)
	}
}

func TestPossibleEdges(t *testing.T) {
	got := testTile.possibleEdges()
	want := []string{
		"..##.#..#.", "...#.##..#", "..###..###", ".#####..#.",
		".#..#.##..", "#..##.#...", "###..###..", ".#..#####.",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindUniqueEdges(t *testing.T) {
	tiles := extractTiles(testData)
	got := findUniqueEdges(tiles, false)
	want := []string{
		"##.#..#..#",
		"###....##.",
		"#.#.#####.",
		"..#.#....#",
		".#....####",
		".###..#...",
		"..###..###",
		"#...##.#..",
		".....##...",
		".#....#...",
		"##.#.#....",
		"#....####.",
	}

	sort.Strings(got)
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindCornerTiles(t *testing.T) {
	tiles := extractTiles(testData)
	got := findCornerTiles(tiles)
	want := []int{1951, 3079, 2971, 1171}

	sort.Ints(got)
	sort.Ints(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFlipTopBottom(t *testing.T) {
	got := testTile.flipTopBottom().grid
	want := []string{
		"..###..###",
		"###...#.#.",
		"..#....#..",
		".#.#.#..##",
		"##...#.###",
		"##.##.###.",
		"####.#...#",
		"#...##..#.",
		"##..#.....",
		"..##.#..#.",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFlipLeftRight(t *testing.T) {
	got := testTile.flipLeftRight().grid
	want := []string{
		".#..#.##..",
		".....#..##",
		".#..##...#",
		"#...#.####",
		".###.##.##",
		"###.#...##",
		"##..#.#.#.",
		"..#....#..",
		".#.#...###",
		"###..###..",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRotate(t *testing.T) {
	got := testTile.rotate(1).grid
	want := []string{
		".#..#####.",
		".#.####.#.",
		"###...#..#",
		"#..#.##..#",
		"#....#.##.",
		"...##.##.#",
		".#...#....",
		"#.#.##....",
		"##.###.#.#",
		"#..##.#...",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTopBorder(t *testing.T) {
	got := testTile.topBorder()
	want := "..##.#..#."

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRightBorder(t *testing.T) {
	got := testTile.rightBorder()
	want := "...#.##..#"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestBottomBorder(t *testing.T) {
	got := testTile.bottomBorder()
	want := "..###..###"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestLeftBorder(t *testing.T) {
	got := testTile.leftBorder()
	want := ".#####..#."

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCountNessies(t *testing.T) {
	tiles := extractTiles(testData)
	formattedPicture := arrangeTiles(tiles).Format()
	nessiesCount := 0

	for _, v := range formattedPicture.Variants() {
		nessiesCount = countNessies(v)

		if nessiesCount > 0 {
			break
		}
	}

	got := nessiesCount
	want := 2

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day20.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day20.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
