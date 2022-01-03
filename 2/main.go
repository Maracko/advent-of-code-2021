package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maracko/advent-of-code-2021/helpers/file"
)

func main() {
	vals, _ := file.ReadFileToSliceOfStrings("data.txt")

	depth, forward := solvePart1(vals)
	fmt.Println("Part 1 solution:", depth*forward)
	depth, forward, _ = solvePart2(vals)
	fmt.Println("Part 2 solution:", depth*forward)
}

func solvePart1(data []string) (forward, depth int) {
	for _, line := range data {
		raw := strings.Split(line, " ")
		direction := raw[0]
		count, _ := strconv.Atoi(raw[1])

		switch direction {
		case "forward":
			forward += count
		case "up":
			depth -= count
		case "down":
			depth += count
		}
	}
	return forward, depth
}

func solvePart2(data []string) (depth, forward, aim int) {
	for _, line := range data {
		raw := strings.Split(line, " ")
		direction := raw[0]
		count, _ := strconv.Atoi(raw[1])

		switch direction {
		case "forward":
			forward += count
			depth += count * aim
		case "up":
			aim -= count
		case "down":
			aim += count
		}

	}
	return
}
