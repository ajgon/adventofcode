package day23

import (
	"strconv"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day23.input", "")

// naive part 1
func pickCups(board *[]int, position int) (pickedCups []int) {
	pickedCups = make([]int, 0)

	if position > len(*board)-1 {
		position = 0
	}

	for {
		newBoard := make([]int, 0)
		pickedCups = append(pickedCups, (*board)[position])
		for i := 0; i < len(*board); i++ {
			if i != position {
				newBoard = append(newBoard, (*board)[i])
			}
		}
		(*board) = newBoard

		if position == len(*board) {
			position = 0
		}

		if len(pickedCups) == 3 {
			break
		}
	}

	return
}

func placeCups(board *[]int, cups []int, position int) {
	newBoard := make([]int, 0)
	for i := 0; i < position; i++ {
		newBoard = append(newBoard, (*board)[i])
	}
	for i := 0; i < len(cups); i++ {
		newBoard = append(newBoard, cups[i])
	}
	for i := position; i < len(*board); i++ {
		newBoard = append(newBoard, (*board)[i])
	}
	(*board) = newBoard
}

func playMove(board []int, currentCup int) ([]int, int) {
	currentPosition := helpers.ArrayIndexOfInt(board, currentCup)
	pickedCups := pickCups(&board, currentPosition+1)
	destinationCup := currentCup

	for {
		if destinationCup == 1 {
			destinationCup = 9
		} else {
			destinationCup -= 1
		}

		if !helpers.ArrayContainsInt(pickedCups, destinationCup) {
			break
		}
	}

	placeCups(&board, pickedCups, helpers.ArrayIndexOfInt(board, destinationCup)+1)

	currentCupIndex := helpers.ArrayIndexOfInt(board, currentCup)
	newCup := 0

	if currentCupIndex == len(board)-1 {
		newCup = board[0]
	} else {
		newCup = board[currentCupIndex+1]
	}

	return board, newCup
}

func playGames(board []int, number int) []int {
	currentCup := board[0]
	for i := 0; i < number; i++ {
		board, currentCup = playMove(board, currentCup)
	}

	return board
}

// part 2 with linked list
type Cup struct {
	Value int
	Prev  *Cup
	Next  *Cup
}

func buildCupList(initialBoard []int, size int) (*Cup, map[int]*Cup) {
	cupMap := make(map[int]*Cup)
	firstCup := &Cup{Value: initialBoard[0]}
	cupMap[firstCup.Value] = firstCup

	previousCup := firstCup
	newCup := &Cup{}

	for i := 1; i < len(initialBoard); i++ {
		newCup = &Cup{Value: initialBoard[i], Prev: previousCup}
		previousCup.Next = newCup
		previousCup = newCup
		cupMap[newCup.Value] = newCup
	}

	for i := 10; i < size+1; i++ {
		newCup = &Cup{Value: i, Prev: previousCup}
		previousCup.Next = newCup
		previousCup = newCup
		cupMap[newCup.Value] = newCup
	}
	newCup.Next = firstCup
	firstCup.Prev = newCup

	return firstCup, cupMap
}

func playLinkedGames(firstCup *Cup, cupMap map[int]*Cup, number int) *Cup {
	var destinationCup int
	maxCupNumber := len(cupMap)
	currentCup := firstCup

	for i := 0; i < number; i++ {
		cup1 := currentCup.Next
		cup2 := currentCup.Next.Next
		cup3 := currentCup.Next.Next.Next
		pickedCups := []int{cup1.Value, cup2.Value, cup3.Value}

		currentCup.Next = cup3.Next
		currentCup.Next.Prev = currentCup

		if currentCup.Value == 1 {
			destinationCup = maxCupNumber
		} else {
			destinationCup = currentCup.Value - 1
		}

		for {
			if !helpers.ArrayContainsInt(pickedCups, destinationCup) {
				break
			}

			if destinationCup == 1 {
				destinationCup = maxCupNumber
			} else {
				destinationCup -= 1
			}
		}

		newCup := cupMap[destinationCup]
		cup1.Prev = newCup
		cup3.Next = newCup.Next
		newCup.Next.Prev = cup3
		newCup.Next = cup1

		currentCup = currentCup.Next
	}

	return firstCup
}

func SimpleSolution() string {
	result := ""
	board := helpers.ConvertStringsToNumbers(input)

	board = playGames(board, 100)
	start := helpers.ArrayIndexOfInt(board, 1)

	for i := start + 1; i < len(board); i++ {
		result += strconv.Itoa(board[i])
	}

	for i := 0; i < start; i++ {
		result += strconv.Itoa(board[i])
	}

	return result
}

func AdvancedSolution() string {
	board := helpers.ConvertStringsToNumbers(input)
	firstCup, cupMap := buildCupList(board, 1000000)

	cup := playLinkedGames(firstCup, cupMap, 10000000)

	for {
		if cup.Value == 1 {
			break
		}
		cup = cup.Next
	}

	return strconv.Itoa(cup.Next.Value * cup.Next.Next.Value)
}
