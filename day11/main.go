package day11

import (
	"errors"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func callerName(skip int) string {
	const unknown = "unknown"
	pcs := make([]uintptr, 1)
	n := runtime.Callers(skip+2, pcs)
	if n < 1 {
		return unknown
	}
	frame, _ := runtime.CallersFrames(pcs).Next()
	if frame.Function == "" {
		return unknown
	}
	return frame.Function
}

func timer() func() {
	name := callerName(1)
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func Blink(stones *List[uint64]) (*List[uint64], error) {
	for stone := stones.Front(); stone != nil; stone = stone.Next() {
		switch {
		case stone.Value < 0:
			return nil, errors.New(fmt.Sprintf("Got value: %d. Something is really bad", stone.Value))
		case stone.Value == 0:
			stone.Value = 1
		case isEvenDigits(stone.Value):
			first, second := splitDigits(stone.Value)
			stone.Value = first
			stones.InsertAfter(second, stone)
			stone = stone.Next()
			// stones.PushBack(second)
		default:
			stone.Value = stone.Value * 2024
			break
		}
	}
	return stones, nil
}

func splitDigits(stone uint64) (uint64, uint64) {
	digits := []uint64{}
	for stone > 0 {
		mod := stone % 10
		digits = append(digits, mod)
		stone = stone / 10
	}
	secondHalf := digits[:len(digits)/2]
	firstHalf := digits[len(digits)/2:]

	return JoinBackwardNumbers(firstHalf), JoinBackwardNumbers(secondHalf)
}

func JoinBackwardNumbers(arr []uint64) uint64 {
	mult := uint64(1)
	acc := uint64(0)
	for _, v := range arr {
		acc += v * mult
		mult *= 10
	}
	return acc
}

func isEvenDigits(stone uint64) bool {
	count := math.Floor(math.Log10(float64(stone)) + 1)
	return math.Mod(count, 2) == 0
}

func ReadFile(filename string) (*List[uint64], error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	line := strings.TrimSuffix(string(f), "\n")
	rawStones := strings.Split(line, " ")
	stones := New[uint64]()
	for _, rawStone := range rawStones {
		stone, err := strconv.ParseUint(rawStone, 10, 64)
		if err != nil {
			return nil, err
		}
		stones.PushBack(stone)
	}
	return stones, nil
}
