package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ProcessInput(filePath string, separators ...string) ([]string, error) {
	var (
		separator string
		result    []byte
	)

	if len(separators) == 0 {
		separator = "\n"
	} else {
		separator = separators[0]
	}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, filePath) {
			result, err = ioutil.ReadFile(path)

			if err != nil {
				return fmt.Errorf("error processing input: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return []string{}, fmt.Errorf("error processing input %w", err)
	}

	return strings.Split(strings.TrimSpace(string(result)), separator), nil
}

func ConvertStringsToNumbers(data []string) (numbers []int) {
	for _, item := range data {
		number, _ := strconv.Atoi(item)
		numbers = append(numbers, number)
	}

	return
}
