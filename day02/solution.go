package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/geezyx/advent-of-code-2024/pkg/utils"
)

type Solution struct {
	input string
}

func New() *Solution {
	return &Solution{}
}

func (s *Solution) LoadInput(path string) error {
	data, err := utils.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}
	s.input = data
	return nil
}

// parseReport converts a string of space-separated numbers into a slice of integers
func parseReport(line string) ([]int, error) {
	var levels []int
	parts := strings.Fields(line)

	for _, part := range parts {
		level, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("parsing number: %w", err)
		}
		levels = append(levels, level)
	}
	return levels, nil
}

// isSafe checks if a report meets the safety criteria:
// 1. All levels are either increasing or decreasing
// 2. Adjacent levels differ by at least 1 and at most 3
func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	// Determine if sequence should be increasing or decreasing
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

func (s *Solution) Part1() (int, error) {
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	safeCount := 0

	for _, line := range lines {
		levels, err := parseReport(line)
		if err != nil {
			return 0, fmt.Errorf("parsing report: %w", err)
		}

		if isSafe(levels) {
			safeCount++
		}
	}

	return safeCount, nil
}

func (s *Solution) Part2() (int, error) {
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	safeCount := 0

	for _, line := range lines {
		levels, err := parseReport(line)
		if err != nil {
			return 0, fmt.Errorf("parsing report: %w", err)
		}

		// Check if report is already safe or can be made safe by removing one number
		if isSafe(levels) || canBeMadeSafe(levels) {
			safeCount++
		}
	}

	return safeCount, nil
}

func main() {
	solution := New()
	if err := solution.LoadInput("input"); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading input: %v\n", err)
		os.Exit(1)
	}

	part1Result, err := solution.Part1()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error solving Part 1: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 result: %d\n", part1Result)

	part2Result, err := solution.Part2()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error solving Part 2: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 result: %d\n", part2Result)
}
