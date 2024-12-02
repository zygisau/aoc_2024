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

func IntAbs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func CompareReports(reports []int) (bool, error) {
	continuityFactor := 0
	isContinuous := true
	var prev *int
	var next *int
	for i, report := range reports {
		if i-1 >= 0 {
			prev = &reports[i-1]
		}

		if i+1 < len(reports) {
			next = &reports[i+1]
		}

		if next != nil {
			previousLevelsContinuity := continuityFactor
			if (*next - report) != 0 {
				continuityFactor = (*next - report) / IntAbs(*next, report)
			}
			previousIsContinuous := isContinuous
			if prev != nil {
				isContinuous = previousLevelsContinuity == continuityFactor
			}

			if previousIsContinuous != isContinuous {
				return false, errors.New("not continuous")
			}
		}

		isAdjacentDifferOneToThree := []bool{true, true}
		if prev != nil {
			diff := IntAbs(*prev, report)
			isAdjacentDifferOneToThree[0] = diff >= 1 && diff <= 3
		}
		if next != nil {
			diff := IntAbs(*next, report)
			isAdjacentDifferOneToThree[1] = diff >= 1 && diff <= 3
		}
		if !isAdjacentDifferOneToThree[0] || !isAdjacentDifferOneToThree[1] {
			return false, errors.New("adjacent values differ less than 1 or more than 3")
		}

		if prev != nil {
			prev = nil
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

func CountSafeReports(reportLogs [][]int) int {
	out := make([]bool, len(reportLogs))
	for i, reports := range reportLogs {
		out[i], _ = CompareReports(reports)
	}
	return SumArray(out)
}

func ReadInputFile(filename string) ([][]int, error) {
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