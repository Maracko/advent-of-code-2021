package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maracko/advent-of-code-2021/helpers/file"
	"github.com/maracko/advent-of-code-2021/helpers/maps"
	"github.com/maracko/advent-of-code-2021/helpers/slice"
)

type display struct {
	digits []string
	output []string
}

// //Key: digit, Val: number of segments needed to show it
// var digitsMap = map[int]int{
// 	0: 6,
// 	1: 2, //unique len
// 	2: 5,
// 	3: 5,
// 	4: 4, //unique len
// 	5: 5,
// 	6: 6,
// 	7: 3, //unique len
// 	8: 7, //unique len
// 	9: 6,
// }

//Key: Len , Val:  Digit
var uniqueDigitsMap = map[int]int{
	2: 1,
	3: 7,
	4: 4,
	7: 8,
}

var uniqueDigitsLen []int

func init() {
	//Get all unique digits len
	for digitLen := range uniqueDigitsMap {
		uniqueDigitsLen = append(uniqueDigitsLen, digitLen)
	}
}

func main() {
	d, _ := file.ReadFileToSliceOfStrings("data.txt")
	data := getSliceOfDisplays(d)

	fmt.Println("Part 1 solution:", solvePart1(data))
	fmt.Println("Part 2 solution:", solvePart2(data))
}

func getSliceOfDisplays(data []string) []display {
	var res []display
	for _, val := range data {
		splitVal := strings.Split(val, " ")

		digits := append([]string(nil), splitVal[:10]...)
		digits = slice.SortAllStringsInSlice(digits)

		output := append([]string(nil), splitVal[11:]...)
		output = slice.SortAllStringsInSlice(output)

		res = append(res, display{digits, output})
	}
	return res
}

func solvePart1(data []display) int {
	res := 0
	for _, d := range data {
		for _, output := range d.output {
			if slice.IntInSlice(uniqueDigitsLen, len(output)) {
				res++
			}
		}
	}
	return res
}

func solvePart2(data []display) int {
	res := 0
	for _, disp := range data {
		vals := append(disp.digits, disp.output...)
		valsMap := make(map[int]string)
		// Find the "easy digits"
		for _, val := range vals {
			size := len(val)
			if slice.IntInSlice(uniqueDigitsLen, size) {
				valsMap[uniqueDigitsMap[size]] = val
			}
		}

		// Find 6 digit
		for _, val := range vals {
			if len(val) == 6 {
				for _, char := range valsMap[1] {
					if !strings.ContainsRune(val, char) {
						valsMap[6] = val
						break
					}
				}
			}
		}

		//Find 0 digit
		for _, val := range vals {
			if len(val) == 6 {
				for _, char := range valsMap[4] {
					if !strings.ContainsRune(val, char) && !maps.StringInMap(valsMap, val) {
						valsMap[0] = val
						break
					}
				}
			}
		}

		//Find 9 after 6 and 0 with length 6
		for _, val := range vals {
			if len(val) == 6 && !maps.StringInMap(valsMap, val) {
				valsMap[9] = val
			}
		}

	Value5Loop:
		//Find 5
		for _, val := range vals {
			if len(val) == 5 {
				for i, char := range val {
					if !strings.ContainsRune(valsMap[6], char) {
						continue Value5Loop
					}
					if i == len(val)-1 {
						valsMap[5] = val
					}
				}
			}
		}

	Value3Loop:
		//Find 3
		for _, val := range vals {
			if len(val) == 5 {
				for _, char := range val {
					if !strings.ContainsRune(valsMap[9], char) {
						continue Value3Loop
					}
				}

				if !maps.StringInMap(valsMap, val) {
					valsMap[3] = val
					fmt.Println(valsMap[3])
					break Value3Loop
				}
			}
		}

		//Find 2 after 3 and 5 with length 5
		for _, val := range vals {
			if len(val) == 5 && !maps.StringInMap(valsMap, val) {
				valsMap[2] = val
			}
		}

		//Calculate the number inside output and add it to total
		outString := ""
		for _, output := range disp.output {
			for number, digit := range valsMap {
				if output == digit {
					outString += fmt.Sprint(number)
					break
				}
			}
		}
		num, _ := strconv.Atoi(outString)
		res += num
	}
	return res
}
