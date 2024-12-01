package day1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Day1T struct{}

var Day1 Day1T

func (Day1T) GetDistanceOfPair(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

type CalculateDistanceFn func(a int, b int) int

func (Day1T) CalculateDistance(listA []int, listB []int) CalculateDistanceFn {
	sum := 0
	return func(a, b int) int {
		sum += Day1.GetDistanceOfPair(a, b)
		return sum
	}
}

func (Day1T) ReadInputFile(filename string) ([]int, []int, error) {
	a, b := []int{}, []int{}
	f, err := os.Open(filename)
	if err != nil {
		return a, b, err
	}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return []int{}, []int{}, err
		}
		line = strings.TrimSuffix(line, "\n")
		splitted := strings.Split(line, "   ")
		if len(splitted) != 2 {
			return []int{}, []int{}, fmt.Errorf("splitted string '%v' does not contain two integers", line)
		}

		aStr, err := strconv.Atoi(splitted[0])
		if err != nil {
			return []int{}, []int{}, fmt.Errorf("a string '%v' cannot be casted to integer", splitted[0])
		}
		bStr, err := strconv.Atoi(splitted[1])
		if err != nil {
			return []int{}, []int{}, fmt.Errorf("b string '%v' cannot be casted to integer", splitted[1])
		}

		a = append(a, aStr)
		b = append(b, bStr)
	}
	return a, b, err
}

type CalculateFn func(a int, b int)
type GetCalculationFn func() int

func (Day1T) GetCalculateSimilarity() (CalculateFn, GetCalculationFn) {
	leftSideFrequency := map[int]int{}
	rightSideFrequency := map[int]int{}
	return func(a int, b int) {
			if _, ok := leftSideFrequency[a]; !ok {
				leftSideFrequency[a] = 0
			}
			leftSideFrequency[a]++

			if _, ok := rightSideFrequency[b]; !ok {
				rightSideFrequency[b] = 0
			}
			rightSideFrequency[b]++
		}, func() int {
			sum := 0
			for a, freqA := range leftSideFrequency {
				freqB := rightSideFrequency[a]
				sum += freqA * (a * freqB)
			}
			return sum
		}
}
