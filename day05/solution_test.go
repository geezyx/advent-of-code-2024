package main

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
			name: "example",
			input: `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`,
			want1: 143, // Sum of middle numbers (61 + 53 + 29) from correctly ordered updates
			want2: 123, // Sum of middle numbers (47 + 29 + 47) from reordered incorrect updates
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
