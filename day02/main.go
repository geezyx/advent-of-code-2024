package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var reports [][]int
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		var levels []int
		parts := strings.Fields(line)

		for _, part := range parts {
			level, err := strconv.Atoi(part)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing number: %v\n", err)
				continue
			}
			levels = append(levels, level)
		}

		reports = append(reports, levels)
	}

	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)

	safeCountWithRemoval := 0
	for _, report := range reports {
		if isSafe(report) || canBeMadeSafe(report) {
			safeCountWithRemoval++
		}
	}

	fmt.Printf("Number of safe reports with single removal: %d\n", safeCountWithRemoval)
}

func canBeMadeSafe(levels []int) bool {
	// Create a single reusable slice
	modified := make([]int, len(levels)-1)

	// Try removing each level one at a time
	for i := 0; i < len(levels); i++ {
		// Copy elements before the removed index
		copy(modified[:i], levels[:i])
		// Copy elements after the removed index
		copy(modified[i:], levels[i+1:])

		if isSafe(modified) {
			return true
		}
	}
	return false
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	// Check first difference to determine if we should be increasing or decreasing
	isIncreasing := levels[1] > levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		// Check if difference is between 1 and 3
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}

		// Check if direction matches initial direction
		if isIncreasing && diff <= 0 {
			return false
		}
		if !isIncreasing && diff >= 0 {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
