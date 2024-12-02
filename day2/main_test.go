package day2

import (
	"testing"
)

func TestCompareReports(t *testing.T) {
	a := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	want := []bool{true, false, false, false, false, true}
	out := make([]bool, 6)
	errs := make([]error, 6)
	for i, report := range a {
		out[i], errs[i] = CompareReports(report)
	}

	for i, w := range want {
		if w != out[i] {
			t.Fatalf(`CompareReports() = %t, want %t, got %t; index %dl error: %t`, out[i], w, out[i], i, errs[i])
		}
	}
}

func TestCountSafeReports(t *testing.T) {
	a := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	want := 2
	out := CountSafeReports(a)

	if want != out {
		t.Fatalf(`CompareReports() = %d, want %d, got %d`, out, want, out)
	}
}

func TestReadAndCountSafeReports(t *testing.T) {
	filename := "input.txt"
	want := 21328497
	a, err := ReadInputFile(filename)
	if err != nil {
		t.Error(err)
	}
	out := CountSafeReports(a)

	if want != out {
		t.Fatalf(`CountSafeReports() = %d, want %d, got %d`, out, want, out)
	}
}
