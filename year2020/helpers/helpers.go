package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var result []byte

func ProcessInput(file_path string) ([]string, error) {

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

	return strings.Split(strings.TrimSpace(string(result)), "\n"), nil
}
