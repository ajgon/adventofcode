package day25

import (
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day25.input")

func extractKeys(data []string) (cardKey, doorKey int) {
	cardKey, _ = strconv.Atoi(data[0])
	doorKey, _ = strconv.Atoi(data[1])

	return
}

func rotate(number, subject int) int {
	return (number * subject) % 20201227
}

func findLoopSize(target, subject int) (loop int) {
	rotation := 1

	for {
		if rotation == target {
			return loop
		}

		loop += 1
		rotation = rotate(rotation, subject)
	}
}

func findEncrypionKey(publicKey, loop int) (encryptionKey int) {
	encryptionKey = 1

	for i := 0; i < loop; i++ {
		encryptionKey = rotate(encryptionKey, publicKey)
	}

	return
}

func SimpleSolution() string {
	cardKey, doorKey := extractKeys(input)
	cardLoop, doorLoop := findLoopSize(cardKey, 7), findLoopSize(doorKey, 7)

	encryptionKeyCard := findEncrypionKey(cardKey, doorLoop)
	encryptionKeyDoor := findEncrypionKey(doorKey, cardLoop)

	if encryptionKeyCard == encryptionKeyDoor {
		return strconv.Itoa(encryptionKeyCard)
	}

	return ""
}
