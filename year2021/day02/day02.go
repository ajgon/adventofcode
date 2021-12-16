package day02

import (
	"fmt"
	"strconv"
	"strings"

	"year2021/helpers"
)

func ProcessInstruction(instruction string) (string, uint32, error) {
	tokens := strings.Split(instruction, " ")
	command := tokens[0]

	value, err := strconv.Atoi(tokens[1])
	if err != nil {
		return "", 0, fmt.Errorf("error while processing instruction: %w", err)
	}

	return command, uint32(value), nil
}

func ProcessInstructionsXY(instructions []string) (x, y uint32) {
	for _, instruction := range instructions {
		command, value, _ := ProcessInstruction(instruction)

		switch command {
		case "forward":
			x += value
		case "up":
			y -= value
		case "down":
			y += value
		}
	}

	return
}

func ProcessInstructionsXYAim(instructions []string) (x, y uint32) {
	aim := 0

	for _, instruction := range instructions {
		command, value, _ := ProcessInstruction(instruction)

		switch command {
		case "forward":
			x += value
			y += uint32(aim * int(value))
		case "up":
			aim -= int(value)
		case "down":
			aim += int(value)
		}
	}

	return
}

func SimpleSolution() string {
	input, _ := helpers.ProcessInput("day02.input")
	x, y := ProcessInstructionsXY(input)
	solution := x * y

	return strconv.Itoa(int(solution))
}

func AdvancedSolution() string {
	input, _ := helpers.ProcessInput("day02.input")
	x, y := ProcessInstructionsXYAim(input)
	solution := x * y

	return strconv.Itoa(int(solution))
}
