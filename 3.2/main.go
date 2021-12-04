package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	vals := make([]string, 0, 1000)
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}

	oxyVals := vals
	co2Vals := vals
	for pos := 0; pos < len(vals[0]); pos++ {
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

	fmt.Println("Part 2 solution:", decOxy*decCO2)
}
