package day11

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestBlink(t *testing.T) {
	aRaw := []uint64{0, 1, 10, 99, 999}
	a := New[uint64]()
	for _, aValue := range aRaw {
		a.PushBack(aValue)
	}
	wantRaw := []uint64{1, 2024, 1, 0, 9, 9, 2021976}
	want := New[uint64]()
	for _, wantValue := range wantRaw {
		want.PushBack(wantValue)
	}

	out, err := Blink(a)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(want, out) {
		t.Fatalf("TestBlink()\nwant: %s\n got: %s", fmt.Sprint(want), fmt.Sprint(out))
	}
}

func TestBlinkOneNumber(t *testing.T) {
	aRaw := []uint64{0}
	a := New[uint64]()
	for _, aValue := range aRaw {
		a.PushBack(aValue)
	}

	for i := 0; i < 25; i++ {
		var err error
		a, err = Blink(a)
		if err != nil {
			t.Error(err)
		}
		out := []uint64{}
		for aa := a.Front(); aa != nil; aa = aa.Next() {
			out = append(out, aa.Value)
		}

		err = os.WriteFile(fmt.Sprintf("out/%d.txt", i), []byte(fmt.Sprint(out)), 0644)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestBlinks(t *testing.T) {
	aRaw := []uint64{125, 17}
	a := New[uint64]()
	for _, aValue := range aRaw {
		a.PushBack(aValue)
	}
	wantsRaw := [][]uint64{
		{253000, 1, 7},
		{253, 0, 2024, 14168},
		{512072, 1, 20, 24, 28676032},
		{512, 72, 2024, 2, 0, 2, 4, 2867, 6032},
		{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32},
	}
	wants := []*List[uint64]{}
	for _, wantRaw := range wantsRaw {
		wants = append(wants, New[uint64]())
		l := wants[len(wants)-1]
		for _, wantValue := range wantRaw {
			l.PushBack(wantValue)
		}
	}

	for i := 0; i < 5; i++ {
		a, err := Blink(a)
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
	for workerStone := a.Front(); workerStone != nil; workerStone = workerStone.Next() {
		subsetA := New[uint64]()
		subsetA.PushBack(workerStone.Value)
		go run(subsetA, count, outChannel, t)
	}

	sum := 0
	for i := 0; i < a.Len(); i++ {
		sum += <-outChannel
	}

	if !reflect.DeepEqual(want, sum) {
		t.Fatalf("TestReadAndBlink()\nwant: %s\n got: %d", fmt.Sprint(want), sum)
	}
}

func run(subsetA *List[uint64], count int, outChannel chan int, t *testing.T) {
	for i := 0; i < count; i++ {
		var err error
		subsetA, err = Blink(subsetA)
		if err != nil {
			t.Error(err)
		}
	}
	outChannel <- subsetA.Len()
}
