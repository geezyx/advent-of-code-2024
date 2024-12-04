package main

import (
	"fmt"
	"os"
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

// findXMAS searches for "XMAS" in all possible directions from a given starting point
func findXMAS(grid [][]rune, row, col int) int {
	// Check if grid is empty or has invalid dimensions
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// Validate row and column bounds
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
		return 0
	}

	// All possible directions: right, down-right, down, down-left, left, up-left, up, up-right
	directions := [][2]int{
		{0, 1},   // right
		{1, 1},   // down-right
		{1, 0},   // down
		{1, -1},  // down-left
		{0, -1},  // left
		{-1, -1}, // up-left
		{-1, 0},  // up
		{-1, 1},  // up-right
	}

	count := 0
	target := "XMAS"

	// Check each direction
	for _, dir := range directions {
		valid := true
		// Pre-check if the word would fit in this direction
		lastRow := row + (len(target)-1)*dir[0]
		lastCol := col + (len(target)-1)*dir[1]

		// Skip if the word wouldn't fit in this direction
		if lastRow < 0 || lastRow >= len(grid) ||
			lastCol < 0 || lastCol >= len(grid[row]) {
			continue
		}

		for i := 0; i < len(target); i++ {
			newRow := row + i*dir[0]
			newCol := col + i*dir[1]

			if newRow < 0 || newRow >= len(grid) ||
				newCol < 0 || newCol >= len(grid[newRow]) ||
				grid[newRow][newCol] != rune(target[i]) {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
	}

	return count
}

// findXMAS_Part2 searches for X-shaped MAS patterns
func findXMAS_Part2(grid [][]rune, row, col int) int {
	// Check if grid is empty or has invalid dimensions
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// Validate row and column bounds (need room for 3x3 pattern)
	if row-1 < 0 || row+1 >= len(grid) || col-1 < 0 || col+1 >= len(grid[0]) {
		return 0
	}

	// Check if the center is 'A'
	if grid[row][col] != 'A' {
		return 0
	}

	count := 0

	// Pattern examples:
	// M . S    S . M    M . M    S . S
	// . A .    . A .    . A .    . A .
	// M . S    S . M    S . S    M . M

	// Check all four corners
	topLeft := grid[row-1][col-1]
	topRight := grid[row-1][col+1]
	bottomLeft := grid[row+1][col-1]
	bottomRight := grid[row+1][col+1]

	// Check all valid combinations
	isValid := func(start1, end1, start2, end2 rune) bool {
		return ((start1 == 'M' && end1 == 'S') || (start1 == 'S' && end1 == 'M')) &&
			((start2 == 'M' && end2 == 'S') || (start2 == 'S' && end2 == 'M'))
	}

	// Check diagonal combinations
	if isValid(topLeft, bottomRight, topRight, bottomLeft) {
		count++
	}
	if isValid(topRight, bottomLeft, topLeft, bottomRight) {
		count++
	}

	return count / 2 // Each valid pattern is counted twice due to symmetry
}

func (s *Solution) Part1() (int, error) {
	// Split input into lines and create grid
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	if len(lines) == 0 {
		return 0, fmt.Errorf("empty input")
	}

	// Convert input to 2D rune grid
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	// Search for XMAS starting from each position
	total := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			total += findXMAS(grid, row, col)
		}
	}

	return total, nil
}

func (s *Solution) Part2() (int, error) {
	// Split input into lines and create grid
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	if len(lines) == 0 {
		return 0, fmt.Errorf("empty input")
	}

	// Convert input to 2D rune grid
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	// Search for X-MAS patterns starting from each position
	total := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			total += findXMAS_Part2(grid, row, col)
		}
	}

	return total, nil
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
