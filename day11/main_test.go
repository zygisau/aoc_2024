package day11

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBlink(t *testing.T) {
	a := []uint64{0, 1, 10, 99, 999}
	want := []uint64{1, 2024, 1, 0, 9, 9, 2021976}

	out, err := Blink(&a)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(want, out) {
		t.Fatalf("TestBlink()\nwant: %s\n got: %s", fmt.Sprint(want), fmt.Sprint(out))
	}
}

func TestBlinks(t *testing.T) {
	a := []uint64{125, 17}
	wants := [][]uint64{
		{253000, 1, 7},
		{253, 0, 2024, 14168},
		{512072, 1, 20, 24, 28676032},
		{512, 72, 2024, 2, 0, 2, 4, 2867, 6032},
		{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32},
	}

	for i := 0; i < 5; i++ {
		a, err := Blink(&a)
		if err != nil {
			t.Error(err)
		}

		want := wants[i]
		if !reflect.DeepEqual(want, a) {
			t.Fatalf("TestBlinks()\nwant: %s\n got: %s", fmt.Sprint(want), fmt.Sprint(a))
		}
	}
}

func TestReadAndBlink(t *testing.T) {
	filename := "input.txt"
	a, err := ReadFile(filename)
	if err != nil {
		t.Error(err)
	}
	want := 199982

	for i := 0; i < 25; i++ {
		var err error
		a, err = Blink(&a)
		if err != nil {
			t.Error(err)
		}
	}

	if !reflect.DeepEqual(want, len(a)) {
		t.Fatalf("TestReadAndBlink()\nwant: %s\n got: %d", fmt.Sprint(want), len(a))
	}
}
