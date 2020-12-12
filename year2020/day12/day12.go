package day12

import (
	"math"
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day12.input")

type Ship struct {
	x         int
	y         int
	direction byte
}

type Waypoint struct {
	x int
	y int
}

type WaypointedShip struct {
	x        int
	y        int
	waypoint Waypoint
}

var rotationMap = map[byte]map[byte]byte{
	'R': map[byte]byte{
		'N': 'E',
		'E': 'S',
		'S': 'W',
		'W': 'N',
	},
	'L': map[byte]byte{
		'N': 'W',
		'W': 'S',
		'S': 'E',
		'E': 'N',
	},
}

func parseInstruction(instruction string) (action byte, value int) {
	action = instruction[0]
	value, _ = strconv.Atoi(instruction[1:])

	return
}

func manhattanDistance(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func (s *Ship) Process(instruction string) {
	action, value := parseInstruction(instruction)

	if action == 'F' {
		action = s.direction
	}

	switch action {
	case 'N':
		s.y -= value
	case 'E':
		s.x += value
	case 'S':
		s.y += value
	case 'W':
		s.x -= value
	case 'R', 'L':
		switch value {
		case 90:
			s.direction = rotationMap[action][s.direction]
		case 180:
			s.direction = rotationMap[action][rotationMap[action][s.direction]]
		case 270:
			s.direction = rotationMap[action][rotationMap[action][rotationMap[action][s.direction]]]
		}
	}
}

func (s *Ship) ManhattanDistance() int {
	return manhattanDistance(s.x, s.y)
}

func (ws *WaypointedShip) Process(instruction string) {
	action, value := parseInstruction(instruction)

	if action == 'F' {
		ws.moveRelativeToWaypoint(value)
		return
	}

	ws.waypoint.Process(instruction)
}

func (ws *WaypointedShip) moveRelativeToWaypoint(value int) {
	if ws.waypoint.x > 0 {
		ws.x += value * int(math.Abs(float64(ws.waypoint.x)))
	} else {
		ws.x -= value * int(math.Abs(float64(ws.waypoint.x)))
	}

	if ws.waypoint.y > 0 {
		ws.y += value * int(math.Abs(float64(ws.waypoint.y)))
	} else {
		ws.y -= value * int(math.Abs(float64(ws.waypoint.y)))
	}
}

func (ws *WaypointedShip) ManhattanDistance() int {
	return manhattanDistance(ws.x, ws.y)
}

func (w *Waypoint) Process(instruction string) {
	action, value := parseInstruction(instruction)

	switch action {
	case 'N':
		w.y -= value
	case 'E':
		w.x += value
	case 'S':
		w.y += value
	case 'W':
		w.x -= value
	case 'R':
		for i := 0; i < (value / 90); i++ {
			tmp := w.x
			w.x = -w.y
			w.y = tmp
		}
	case 'L':
		for i := 0; i < (value / 90); i++ {
			tmp := w.x
			w.x = w.y
			w.y = -tmp
		}
	}
}

func SimpleSolution() string {
	ship := Ship{x: 0, y: 0, direction: 'E'}

	for _, instruction := range input {
		ship.Process(instruction)
	}

	return strconv.Itoa(ship.ManhattanDistance())
}

func AdvancedSolution() string {
	ship := WaypointedShip{x: 0, y: 0, waypoint: Waypoint{x: 10, y: -1}}

	for _, instruction := range input {
		ship.Process(instruction)
	}

	return strconv.Itoa(ship.ManhattanDistance())
}
