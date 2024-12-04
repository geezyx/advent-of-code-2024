package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const (
	solutionTemplate = `package main

import (
	"fmt"
	"github.com/geezyx/advent-of-code-2024/pkg/utils"
	"os"
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
	// TODO: Implement part 1
	return 0, nil
}

func (s *Solution) Part2() (int, error) {
	// TODO: Implement part 2
	return 0, nil
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
`

	testTemplate = `package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want1   int
		want2   int
		wantErr bool
	}{
		{
			name:  "example",
			input: "", // TODO: Add example input
			want1: 0,  // TODO: Add expected output for part 1
			want2: 0,  // TODO: Add expected output for part 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			s.input = tt.input

			got1, err := s.Part1()
			if (err != nil) != tt.wantErr {
				t.Errorf("Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want1 {
				t.Errorf("Part1() = %v, want %v", got1, tt.want1)
			}

			got2, err := s.Part2()
			if (err != nil) != tt.wantErr {
				t.Errorf("Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got2 != tt.want2 {
				t.Errorf("Part2() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
`

	promptsTemplate = `# Day {{.DayNum}}

## Part 1

Chat history:
` + "```" + `
1. User: [Your prompt]
2. Assistant: [Summary of response]
3. User: [Follow-up prompt]
...
` + "```" + `

Key observations:
- What worked in the prompts
- What needed clarification
- Any interesting AI behaviors

## Part 2

` + "```" + `
1. User: [Your prompt]
2. Assistant: [Summary of response]
3. User: [Follow-up prompt]
...
` + "```" + `

Key observations:
- What worked in the prompts
- What needed clarification
- Any interesting AI behaviors

## Solution Evolution
Document any significant changes or improvements:
- Initial approach
- Problems encountered
- How solutions were refined

## Learnings
- What worked well in the prompt engineering
- What to avoid
- Ideas for future improvements
- Interesting AI capabilities or limitations discovered

## Time Analysis
- Time spent on prompt engineering
- Time spent on implementation
- Comparison with traditional programming approach

## Code Stats
- Lines of code
- Number of prompts needed
- Number of iterations to solution
`
)

type TemplateData struct {
	DayNum string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run scripts/newday.go <day>")
		os.Exit(1)
	}

	dayNum := os.Args[1]
	dayDir := fmt.Sprintf("day%02s", dayNum)

	// Create directory
	if err := os.MkdirAll(dayDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Create files
	files := map[string]string{
		"solution.go":      solutionTemplate,
		"solution_test.go": testTemplate,
		"prompts.md":       promptsTemplate,
		"puzzle.md":        "# Day {{.DayNum}}: [Title]\n\n[Puzzle description goes here]\n",
		"input":            "",
	}

	data := TemplateData{DayNum: dayNum}

	for filename, tmplContent := range files {
		tmpl, err := template.New(filename).Parse(tmplContent)
		if err != nil {
			fmt.Printf("Error parsing template for %s: %v\n", filename, err)
			continue
		}

		filepath := filepath.Join(dayDir, filename)
		if _, err := os.Stat(filepath); err == nil {
			fmt.Printf("File already exists: %s\n", filepath)
			continue
		}

		f, err := os.Create(filepath)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filepath, err)
			continue
		}
		defer f.Close()

		if err := tmpl.Execute(f, data); err != nil {
			fmt.Printf("Error executing template for %s: %v\n", filename, err)
		}
	}

	fmt.Printf("Created new day %s in %s/\n", dayNum, dayDir)
	fmt.Println("Next steps:")
	fmt.Println("1. Add puzzle description to puzzle.md")
	fmt.Println("2. Add puzzle input to input")
	fmt.Println("3. Implement solution in solution.go")
	fmt.Println("4. Add test cases in solution_test.go")
	fmt.Println("5. Document the process in prompts.md")
}
