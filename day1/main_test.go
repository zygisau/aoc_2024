package day1

import (
	"sort"
	"testing"
)

func TestGetDistanceOfPair_shouldCompareTwoNumbers(t *testing.T) {
	a, b := int(5), int(10)
	want := b - a
	out := Day1.GetDistanceOfPair(a, b)
	if want != out {
		t.Fatalf(`GetDistanceOfPair() = %d, want %d, got %d`, out, want, out)
	}
}

func TestCalculateDistance(t *testing.T) {
	a, b := []int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}
	want := 11

	sort.Ints(a)
	sort.Ints(b)
	calculateDistance := Day1.CalculateDistance(a, b)
	for i := 0; i < len(a); i++ {
		itemA := a[i]
		itemB := b[i]
		calculateDistance(itemA, itemB)
	}

	out := calculateDistance(0, 0)
	if want != out {
		t.Fatalf(`CalculateDistance() = %d, want %d, got %d`, out, want, out)
	}
}

func TestReadInputFileAndCalculateDistance(t *testing.T) {
	filename := "input.txt"
	want := 2742123
	a, b, err := Day1.ReadInputFile(filename)
	if err != nil {
		t.Error(err)
	}

	sort.Ints(a)
	sort.Ints(b)
	calculateDistance := Day1.CalculateDistance(a, b)
	for i := 0; i < len(a); i++ {
		itemA := a[i]
		itemB := b[i]
		calculateDistance(itemA, itemB)
	}

	out := calculateDistance(0, 0)
	if want != out {
		t.Fatalf(`CalculateDistance() = %d, want %d, got %d`, out, want, out)
	}
}

func TestGetCalculateSimilarity(t *testing.T) {
	a, b := []int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}
	want := 31

	sort.Ints(a)
	sort.Ints(b)
	appendToSimilarity, calculateScore := Day1.GetCalculateSimilarity()
	for i := 0; i < len(a); i++ {
		itemA := a[i]
		itemB := b[i]
		appendToSimilarity(itemA, itemB)
	}

	out := calculateScore()
	if want != out {
		t.Fatalf(`CalculateSimilarity() = %d, want %d, got %d`, out, want, out)
	}
}

func TestReadInputFileAndCalculateSimilarity(t *testing.T) {
	filename := "input.txt"
	want := 21328497
	a, b, err := Day1.ReadInputFile(filename)
	if err != nil {
		t.Error(err)
	}

	sort.Ints(a)
	sort.Ints(b)
	appendToSimilarity, calculateScore := Day1.GetCalculateSimilarity()
	for i := 0; i < len(a); i++ {
		itemA := a[i]
		itemB := b[i]
		appendToSimilarity(itemA, itemB)
	}

	out := calculateScore()
	if want != out {
		t.Fatalf(`CalculateSimilarity() = %d, want %d, got %d`, out, want, out)
	}
}
