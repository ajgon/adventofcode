package day02

import (
	"regexp"
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day02.input")

type Data struct {
	low      uint8
	high     uint8
	letter   byte
	password string
}

func extractData(row string) Data {
	re := regexp.MustCompile("^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$")
	match := re.FindStringSubmatch(row)

	low, _ := strconv.ParseInt(match[1], 10, 8)
	high, _ := strconv.ParseInt(match[2], 10, 8)

	return Data{
		low:      uint8(low),
		high:     uint8(high),
		letter:   byte(match[3][0]),
		password: match[4],
	}
}

func simpleValidateData(data Data) bool {
	var letterCount uint8 = 0

	for i := 0; i < len(data.password); i++ {
		if data.password[i] == data.letter {
			letterCount += 1
		}
	}

	return data.low <= letterCount && letterCount <= data.high
}

func advancedValidateData(data Data) bool {
	low := data.low - 1
	high := data.high - 1
	return (data.password[low] == data.letter && data.password[high] != data.letter) || (data.password[low] != data.letter && data.password[high] == data.letter)
}

func SimpleSolution() string {
	var count uint16 = 0

	for _, row := range input {
		if simpleValidateData(extractData(row)) {
			count += 1
		}
	}

	return strconv.Itoa(int(count))
}

func AdvancedSolution() string {
	var count uint16 = 0

	for _, row := range input {
		if advancedValidateData(extractData(row)) {
			count += 1
		}
	}

	return strconv.Itoa(int(count))
}
