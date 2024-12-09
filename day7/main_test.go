package day7

import (
	"testing"
)

func TestIsOperationCorrect(t *testing.T) {
	a := []*Operation{
		NewOperationSpread(190, 10, 19),
		NewOperationSpread(3267, 81, 40, 27),
		NewOperationSpread(83, 17, 5),
		NewOperationSpread(156, 15, 6),
		NewOperationSpread(7290, 6, 8, 6, 15),
		NewOperationSpread(161011, 16, 10, 13),
		NewOperationSpread(192, 17, 8, 14),
		NewOperationSpread(21037, 9, 7, 18, 13),
		NewOperationSpread(292, 11, 6, 16, 20),
	}
	want := []bool{true, true, false, false, false, false, false, false, true}

	for i, aa := range a {
		w := want[i]
		out := IsOperationCorrect(aa)

		if out != w {
			t.Fatalf(`IsOperationCorrect() = %t, want %t, got %t`, out, w, out)
		}
	}
}

func TestReadAndCheckIsOperationCorrect(t *testing.T) {
	filename := "input.txt"
	want := 0
	a, err := ReadFromFile(filename)
	if err != nil {
		t.Error(err)
	}

	sum := 0
	for _, aa := range a {
		isCorrect := IsOperationCorrect(aa)

		if isCorrect {
			sum += int(aa.Result)
		}
	}

	if sum != want {
		t.Fatalf(`IsOperationCorrect() = %d, want %d, got %d`, sum, want, sum)
	}
}
