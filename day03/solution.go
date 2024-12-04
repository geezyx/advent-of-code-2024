package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

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

// extractValidMultiplications finds all valid mul(X,Y) instructions and returns their products
func extractValidMultiplications(input string) ([]int, error) {
	// Regular expression to match valid mul(X,Y) patterns
	// - Matches 'mul' followed by exactly one opening parenthesis
	// - Captures 1-3 digits for first number
	// - Matches comma
	// - Captures 1-3 digits for second number
	// - Matches exactly one closing parenthesis
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches in the input
	matches := re.FindAllStringSubmatch(input, -1)

	var results []int
	for _, match := range matches {
		// Parse first number
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, fmt.Errorf("parsing first number: %w", err)
		}

		// Parse second number
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, fmt.Errorf("parsing second number: %w", err)
		}

		// Calculate product and add to results
		results = append(results, num1*num2)
	}

	return results, nil
}

// instruction represents a parsed instruction from the input
type instruction struct {
	typ  string // "mul", "do", or "don't"
	num1 int    // first number for mul instructions
	num2 int    // second number for mul instructions
	pos  int    // position in input string
}

// parseInstructions finds all valid instructions in order
func parseInstructions(input string) ([]instruction, error) {
	// Combined regex to match both mul and conditional instructions
	re := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)

	matches := re.FindAllStringSubmatchIndex(input, -1)
	var instructions []instruction

	for _, match := range matches {
		fullMatch := input[match[0]:match[1]]

		switch {
		case fullMatch == "do()":
			instructions = append(instructions, instruction{
				typ: "do",
				pos: match[0],
			})
		case fullMatch == "don't()":
			instructions = append(instructions, instruction{
				typ: "don't",
				pos: match[0],
			})
		default:
			// Must be a mul instruction
			mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
			mulMatch := mulRe.FindStringSubmatch(fullMatch)
			if mulMatch == nil {
				continue
			}

			num1, err := strconv.Atoi(mulMatch[1])
			if err != nil {
				return nil, fmt.Errorf("parsing first number: %w", err)
			}

			num2, err := strconv.Atoi(mulMatch[2])
			if err != nil {
				return nil, fmt.Errorf("parsing second number: %w", err)
			}

			instructions = append(instructions, instruction{
				typ:  "mul",
				num1: num1,
				num2: num2,
				pos:  match[0],
			})
		}
	}

	return instructions, nil
}

func (s *Solution) Part1() (int, error) {
	// Extract all valid multiplications
	products, err := extractValidMultiplications(s.input)
	if err != nil {
		return 0, fmt.Errorf("extracting multiplications: %w", err)
	}

	// Sum all products
	sum := 0
	for _, product := range products {
		sum += product
	}

	return sum, nil
}

func (s *Solution) Part2() (int, error) {
	// Parse all instructions in order
	instructions, err := parseInstructions(s.input)
	if err != nil {
		return 0, fmt.Errorf("parsing instructions: %w", err)
	}

	// Process instructions in order
	enabled := true // mul instructions are enabled by default
	sum := 0

	for _, inst := range instructions {
		switch inst.typ {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if enabled {
				sum += inst.num1 * inst.num2
			}
		}
	}

	return sum, nil
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
