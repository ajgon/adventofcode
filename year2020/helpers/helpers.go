package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var result []byte

func ProcessInput(file_path string, separators ...string) ([]string, error) {
	var separator string

	if len(separators) == 0 {
		separator = "\n"
	} else {
		separator = separators[0]
	}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, file_path) {
			result, err = ioutil.ReadFile(path)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return []string{}, err
	}

	return strings.Split(strings.TrimSpace(string(result)), separator), nil
}

func ArrayContainsStr(stack []string, needle string) bool {
	for _, item := range stack {
		if item == needle {
			return true
		}
	}
	return false
}

func ArrayContainsInt(stack []int, needle int) bool {
	for _, item := range stack {
		if item == needle {
			return true
		}
	}
	return false
}

func ArrayIntersectStr(left []string, right []string) (result []string) {
	result = make([]string, 0)

	for l := 0; l < len(left); l++ {
		if ArrayContainsStr(right, left[l]) {
			result = append(result, left[l])
		}
	}

	return
}

func ArrayExceptStr(left []string, right []string) (result []string) {
	result = make([]string, 0)

	for l := 0; l < len(left); l++ {
		if !ArrayContainsStr(right, left[l]) {
			result = append(result, left[l])
		}
	}

	return
}

func ArrayMaxInt(stack []int) (max int) {
	max = stack[0]

	for i := 0; i < len(stack); i++ {
		if stack[i] > max {
			max = stack[i]
		}
	}
	return
}

func ArrayShiftInt(stack *[]int) int {
	top := (*stack)[0]
	*stack = (*stack)[1:len(*stack)]

	return top
}

func ArrayIndexOfInt(stack []int, needle int) int {
	for i := 0; i < len(stack); i++ {
		if stack[i] == needle {
			return i
		}
	}

	return -1
}

func ConvertStringsToNumbers(data []string) []int {
	var numbers []int
	for _, item := range data {
		number, _ := strconv.Atoi(item)
		numbers = append(numbers, number)
	}

	return numbers
}

func StringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
