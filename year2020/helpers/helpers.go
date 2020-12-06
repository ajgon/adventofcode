package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"
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

func ArrayContains(stack []string, needle string) bool {
	for _, item := range stack {
		if item == needle {
			return true
		}
	}
	return false
}
