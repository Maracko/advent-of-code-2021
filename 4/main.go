package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maracko/advent-of-code-2021/helpers/file"
)

type line [5]int
type board [5]line

var numbers []int

func main() {
	data, _ := file.ReadFileToSliceOfStrings("data.txt")
	numbers = getAllDrawnNumbers(data[0])

	fmt.Println("Part 1 solution:", solvePart1(data))
	fmt.Println("Part 2 solution:", solvePart2(data))

}

func getAllDrawnNumbers(nums string) []int {
	n := strings.Split(nums, ",")
	numbers := make([]int, len(n))
	for i, v := range n {
		numbers[i], _ = strconv.Atoi(v)
	}
	return numbers
}

func sumAllRemainingNumbersOnBoard(nDrawn int, b board) int {
	var sum int
	for _, row := range b {
		for _, num := range row {
			if !numInArray(num, numbers[:nDrawn]) {
				sum += num
			}
		}
	}
	return sum
}

func lineWon(nDrawn int, lns [5]line) bool {
	for _, ln := range lns {
		if isWinningLine(nDrawn, ln) {
			return true
		}
	}
	return false
}

func getAllRowsAndColumns(b board) ([5]line, [5]line) {
	var rows [5]line
	var columns [5]line

	for rowIdx, row := range b {
		rows[rowIdx] = row
		for colIdx, num := range row {
			columns[colIdx][rowIdx] = num
		}

	}

	return rows, columns
}

func isWinningLine(nDrawn int, ln line) bool {
	if nDrawn < 5 {
		return false
	}

	nMatched := 0
	for _, num := range ln {
		if numInArray(num, numbers[:nDrawn]) {
			nMatched++
		}
	}

	return nMatched == 5
}

func numInArray(num int, arr []int) bool {
	for _, val := range arr {
		if val == num {
			return true
		}
	}
	return false
}

func solvePart1(data []string) int {
	n := strings.Split(data[0], ",")
	numbers = make([]int, len(n))
	for i, v := range n {
		numbers[i], _ = strconv.Atoi(v)
	}

	var boards []board
	var b board
	skipped := 0
	for i, line := range data[2:] {
		if line == "" {
			skipped++
			continue
		}
		var rowIdx int
		if i < 5 {
			rowIdx = i
		} else {
			rowIdx = (i - skipped) % 5
		}

		if rowIdx == 0 {
			b = board{}
		}

		col := 0
		for _, num := range strings.Split(line, " ") {
			if num == "" {
				continue
			}
			b[rowIdx][col], _ = strconv.Atoi(num)
			col += 1
		}

		if rowIdx == 4 {
			boards = append(boards, b)
		}

	}

	numbersDrawn := 0
	number := 0
	winningBoard := board{}
numLoop:
	for i, num := range numbers {
		nDrawn := i + 1
		for _, board := range boards {
			rows, columns := getAllRowsAndColumns(board)
			boardWon := false

			if won := lineWon(nDrawn, rows); won {
				boardWon = true

			}
			if won := lineWon(nDrawn, columns); won {
				boardWon = true
			}

			if boardWon {
				// fmt.Printf("Board %d is first winner after %d numbers drawn. Last drawn number: %d\n", boardIdx+1, nDrawn, num)
				numbersDrawn = nDrawn
				number = num
				winningBoard = board
				break numLoop
			}
		}
	}

	return sumAllRemainingNumbersOnBoard(numbersDrawn, winningBoard) * number
}

func solvePart2(data []string) int {
	var boards []board
	var b board
	skipped := 0
	for i, line := range data[2:] {
		if line == "" {
			skipped++
			continue
		}
		var rowIdx int
		if i < 5 {
			rowIdx = i
		} else {
			rowIdx = (i - skipped) % 5
		}

		if rowIdx == 0 {
			b = board{}
		}

		col := 0
		for _, num := range strings.Split(line, " ") {
			if num == "" {
				continue
			}
			b[rowIdx][col], _ = strconv.Atoi(num)
			col += 1
		}

		if rowIdx == 4 {
			boards = append(boards, b)
		}

	}

	lastWinningBoard := board{}
	lastDrawnNumIdx := 0
	for i := range numbers {
		nDrawn := i + 1
		for j := 0; j < len(boards); {
			for boardIdx, board := range boards {
				j++
				rows, columns := getAllRowsAndColumns(board)
				winnerBoard := false

				if won := lineWon(nDrawn, rows); won {
					winnerBoard = true
				}
				if won := lineWon(nDrawn, columns); won {
					winnerBoard = true
				}

				if winnerBoard {
					lastWinningBoard = board
					lastDrawnNumIdx = i
					boards = append(boards[:boardIdx], boards[boardIdx+1:]...)
					j = 0
					break
				}
			}

		}
	}
	return sumAllRemainingNumbersOnBoard(lastDrawnNumIdx+1, lastWinningBoard) * numbers[lastDrawnNumIdx]
}
