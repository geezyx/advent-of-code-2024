package main

import (
	"fmt"
	"os"
	"sort"
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

func (s *Solution) Part1() (int, error) {
	// Split input into lines and create two slices for numbers
	var firstList, secondList []int

	// Process each line
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		// Split line into two numbers
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid line format: %s", line)
		}

		// Parse first number
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, fmt.Errorf("parsing first number: %w", err)
		}

		// Parse second number
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, fmt.Errorf("parsing second number: %w", err)
		}

		firstList = append(firstList, num1)
		secondList = append(secondList, num2)
	}

	// Sort both lists
	sort.Ints(firstList)
	sort.Ints(secondList)

	// Calculate total distance
	totalDistance := 0
	for i := 0; i < len(firstList); i++ {
		distance := abs(firstList[i] - secondList[i])
		totalDistance += distance
	}

	return totalDistance, nil
}

// Helper function to calculate absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (s *Solution) Part2() (int, error) {
	// Split input into lines and create two slices for numbers
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	var firstList, secondList []int

	// Process each line
	for _, line := range lines {
		if line == "" {
			continue
		}

		// Split line into two numbers
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid line format: %s", line)
		}

		// Parse first number
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, fmt.Errorf("parsing first number: %w", err)
		}

		// Parse second number
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, fmt.Errorf("parsing second number: %w", err)
		}

		firstList = append(firstList, num1)
		secondList = append(secondList, num2)
	}

	// Create frequency map for the right list
	freqMap := make(map[int]int)
	for _, num := range secondList {
		freqMap[num]++
	}

	// Calculate similarity score
	totalSimilarity := 0
	for _, num := range firstList {
		frequency := freqMap[num]
		totalSimilarity += num * frequency
	}

	return totalSimilarity, nil
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
