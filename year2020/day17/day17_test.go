package day17

import (
	"reflect"
	"strings"
	"testing"
	"year2020/helpers"
)

var testData []string = strings.Split(`.#.
..#
###`, "\n")

func TestInitializeGrid3D(t *testing.T) {
	got := initializeGrid(testData)
	want := [][][][]string{
		[][][]string{
			[][]string{[]string{".", "#", "."}, []string{".", ".", "#"}, []string{"#", "#", "#"}},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCountNeighbours3D01(t *testing.T) {
	got := countNeighbours([][][][]string{
		[][][]string{
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", "#", "#"}},
			[][]string{[]string{"#", "#", "#"}, []string{"#", "X", "#"}, []string{"#", "#", "#"}},
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", "#", "#"}},
		},
	}, 0, 1, 1, 1)
	want := 26

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountNeighbours3D02(t *testing.T) {
	got := countNeighbours([][][][]string{
		[][][]string{
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "."}, []string{"#", "#", "#"}},
			[][]string{[]string{"#", ".", "X"}, []string{"#", "#", "#"}, []string{"#", ".", "#"}},
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", ".", "#"}},
		}}, 0, 1, 0, 2)
	want := 9

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestNormalizeGrid3D(t *testing.T) {
	got := normalizeGrid([][][][]string{
		[][][]string{
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}},
		},
	})[1]
	want := [][][]string{
		[][]string{[]string{".", ".", ".", ".", "."}, []string{".", ".", ".", ".", "."}, []string{".", ".", ".", ".", "."}, []string{".", ".", ".", ".", "."}},
		[][]string{[]string{".", ".", ".", ".", "."}, []string{".", "#", "#", "#", "."}, []string{".", "#", "#", "#", "."}, []string{".", ".", ".", ".", "."}},
		[][]string{[]string{".", ".", ".", ".", "."}, []string{".", ".", ".", ".", "."}, []string{".", ".", ".", ".", "."}, []string{".", ".", ".", ".", "."}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNextGeneration3D(t *testing.T) {
	got := nextGeneration([][][][]string{
		[][][]string{
			[][]string{
				[]string{".", "#", "."},
				[]string{".", ".", "#"},
				[]string{"#", "#", "#"},
			},
		},
	}, false)[1]
	want := [][][]string{
		[][]string{
			[]string{".", ".", ".", ".", "."},
			[]string{".", ".", ".", ".", "."},
			[]string{".", "#", ".", ".", "."},
			[]string{".", ".", ".", "#", "."},
			[]string{".", ".", "#", ".", "."},
		},
		[][]string{
			[]string{".", ".", ".", ".", "."},
			[]string{".", ".", ".", ".", "."},
			[]string{".", "#", ".", "#", "."},
			[]string{".", ".", "#", "#", "."},
			[]string{".", ".", "#", ".", "."},
		},
		[][]string{
			[]string{".", ".", ".", ".", "."},
			[]string{".", ".", ".", ".", "."},
			[]string{".", "#", ".", ".", "."},
			[]string{".", ".", ".", "#", "."},
			[]string{".", ".", "#", ".", "."},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBootProcessBlocksCount3D(t *testing.T) {
	genesis := initializeGrid(testData)
	got := bootProcessBlocksCount(genesis, false)
	want := 112

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestInitializeGrid4D(t *testing.T) {
	got := initializeGrid(testData)
	want := [][][][]string{
		[][][]string{[][]string{[]string{".", "#", "."}, []string{".", ".", "#"}, []string{"#", "#", "#"}}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCountNeighbours4D01(t *testing.T) {
	got := countNeighbours([][][][]string{
		[][][]string{
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", "#", "#"}},
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", "#", "#"}},
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", "#", "#"}},
		},
		[][][]string{
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", "#", "#"}},
			[][]string{[]string{"#", "#", "#"}, []string{"#", "X", "#"}, []string{"#", "#", "#"}},
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", "#", "#"}},
		},
		[][][]string{
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", "#", "#"}},
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", "#", "#"}},
			[][]string{[]string{"#", "#", "#"}, []string{"#", "#", "#"}, []string{"#", "#", "#"}},
		},
	}, 1, 1, 1, 1)
	want := 80

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestNormalizeGrid4D(t *testing.T) {
	got := normalizeGrid([][][][]string{
		[][][]string{
			[][]string{[]string{"#", "#", "#", "#"}, []string{"#", "#", "#", "#"}},
			[][]string{[]string{"#", "#", "#", "#"}, []string{"#", "#", "#", "#"}},
			[][]string{[]string{"#", "#", "#", "#"}, []string{"#", "#", "#", "#"}},
		},
	})

	want := [][][][]string{
		[][][]string{
			[][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			},
		},
		[][][]string{
			[][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", "#", "#", "#", "#", "."},
				[]string{".", "#", "#", "#", "#", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", "#", "#", "#", "#", "."},
				[]string{".", "#", "#", "#", "#", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", "#", "#", "#", "#", "."},
				[]string{".", "#", "#", "#", "#", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			},
		},
		[][][]string{
			[][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			},
			[][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", ".", "."},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNextGeneration4D(t *testing.T) {
	got := nextGeneration([][][][]string{
		[][][]string{
			[][]string{
				[]string{".", "#", "."},
				[]string{".", ".", "#"},
				[]string{"#", "#", "#"},
			},
		},
	}, true)

	want := [][][][]string{
		[][][]string{
			[][]string{
				[]string{".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", "."},
				[]string{".", "#", ".", ".", "."},
				[]string{".", ".", ".", "#", "."},
				[]string{".", ".", "#", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", "."},
				[]string{".", "#", ".", ".", "."},
				[]string{".", ".", ".", "#", "."},
				[]string{".", ".", "#", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", "."},
				[]string{".", "#", ".", ".", "."},
				[]string{".", ".", ".", "#", "."},
				[]string{".", ".", "#", ".", "."},
			},
		},
		[][][]string{
			[][]string{
				[]string{".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", "."},
				[]string{".", "#", ".", ".", "."},
				[]string{".", ".", ".", "#", "."},
				[]string{".", ".", "#", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", "."},
				[]string{".", "#", ".", "#", "."},
				[]string{".", ".", "#", "#", "."},
				[]string{".", ".", "#", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", "."},
				[]string{".", "#", ".", ".", "."},
				[]string{".", ".", ".", "#", "."},
				[]string{".", ".", "#", ".", "."},
			},
		},
		[][][]string{
			[][]string{
				[]string{".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", "."},
				[]string{".", "#", ".", ".", "."},
				[]string{".", ".", ".", "#", "."},
				[]string{".", ".", "#", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", "."},
				[]string{".", "#", ".", ".", "."},
				[]string{".", ".", ".", "#", "."},
				[]string{".", ".", "#", ".", "."},
			}, [][]string{
				[]string{".", ".", ".", ".", "."},
				[]string{".", ".", ".", ".", "."},
				[]string{".", "#", ".", ".", "."},
				[]string{".", ".", ".", "#", "."},
				[]string{".", ".", "#", ".", "."},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBootProcessBlocksCount4D(t *testing.T) {
	genesis := initializeGrid(testData)
	got := bootProcessBlocksCount(genesis, true)
	want := 848

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day17.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day17.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
