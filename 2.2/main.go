package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	depth := 0
	forward := 0
	aim := 0
	for scanner.Scan(){
		line := scanner.Text()
		raw := strings.Split(line," ")

		direction := raw[0]
		count, _ := strconv.Atoi(raw[1])
		
		switch direction{
		case "forward":
			forward += count
			depth += count * aim
		case "up":
			aim -= count
		case "down":
			aim += count
		}
		
}
fmt.Println("Part 2 solution:",forward*depth)
}