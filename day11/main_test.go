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

	if !reflect.DeepEqual(want, *out) {
		t.Fatalf("TestBlink()\nwant: %s\n got: %s", fmt.Sprint(want), fmt.Sprint(*out))
	}
}

func TestBlinks(t *testing.T) {
	original := []uint64{125, 17}
	wants := [][]uint64{
		{253000, 1, 7},
		{253, 0, 2024, 14168},
		{512072, 1, 20, 24, 28676032},
		{512, 72, 2024, 2, 0, 2, 4, 2867, 6032},
		{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32},
	}

	a := &original
	for i := 0; i < 5; i++ {
		a, err := Blink(a)
		if err != nil {
			t.Error(err)
		}

		want := wants[i]
		if !reflect.DeepEqual(want, *a) {
			t.Fatalf("TestBlinks()\nwant: %s\n got: %s", fmt.Sprint(want), fmt.Sprint(*a))
		}
	}
}

func TestReadAndBlink(t *testing.T) {
	count := 25
	want := 199982
	RunReadAndBlinkTest(t, count, want)
}

func TestReadAndBlinkMore(t *testing.T) {
	count := 75
	want := 199982
	RunReadAndBlinkTest(t, count, want)
}

func RunReadAndBlinkTest(t *testing.T, count int, want int) {
	filename := "input.txt"
	a, err := ReadFile(filename)
	if err != nil {
		t.Error(err)
	}

	outChannel := make(chan int)
	for workerIdx := 0; workerIdx < len(a); workerIdx++ {
		subsetA := []uint64{a[workerIdx]}
		go run(&subsetA, count, outChannel, t)
	}

	sum := 0
	for i := 0; i < len(a); i++ {
		sum += <-outChannel
	}

	if !reflect.DeepEqual(want, sum) {
		t.Fatalf("TestReadAndBlink()\nwant: %s\n got: %d", fmt.Sprint(want), sum)
	}
}

func run(subsetA *[]uint64, count int, outChannel chan int, t *testing.T) {
	for i := 0; i < count; i++ {
		var err error
		subsetA, err = Blink(subsetA)
		if err != nil {
			t.Error(err)
		}
	}
	outChannel <- len(*subsetA)
}
