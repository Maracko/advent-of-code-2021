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

	y := 0
	z := 0
	for scanner.Scan() {
		line := scanner.Text()
		raw := strings.Split(line, " ")

		direction := raw[0]
		count, _ := strconv.Atoi(raw[1])

		switch direction {
		case "forward":
			y += count
		case "up":
			z -= count
		case "down":
			z += count
		}
	}
	fmt.Println("Part 1 solution:", y*z)
}
