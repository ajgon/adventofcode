package day08

import (
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day08.input")

type CPU struct {
	pc      int16
	acc     int16
	history []int
}

func (c *CPU) ProcessMnemonic(mnemonic string) {
	mnemonicParts := strings.Split(mnemonic, " ")
	mnemonicValue, _ := strconv.Atoi(mnemonicParts[1])
	mnemonic, value := mnemonicParts[0], int16(mnemonicValue)

	switch mnemonic {
	case "acc":
		c.acc += value
		c.pc += 1
	case "jmp":
		c.pc += value
	case "nop":
		c.pc += 1
	default:
		panic("unknown mnemonic")
	}
}

func (c *CPU) Reset() {
	c.pc = 0
	c.acc = 0
	c.history = make([]int, 0)
}

func (c *CPU) RunUntilLoop(program []string) {
	program = append(program, "END")

	for {
		if helpers.ArrayContainsInt(c.history, int(c.pc)) {
			return
		}
		c.history = append(c.history, int(c.pc))

		c.ProcessMnemonic(program[c.pc])

		if program[c.pc] == "END" {
			c.pc = -1
			return
		}
	}
}

func SimpleSolution() string {
	cpu := CPU{pc: 0, acc: 0}
	cpu.Reset()
	cpu.RunUntilLoop(input)

	return strconv.Itoa(int(cpu.acc))
}

func AdvancedSolution() string {
	cpu := CPU{pc: 0, acc: 0}

	for index, _ := range input {
		cpu.Reset()
		alteredInput := make([]string, len(input))
		copy(alteredInput, input)

		if alteredInput[index][0:3] == "acc" {
			continue
		}

		if alteredInput[index][0:3] == "jmp" {
			alteredInput[index] = "nop" + alteredInput[index][3:]
		} else {
			alteredInput[index] = "jmp" + alteredInput[index][3:]
		}

		cpu.RunUntilLoop(alteredInput)

		if cpu.pc == -1 {
			return strconv.Itoa(int(cpu.acc))
		}
	}

	return ""
}
