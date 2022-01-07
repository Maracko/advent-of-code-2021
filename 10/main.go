package main

import (
	"fmt"

	"github.com/maracko/advent-of-code-2021/helpers/file"
	"github.com/maracko/advent-of-code-2021/helpers/slice"
	"github.com/maracko/advent-of-code-2021/helpers/stack"
)

var openers = []rune{'(', '[', '{', '<'}
var closers = []rune{')', ']', '}', '>'}

var pointsMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	d, _ := file.ReadFileToSliceOfStrings("data.txt")

	fmt.Println("Part 1 solution:", solvePart1(d))
}

func solvePart1(data []string) int {
	corruptedLines := make(map[string]rune, len(data))

lineLoop:
	for _, line := range data {
		stack := stack.Stack{}
		for _, char := range line {
			isOpener, _ := slice.RuneInSlice(openers, char)
			// Saw in test data that this case is never hit
			// if i == 0 && !isOpener {
			// 	corruptedLines = append(corruptedLines, line)
			// 	continue lineLoop
			// }

			if isOpener {
				stack.Push(char)
			}

			if isCloser, closerIdx := slice.RuneInSlice(closers, char); isCloser {
				elem := stack.Pop().(rune)
				_, openerIdx := slice.RuneInSlice(openers, elem)

				if openerIdx != closerIdx {
					corruptedLines[line] = closers[closerIdx]
					continue lineLoop
				}
			}
		}
	}
	res := 0
	counter := map[rune]int{}
	for _, closer := range corruptedLines {
		counter[closer]++
	}
	for closer, count := range counter {
		res += pointsMap[closer] * count
	}
	return res
}

// func solvePart2(){}
