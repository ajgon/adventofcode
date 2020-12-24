package day24

import (
	"reflect"
	"strings"
	"testing"
	"year2020/helpers"
)

var testData = strings.Split(`sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`, "\n")

func TestExtractDirections(t *testing.T) {
	got := extractDirections("sesenwnenenewseeswwswswwnenewsewsw")

	want := []string{"se", "se", "nw", "ne", "ne", "ne", "w", "se", "e", "sw", "w", "sw", "sw", "w", "ne", "ne", "w", "se", "w", "sw"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestProcessLine(t *testing.T) {
	grid := make(map[int]map[int]bool)
	processLine(&grid, "nwwswee")

	got := grid
	want := map[int]map[int]bool{
		-2: map[int]bool{},
		-1: map[int]bool{},
		0: map[int]bool{
			0: true,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCalculateBlackTiles(t *testing.T) {
	grid := make(map[int]map[int]bool)
	processLine(&grid, "nwwswee")

	got := calculateBlackTiles(grid)
	want := 1

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestProcessLineSet(t *testing.T) {
	got := calculateBlackTiles(processLineSet(testData))
	want := 10

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountNeighbours(t *testing.T) {
	grid := map[int]map[int]bool{
		-1: map[int]bool{
			0: true,
		},
		0: map[int]bool{
			1: true,
		},
		1: map[int]bool{
			-1: true,
			0:  true,
		},
	}

	got := countNeighbours(grid, 0, 0)
	want := 4

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestEvolveTiles(t *testing.T) {
	grid := processLineSet(testData)

	got := calculateBlackTiles(evolveTiles(grid))
	want := 15

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day24.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day24.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
