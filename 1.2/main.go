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
		fmt.Println("Cannot open file",err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	vals := make([]int,2000)
	for scanner.Scan(){
		val, _ := strconv.Atoi(scanner.Text())
		vals = append(vals, val)
	}

	timesIncreased := 0
	for i := range vals {
		if i + 3 >= len(vals) {
			break
		}
		
		if (vals[i] + vals[i+1] + vals[i+2]) < (vals[i+1] + vals[i+2] + vals[i+3]) {
			timesIncreased++
		}
	}
	fmt.Println(timesIncreased)
}
