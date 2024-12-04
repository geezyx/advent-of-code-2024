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
			input: "3 4\n4 3\n2 5\n1 3\n3 9\n3 3", // Example input
			want1: 11,                             // Expected output for part 1
			want2: 31,                             // Corrected expected output for part 2
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
