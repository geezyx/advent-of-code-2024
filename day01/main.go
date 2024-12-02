package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	var firstInts, secondInts []int

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Invalid line format: %s\n", line)
			continue
		}

		firstInt, err1 := strconv.Atoi(parts[0])
		secondInt, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Fprintf(os.Stderr, "Error parsing integers: %s\n", line)
			continue
		}

		firstInts = append(firstInts, firstInt)
		secondInts = append(secondInts, secondInt)
	}

	sort.Ints(firstInts)
	sort.Ints(secondInts)

	var totalDistance int
	for i := 0; i < len(firstInts); i++ {
		distance := abs(firstInts[i] - secondInts[i])
		totalDistance += distance
	}

	fmt.Printf("Sorted first integers: %v\n", firstInts)
	fmt.Printf("Sorted second integers: %v\n", secondInts)
	fmt.Printf("Total distance: %d\n", totalDistance)

	freqMap := make(map[int]int)
	for _, num := range secondInts {
		freqMap[num]++
	}

	// Calculate similarity score
	var totalSimilarity int
	for _, num := range firstInts {
		frequency := freqMap[num]
		totalSimilarity += num * frequency
	}

	fmt.Printf("Total similarity score: %d\n", totalSimilarity)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
