package day6

import "testing"

func TestFindPathFromFile(t *testing.T) {
	filename := "input.txt"
	a, pos, err := ReadMap(filename)
	if err != nil {
		t.Error(err)
	}

	want := 4665
	out := CountSteps(a, pos)
	if want != out {
		t.Fatalf(`TestCountSteps() = %d, want %d, got %d`, out, want, out)
	}
}
