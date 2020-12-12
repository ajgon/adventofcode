package day12

import (
	"reflect"
	"testing"
	"year2020/helpers"
)

func TestProcess(t *testing.T) {
	ship := Ship{
		x:         0,
		y:         0,
		direction: 'E',
	}
	ship.Process("F10")
	ship.Process("N3")
	ship.Process("F7")
	ship.Process("R90")
	ship.Process("F11")

	got := ship
	want := Ship{x: 17, y: 8, direction: 'S'}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestManhattanDistance(t *testing.T) {
	ship := Ship{
		x:         0,
		y:         0,
		direction: 'E',
	}
	ship.Process("F10")
	ship.Process("N3")
	ship.Process("F7")
	ship.Process("R90")
	ship.Process("F11")

	got := ship.ManhattanDistance()
	want := 25

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestProcessWaypointedShip(t *testing.T) {
	ship := WaypointedShip{
		x:        0,
		y:        0,
		waypoint: Waypoint{x: 10, y: -1},
	}
	ship.Process("F10")

	got := ship
	want := WaypointedShip{x: 100, y: -10, waypoint: Waypoint{x: 10, y: -1}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	ship.Process("N3")

	got = ship
	want = WaypointedShip{x: 100, y: -10, waypoint: Waypoint{x: 10, y: -4}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	ship.Process("F7")

	got = ship
	want = WaypointedShip{x: 170, y: -38, waypoint: Waypoint{x: 10, y: -4}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	ship.Process("R90")

	got = ship
	want = WaypointedShip{x: 170, y: -38, waypoint: Waypoint{x: 4, y: 10}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	ship.Process("F11")

	got = ship
	want = WaypointedShip{x: 214, y: 72, waypoint: Waypoint{x: 4, y: 10}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	gotDistance := ship.ManhattanDistance()
	wantDistance := 286

	if got != want {
		t.Errorf("got %d, want %d", gotDistance, wantDistance)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day12.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day12.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
