package main

import (
	"fmt"
	"strconv"

	"github.com/maracko/advent-of-code-2021/helpers/file"
)

func main() {

	data, _ := file.ReadFileToSliceOfStrings("data.txt")

	fmt.Println("Part 1 solution:", solvePart1(data))
	fmt.Println("Part 2 solution:", solvePart2(data))
}

func createCounter(len int) map[int]map[string]int {
	counter := make(map[int]map[string]int, len)
	for i := 0; i < len; i++ {
		counter[i] = make(map[string]int, 2)
	}
	return counter
}

func solvePart1(data []string) int {
	counter := createCounter(len(data[0]))

	for _, line := range data {
		for i, char := range line {
			counter[i][string(char)] += 1
		}
	}

	gamma := ""
	epsilon := ""
	for i := 0; i < len(counter); i++ {
		if counter[i]["1"] > counter[i]["0"] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	decGamma, _ := strconv.ParseInt(gamma, 2, 64)
	decEpsilon, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(decGamma) * int(decEpsilon)
}

func solvePart2(data []string) int {
	oxyVals := data
	co2Vals := data
	for pos := 0; pos < len(data[0]); pos++ {
		counterOxy := make(map[string]int, 2)
		for _, line := range oxyVals {
			counterOxy[string(line[pos])] += 1
		}

		counterCO2 := make(map[string]int, 2)
		for _, line := range co2Vals {
			counterCO2[string(line[pos])] += 1
		}

		var mostCommonOxy string
		if counterOxy["1"] > counterOxy["0"] || counterOxy["1"] == counterOxy["0"] {
			mostCommonOxy = "1"
		} else {
			mostCommonOxy = "0"
		}

		var leastCommonCO2 string
		if counterCO2["0"] < counterCO2["1"] || counterCO2["0"] == counterCO2["1"] {
			leastCommonCO2 = "0"
		} else {
			leastCommonCO2 = "1"
		}

		if len(oxyVals) != 1 {
			var newOxyVals []string
			for _, val := range oxyVals {
				if string(val[pos]) == mostCommonOxy {
					newOxyVals = append(newOxyVals, val)
				}
			}
			oxyVals = newOxyVals
		}

		if len(co2Vals) != 1 {
			var newCO2Vals []string
			for _, val := range co2Vals {
				if string(val[pos]) == leastCommonCO2 {
					newCO2Vals = append(newCO2Vals, val)
				}
			}
			co2Vals = newCO2Vals
		}

		if len(oxyVals) == 1 && len(co2Vals) == 1 {
			break
		}
	}

	decOxy, _ := strconv.ParseInt(oxyVals[0], 2, 64)
	decCO2, _ := strconv.ParseInt(co2Vals[0], 2, 64)

	return int(decOxy) * int(decCO2)
}
