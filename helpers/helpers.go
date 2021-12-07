package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileToSliceOfStrings(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Cannot open file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, nil
}
