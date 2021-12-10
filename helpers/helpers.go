package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func ConvertSliceOfStringsToSliceOfInts(slice []string) []int {
	var res []int
	for _, s := range slice {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return res
}

func ReadStringToSliceOfInts(s, delimiter string) []int {
	var res []int
	for _, s := range strings.Split(s, delimiter) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}
	return res
}

func SumSliceOfInts(slice []int) int {
	var res int
	for _, i := range slice {
		res += i
	}
	return res
}
