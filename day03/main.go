package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Read input file
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Regular expression to match valid mul(X,Y) patterns
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(string(data), -1)

	total := 0
	for _, match := range matches {
		// Convert string numbers to integers
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		total += x * y
	}

	fmt.Println("Result:", total)

	// part 2

	// Read input file
	data, err = os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	input := string(data)

	// Combined regex to match all instruction types
	re = regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)

	// Find all instructions in order
	matches = re.FindAllStringSubmatch(input, -1)

	// Initial state: mul instructions are enabled
	mulEnabled := true
	total = 0

	// Process instructions in order
	for _, match := range matches {
		instruction := match[1]

		switch {
		case instruction == "do()":
			mulEnabled = true
		case instruction == "don't()":
			mulEnabled = false
		default: // must be a mul instruction
			if mulEnabled {
				x, _ := strconv.Atoi(match[2])
				y, _ := strconv.Atoi(match[3])
				total += x * y
			}
		}
	}

	fmt.Println("Result:", total)
}
