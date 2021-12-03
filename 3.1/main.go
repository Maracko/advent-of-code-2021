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

	counter := make(map[int]map[string]int,12)

	for i := 0; i < 12; i++ {
		counter[i] = make(map[string]int,2)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		line := scanner.Text()
		for i,char := range line {
			counter[i][string(char)] += 1
		}
}
	gamma := ""
	for i:=0;i< len(counter);i++ {
		if counter[i]["1"] > counter[i]["0"] {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}

	epsilon := ""
	for _,v := range gamma {
		if string(v) == "1" {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}
	
	fmt.Printf("bin gamma: %s\nbin epsilon: %s\n",gamma,epsilon)

	decGamma, _ := strconv.ParseInt(gamma,2,64)
	decEpsilon, _ := strconv.ParseInt(epsilon,2,64)

	fmt.Printf("dec gamma: %d\ndec epsilon: %d\n",decGamma,decEpsilon)

	fmt.Println("Part 1 solution:", decGamma * decEpsilon)
}
