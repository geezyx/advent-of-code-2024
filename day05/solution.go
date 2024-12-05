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

type Rule struct {
	before int
	after  int
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

// parseInput parses the input into rules and updates
func parseInput(input string) ([]Rule, [][]int, error) {
	// Ensure consistent newline format
	input = strings.ReplaceAll(input, "\r\n", "\n")

	// Split input into rules and updates sections
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(sections) != 2 {
		return nil, nil, fmt.Errorf("invalid input format")
	}

	// Parse rules
	var rules []Rule
	for _, line := range strings.Split(sections[0], "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid rule format: %s", line)
		}
		before, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, nil, fmt.Errorf("parsing before number: %w", err)
		}
		after, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, nil, fmt.Errorf("parsing after number: %w", err)
		}
		rules = append(rules, Rule{before, after})
	}

	// Parse updates
	var updates [][]int
	for _, line := range strings.Split(sections[1], "\n") {
		if line == "" {
			continue
		}
		var update []int
		for _, numStr := range strings.Split(line, ",") {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return nil, nil, fmt.Errorf("parsing update number: %w", err)
			}
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	return rules, updates, nil
}

// isValidOrder checks if the update follows all applicable rules
func isValidOrder(update []int, rules []Rule) bool {
	// Create a map of number positions for quick lookup
	positions := make(map[int]int)
	for i, num := range update {
		positions[num] = i
	}

	// Check each rule
	for _, rule := range rules {
		// Skip rules that don't apply to this update
		beforePos, beforeExists := positions[rule.before]
		afterPos, afterExists := positions[rule.after]
		if !beforeExists || !afterExists {
			continue
		}
		// If both numbers exist, check their order
		if beforePos >= afterPos {
			return false
		}
	}
	return true
}

// topologicalSort performs a topological sort of numbers based on the rules
func topologicalSort(numbers []int, rules []Rule) ([]int, error) {
	// Create adjacency list
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	// Initialize all numbers in the maps
	for _, num := range numbers {
		graph[num] = []int{}
		inDegree[num] = 0
	}

	// Build the graph
	for _, rule := range rules {
		// Only consider rules where both numbers are in our set
		if _, hasBefore := graph[rule.before]; hasBefore {
			if _, hasAfter := graph[rule.after]; hasAfter {
				graph[rule.before] = append(graph[rule.before], rule.after)
				inDegree[rule.after]++
			}
		}
	}

	// Find all nodes with no incoming edges
	var queue []int
	for num := range graph {
		if inDegree[num] == 0 {
			queue = append(queue, num)
		}
	}

	// Process the queue
	var result []int
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check if we have a valid ordering
	if len(result) != len(numbers) {
		return nil, fmt.Errorf("cycle detected in dependencies")
	}

	return result, nil
}

func (s *Solution) Part1() (int, error) {
	// Parse input
	rules, updates, err := parseInput(s.input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	// Process each update
	sum := 0
	for _, update := range updates {
		if isValidOrder(update, rules) {
			// Get middle number
			middleIndex := len(update) / 2
			sum += update[middleIndex]
		}
	}

	return sum, nil
}

func (s *Solution) Part2() (int, error) {
	// Parse input
	rules, updates, err := parseInput(s.input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %w", err)
	}

	sum := 0
	for _, update := range updates {
		// Skip updates that are already in correct order
		if isValidOrder(update, rules) {
			continue
		}

		// Reorder the update
		ordered, err := topologicalSort(update, rules)
		if err != nil {
			return 0, fmt.Errorf("reordering update: %w", err)
		}

		// Add middle number to sum
		middleIndex := len(ordered) / 2
		sum += ordered[middleIndex]
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
