package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Cannot open file", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lastVal := 0
	i := 0
	timesHigher := 0
	for scanner.Scan() {
		i++
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Cannot convert to int", err)
			return
		}

		if i == 1 {
			lastVal = val
			continue
		}
		if val > lastVal {
			timesHigher++
		}
		lastVal = val
	}
	fmt.Println("Times value was higher than last one", timesHigher)
}
