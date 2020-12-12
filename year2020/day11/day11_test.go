package day11

import (
	"reflect"
	"testing"
	"year2020/helpers"
)

var testData string = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

var testRayData01 string = `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`

var testRayData02 string = `.............
.L.L.#.#.#.#.
.............`

var testRayData03 string = `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`

func TestGetMap(t *testing.T) {
	got := getMap(testData)
	want := [][]byte{
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
		{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestGetNearby(t *testing.T) {
	testDataMap := getMap(testData)

	got := []byte{
		getNearby(testDataMap, 0, 0, -1, -1),
		getNearby(testDataMap, 0, 0, 0, -1),
		getNearby(testDataMap, 0, 0, 1, -1),
		getNearby(testDataMap, 0, 0, -1, 0),
		getNearby(testDataMap, 0, 0, 1, 0),
		getNearby(testDataMap, 0, 0, -1, 1),
		getNearby(testDataMap, 0, 0, 0, 1),
		getNearby(testDataMap, 0, 0, 1, 1),
		getNearby(testDataMap, 3, 3, -1, -1),
		getNearby(testDataMap, 3, 3, 0, -1),
		getNearby(testDataMap, 3, 3, 1, -1),
		getNearby(testDataMap, 3, 3, -1, 0),
		getNearby(testDataMap, 3, 3, 1, 0),
		getNearby(testDataMap, 3, 3, -1, 1),
		getNearby(testDataMap, 3, 3, 0, 1),
		getNearby(testDataMap, 3, 3, 1, 1),
	}

	want := []byte{
		'.', '.', '.', '.', '.', '.', 'L', 'L',
		'L', '.', 'L', 'L', '.', 'L', 'L', '.',
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestGetRay(t *testing.T) {
	testRayData01Map := getMap(testRayData01)
	testRayData02Map := getMap(testRayData02)
	testRayData03Map := getMap(testRayData03)

	got := []byte{
		getRay(testRayData01Map, 3, 4, -1, -1),
		getRay(testRayData01Map, 3, 4, 0, -1),
		getRay(testRayData01Map, 3, 4, 1, -1),
		getRay(testRayData01Map, 3, 4, -1, 0),
		getRay(testRayData01Map, 3, 4, 1, 0),
		getRay(testRayData01Map, 3, 4, -1, 1),
		getRay(testRayData01Map, 3, 4, 0, 1),
		getRay(testRayData01Map, 3, 4, 1, 1),
		getRay(testRayData02Map, 1, 1, 1, 0),
		getRay(testRayData03Map, 3, 3, -1, -1),
		getRay(testRayData03Map, 3, 3, 0, -1),
		getRay(testRayData03Map, 3, 3, 1, -1),
		getRay(testRayData03Map, 3, 3, -1, 0),
		getRay(testRayData03Map, 3, 3, 1, 0),
		getRay(testRayData03Map, 3, 3, -1, 1),
		getRay(testRayData03Map, 3, 3, 0, 1),
		getRay(testRayData03Map, 3, 3, 1, 1),
	}
	want := []byte{
		'#', '#', '#', '#', '#', '#', '#', '#',
		'L',
		'.', '.', '.', '.', '.', '.', '.', '.',
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestGetOccupiedSeats(t *testing.T) {
	testDataMap := getMap(testData)
	testRayData01Map := getMap(testRayData01)
	testRayData02Map := getMap(testRayData02)
	testRayData03Map := getMap(testRayData03)

	got := []int{
		getOccupiedSeats(testDataMap, 0, 0, true),
		getOccupiedSeats(testDataMap, 3, 3, true),
		getOccupiedSeats(testRayData01Map, 3, 4, true),
		getOccupiedSeats(testRayData02Map, 1, 1, true),
		getOccupiedSeats(testRayData03Map, 3, 3, true),
		getOccupiedSeats(testDataMap, 0, 0, false),
		getOccupiedSeats(testDataMap, 3, 3, false),
		getOccupiedSeats(testRayData01Map, 3, 4, false),
		getOccupiedSeats(testRayData02Map, 1, 1, false),
		getOccupiedSeats(testRayData03Map, 3, 3, false),
	}
	want := []int{
		0, 0, 2, 0, 0,
		0, 0, 8, 0, 0,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestReorganizePeopleWithNearby(t *testing.T) {
	testDataMap := getMap(testData)

	got1 := regorganizePeople(testDataMap, true, 4)
	got2 := regorganizePeople(got1, true, 4)
	got3 := regorganizePeople(got2, true, 4)
	got4 := regorganizePeople(got3, true, 4)
	got5 := regorganizePeople(got4, true, 4)
	got6 := regorganizePeople(got5, true, 4)
	want1 := [][]byte{
		{'#', '.', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '.', '#', '.', '.', '#', '.', '.'},
		{'#', '#', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '.', '#', '#'},
		{'.', '.', '#', '.', '#', '.', '.', '.', '.', '.'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '#', '.', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '.', '#', '#'},
	}
	want2 := [][]byte{
		{'#', '.', 'L', 'L', '.', 'L', '#', '.', '#', '#'},
		{'#', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', '#'},
		{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
		{'#', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', '#'},
		{'#', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'#', '.', 'L', 'L', 'L', 'L', '#', '.', '#', '#'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'#', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'#', '.', '#', 'L', 'L', 'L', 'L', '.', '#', '#'},
	}
	want3 := [][]byte{
		{'#', '.', '#', '#', '.', 'L', '#', '.', '#', '#'},
		{'#', 'L', '#', '#', '#', 'L', 'L', '.', 'L', '#'},
		{'L', '.', '#', '.', '#', '.', '.', '#', '.', '.'},
		{'#', 'L', '#', '#', '.', '#', '#', '.', 'L', '#'},
		{'#', '.', '#', '#', '.', 'L', 'L', '.', 'L', 'L'},
		{'#', '.', '#', '#', '#', 'L', '#', '.', '#', '#'},
		{'.', '.', '#', '.', '#', '.', '.', '.', '.', '.'},
		{'#', 'L', '#', '#', '#', '#', '#', '#', 'L', '#'},
		{'#', '.', 'L', 'L', '#', '#', '#', 'L', '.', 'L'},
		{'#', '.', '#', 'L', '#', '#', '#', '.', '#', '#'},
	}
	want4 := [][]byte{
		{'#', '.', '#', 'L', '.', 'L', '#', '.', '#', '#'},
		{'#', 'L', 'L', 'L', '#', 'L', 'L', '.', 'L', '#'},
		{'L', '.', 'L', '.', 'L', '.', '.', '#', '.', '.'},
		{'#', 'L', 'L', 'L', '.', '#', '#', '.', 'L', '#'},
		{'#', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'#', '.', 'L', 'L', '#', 'L', '#', '.', '#', '#'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'#', 'L', '#', 'L', 'L', 'L', 'L', '#', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'#', '.', '#', 'L', '#', 'L', '#', '.', '#', '#'},
	}
	want5 := [][]byte{
		{'#', '.', '#', 'L', '.', 'L', '#', '.', '#', '#'},
		{'#', 'L', 'L', 'L', '#', 'L', 'L', '.', 'L', '#'},
		{'L', '.', '#', '.', 'L', '.', '.', '#', '.', '.'},
		{'#', 'L', '#', '#', '.', '#', '#', '.', 'L', '#'},
		{'#', '.', '#', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'#', '.', '#', 'L', '#', 'L', '#', '.', '#', '#'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'#', 'L', '#', 'L', '#', '#', 'L', '#', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'#', '.', '#', 'L', '#', 'L', '#', '.', '#', '#'},
	}

	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("got %+v, want %+v", got1, want1)
	}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %+v, want %+v", got2, want2)
	}

	if !reflect.DeepEqual(got3, want3) {
		t.Errorf("got %+v, want %+v", got3, want3)
	}

	if !reflect.DeepEqual(got4, want4) {
		t.Errorf("got %+v, want %+v", got4, want4)
	}

	if !reflect.DeepEqual(got5, want5) {
		t.Errorf("got %+v, want %+v", got5, want5)
	}

	if !reflect.DeepEqual(got6, want5) {
		t.Errorf("got %+v, want %+v", got6, want5)
	}
}

func TestReorganizePeopleWithRay(t *testing.T) {
	testDataMap := getMap(testData)

	got1 := regorganizePeople(testDataMap, false, 5)
	got2 := regorganizePeople(got1, false, 5)
	got3 := regorganizePeople(got2, false, 5)
	got4 := regorganizePeople(got3, false, 5)
	got5 := regorganizePeople(got4, false, 5)
	got6 := regorganizePeople(got5, false, 5)
	got7 := regorganizePeople(got6, false, 5)
	want1 := [][]byte{
		{'#', '.', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '.', '#', '.', '.', '#', '.', '.'},
		{'#', '#', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '.', '#', '#'},
		{'.', '.', '#', '.', '#', '.', '.', '.', '.', '.'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '#', '.', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '.', '#', '#'},
	}
	want2 := [][]byte{
		{'#', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', '#'},
		{'#', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
		{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', '#'},
	}
	want3 := [][]byte{
		{'#', '.', 'L', '#', '.', '#', '#', '.', 'L', '#'},
		{'#', 'L', '#', '#', '#', '#', '#', '.', 'L', 'L'},
		{'L', '.', '#', '.', '#', '.', '.', '#', '.', '.'},
		{'#', '#', 'L', '#', '.', '#', '#', '.', '#', '#'},
		{'#', '.', '#', '#', '.', '#', 'L', '.', '#', '#'},
		{'#', '.', '#', '#', '#', '#', '#', '.', '#', 'L'},
		{'.', '.', '#', '.', '#', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', '#', '#', '#', '#', 'L', 'L', '#'},
		{'#', '.', 'L', '#', '#', '#', '#', '#', '.', 'L'},
		{'#', '.', 'L', '#', '#', '#', '#', '.', 'L', '#'},
	}
	want4 := [][]byte{
		{'#', '.', 'L', '#', '.', 'L', '#', '.', 'L', '#'},
		{'#', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', '#', '.', '.'},
		{'#', '#', 'L', 'L', '.', 'L', 'L', '.', 'L', '#'},
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', '#', '.', 'L'},
		{'#', '.', 'L', '#', 'L', 'L', '#', '.', 'L', '#'},
	}
	want5 := [][]byte{
		{'#', '.', 'L', '#', '.', 'L', '#', '.', 'L', '#'},
		{'#', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', '#', '.', '.'},
		{'#', '#', 'L', '#', '.', '#', 'L', '.', 'L', '#'},
		{'L', '.', 'L', '#', '.', '#', 'L', '.', 'L', '#'},
		{'#', '.', 'L', '#', '#', '#', '#', '.', 'L', 'L'},
		{'.', '.', '#', '.', '#', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', '#', '#', '#', 'L', 'L', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', '#', '.', 'L'},
		{'#', '.', 'L', '#', 'L', 'L', '#', '.', 'L', '#'},
	}
	want6 := [][]byte{
		{'#', '.', 'L', '#', '.', 'L', '#', '.', 'L', '#'},
		{'#', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', '#', '.', '.'},
		{'#', '#', 'L', '#', '.', '#', 'L', '.', 'L', '#'},
		{'L', '.', 'L', '#', '.', 'L', 'L', '.', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', '#', '.', 'L', 'L'},
		{'.', '.', '#', '.', 'L', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', '#', '#', '#', 'L', 'L', 'L', '#'},
		{'#', '.', 'L', 'L', 'L', 'L', 'L', '#', '.', 'L'},
		{'#', '.', 'L', '#', 'L', 'L', '#', '.', 'L', '#'},
	}

	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("got %+v, want %+v", got1, want1)
	}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %+v, want %+v", got2, want2)
	}

	if !reflect.DeepEqual(got3, want3) {
		t.Errorf("got %+v, want %+v", got3, want3)
	}

	if !reflect.DeepEqual(got4, want4) {
		t.Errorf("got %+v, want %+v", got4, want4)
	}

	if !reflect.DeepEqual(got5, want5) {
		t.Errorf("got %+v, want %+v", got5, want5)
	}

	if !reflect.DeepEqual(got6, want6) {
		t.Errorf("got %+v, want %+v", got6, want5)
	}

	if !reflect.DeepEqual(got7, want6) {
		t.Errorf("got %+v, want %+v", got6, want5)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day11.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day11.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
