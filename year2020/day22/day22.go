package day22

import (
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day22.input")

func extractDecks(data []string) (deckA, deckB []int) {
	currentDeck := &deckA

	for i := 0; i < len(data); i++ {
		if len(data[i]) > 6 && data[i][0:6] == "Player" {
			continue
		}

		if data[i] == "" {
			currentDeck = &deckB
			continue
		}

		number, _ := strconv.Atoi(data[i])
		*currentDeck = append(*currentDeck, number)
	}

	return
}

func handleWinner(cardA, cardB int, deckA, deckB []int) ([]int, []int) {
	if cardA > cardB {
		deckA = append(deckA, cardA)
		deckA = append(deckA, cardB)
	} else {
		deckB = append(deckB, cardB)
		deckB = append(deckB, cardA)
	}

	return deckA, deckB
}

func playRound(deckA, deckB []int) ([]int, []int) {
	cardA := helpers.ArrayShiftInt(&deckA)
	cardB := helpers.ArrayShiftInt(&deckB)

	deckA, deckB = handleWinner(cardA, cardB, deckA, deckB)

	return deckA, deckB
}

func playRecursiveRound(deckA, deckB []int) ([]int, []int) {
	cardA := deckA[0]
	cardB := deckB[0]

	smallDeckA := deckA[1:]
	smallDeckB := deckB[1:]

	if cardA <= len(smallDeckA) && cardB <= len(smallDeckB) {
		smallDeckAClone := make([]int, len(smallDeckA))
		smallDeckBClone := make([]int, len(smallDeckB))
		copy(smallDeckAClone, smallDeckA)
		copy(smallDeckBClone, smallDeckB)
		_, winner := playGame(smallDeckA[0:cardA], smallDeckB[0:cardB], playRecursiveRound)
		copy(smallDeckA, smallDeckAClone)
		copy(smallDeckB, smallDeckBClone)
		if winner == 'A' {
			smallDeckA = append(smallDeckA, cardA)
			smallDeckA = append(smallDeckA, cardB)
		} else {
			smallDeckB = append(smallDeckB, cardB)
			smallDeckB = append(smallDeckB, cardA)
		}
	} else {
		smallDeckA, smallDeckB = handleWinner(cardA, cardB, smallDeckA, smallDeckB)
	}

	return smallDeckA, smallDeckB
}

func includesDeckSet(stack [][][]int, deckA []int, deckB []int) bool {
	for i := 0; i < len(stack); i++ {
		aMatches := false
		bMatches := false

		if len(stack[i][0]) == len(deckA) {
			aMatches = true
			for a := 0; a < len(deckA); a++ {
				if stack[i][0][a] != deckA[a] {
					aMatches = false
				}
			}

			if !aMatches {
				continue
			}
		}

		if len(stack[i][1]) == len(deckB) {
			bMatches = true
			for b := 0; b < len(deckB); b++ {
				if stack[i][1][b] != deckB[b] {
					bMatches = false
				}
			}

			if bMatches {
				return true
			}
		}
	}

	return false
}

func playGame(deckA, deckB []int, playFunc func([]int, []int) ([]int, []int)) ([]int, rune) {
	var winningDeck []int
	winner := ' '
	rounds := make([][][]int, 0)

	for {
		if includesDeckSet(rounds, deckA, deckB) || len(deckB) == 0 {
			winningDeck = deckA
			winner = 'A'
			break
		}
		if len(deckA) == 0 {
			winningDeck = deckB
			winner = 'B'
			break
		}

		rounds = append(rounds, [][]int{deckA, deckB})
		deckA, deckB = playFunc(deckA, deckB)
	}

	return winningDeck, winner
}

func SimpleSolution() string {
	deckA, deckB := extractDecks(input)
	result := 0

	winningDeck, _ := playGame(deckA, deckB, playRound)

	for i := 0; i < len(winningDeck); i++ {
		result += winningDeck[i] * (len(winningDeck) - i)
	}

	return strconv.Itoa(result)
}

func AdvancedSolution() string {
	deckA, deckB := extractDecks(input)
	result := 0

	winningDeck, _ := playGame(deckA, deckB, playRecursiveRound)

	for i := 0; i < len(winningDeck); i++ {
		result += winningDeck[i] * (len(winningDeck) - i)
	}

	return strconv.Itoa(result)
}
