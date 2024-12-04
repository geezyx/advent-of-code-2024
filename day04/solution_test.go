package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	s := New()
	s.input = input

	got, err := s.Part1()
	if err != nil {
		t.Errorf("Part1() unexpected error = %v", err)
		return
	}
	want := 18
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

	s := New()
	s.input = input

	got, err := s.Part2()
	if err != nil {
		t.Errorf("Part2() unexpected error = %v", err)
		return
	}
	want := 9
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}
