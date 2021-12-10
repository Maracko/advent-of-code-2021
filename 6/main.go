package main

import (
	"fmt"

	"github.com/maracko/advent-of-code-2021/helpers"
)

type lanternFish struct {
	timer int64
}

var newFishTimer = 8
var resetFishTimer = 6

func main() {

	fmt.Println("Part 1 solution:", simulate(80))
	fmt.Println("Part 2 solution:", simulate(256))
}

func simulate(nDays int) int {
	data, _ := helpers.ReadFileToSliceOfStrings("data.txt")
	vals := helpers.ReadStringToSliceOfInts(data[0], ",")
	fishes := [9]int{}
	for _, val := range vals {
		fishes[val]++
	}

	for day := 1; day <= nDays; day++ {
		nFishAtZero := fishes[0]
		for i := 1; i < len(fishes); i++ {
			fishes[i-1] = fishes[i]
		}
		fishes[resetFishTimer] += nFishAtZero
		fishes[newFishTimer] = nFishAtZero

	}
	return helpers.SumSliceOfInts(fishes[:])
}
