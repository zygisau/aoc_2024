package day2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Day2T struct{}

var Day2 Day2T

type CompareError struct {
	Error error
	Idx   int
}

var Error CompareError

func (e *CompareError) New(message string, idx int) *CompareError {
	return &CompareError{Error: errors.New(message), Idx: idx}
}

func (e *CompareError) String() string {
	return fmt.Sprintf("Error: %q, index: %d", e.Error, Error.Idx)
}

func IntAbs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func CheckContinuity(a int, b int) int {
	diff := a - b
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

type CompareDelegateFn func(reports []int) (bool, *CompareError)

func (Day2T) CompareReports(reports []int) (bool, *CompareError) {
	continuityFactor := 0
	if len(reports) > 1 {
		continuityFactor = CheckContinuity(reports[1], reports[0])
		diff := IntAbs(reports[1], reports[0])
		isAdjacentDifferOneToThree := diff >= 1 && diff <= 3
		if !isAdjacentDifferOneToThree {
			return false, Error.New("adjacent values differ less than 1 or more than 3", 1)
		}
	}
	var next *int

	for i := 1; i < len(reports); i++ {
		current := reports[i]

		if i+1 < len(reports) {
			next = &reports[i+1]
		}

		if next != nil {
			previousLevelsContinuity := continuityFactor
			continuityFactor = CheckContinuity(*next, current)

			if previousLevelsContinuity != continuityFactor {
				return false, Error.New("not continuous", i+1)
			}
		}

		isAdjacentDifferOneToThree := true
		if next != nil {
			diff := IntAbs(*next, current)
			isAdjacentDifferOneToThree = diff >= 1 && diff <= 3
		}
		if !isAdjacentDifferOneToThree {
			return false, Error.New("adjacent values differ less than 1 or more than 3", i+1)
		}

		if next != nil {
			next = nil
		}
	}
	return true, nil
}

func Bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func SumArray(arr []bool) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		result += Bool2int(arr[i])
	}
	return result
}

func (Day2T) CountSafeReports(reportLogs [][]int, CompareDelegate CompareDelegateFn) int {
	out := make([]bool, len(reportLogs))
	for i, reports := range reportLogs {
		out[i], _ = CompareDelegate(reports)
	}
	return SumArray(out)
}

func (Day2T) ReadInputFile(filename string) ([][]int, error) {
	a := [][]int{}
	f, err := os.Open(filename)
	if err != nil {
		return a, err
	}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return [][]int{}, err
		}
		line = strings.TrimSuffix(line, "\n")
		splitted := strings.Split(line, " ")

		a = append(a, make([]int, len(splitted)))

		for i, str := range splitted {
			integer, err := strconv.Atoi(str)
			if err != nil {
				return [][]int{}, fmt.Errorf("a string '%v' cannot be casted to integer", splitted[0])
			}
			a[len(a)-1][i] = integer
		}
	}
	return a, err
}

func SaveRemove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func (Day2T) CompareReportsWithFix(reports []int) (bool, *CompareError) {
	out, err := Day2.CompareReports(reports)
	if err != nil {
		reports = SaveRemove(reports, err.Idx)
		secondOut, secondErr := Day2.CompareReports(reports)
		return secondOut, secondErr
	}
	return out, err
}
