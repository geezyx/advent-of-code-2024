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
			name:  "example",
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)do()?mul(8,5))",
			want1: 161,
			want2: 48,
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
