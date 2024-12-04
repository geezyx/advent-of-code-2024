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
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			want1: 2, // 2 reports are safe
			want2: 4, // 4 reports are safe with Problem Dampener
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
