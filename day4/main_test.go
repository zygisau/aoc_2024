package day4

import (
	"testing"
)

func TestCountWords(t *testing.T) {
	// a := []string{
	// 	"MMMSXXMASM",
	// 	"MSAMXMSMSA",
	// 	"AMXSXMAAMM",
	// 	"MSAMASMSMX",
	// 	"XMASAMXAMM",
	// 	"XXAMMXXAMA",
	// 	"SMSMSASXSS",
	// 	"SAXAMASAAA",
	// 	"MAMMMXMMMM",
	// 	"MXMXAXMASX",
	// }
	a := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	want := 18
	out := CountWords(a)
	if want != out {
		t.Fatalf(`TestCountWords() = %d, want %d, got %d`, out, want, out)
	}
}

func TestReadFileCountWords(t *testing.T) {
	filename := "input.txt"
	a, err := ReadInputFile(filename)
	if err != nil {
		t.Error(err)
	}

	want := 2493
	out := CountWords(a)
	if want != out {
		t.Fatalf(`TestCountWords() = %d, want %d, got %d`, out, want, out)
	}
}
