package day11

import (
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Blink(stones *[]uint64) (*[]uint64, error) {
	for i := 0; i < len(*stones); i++ {
		stone := (*stones)[i]

		switch {
		case stone == 0:
			(*stones)[i] = 1
		case isEvenDigits(stone):
			first, second := splitDigits(stone)
			(*stones)[i] = first
			*stones = slices.Insert(*stones, i+1, second)
			i++
		default:
			(*stones)[i] = (*stones)[i] * 2024
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

func ReadFile(filename string) ([]uint64, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	line := strings.TrimSuffix(string(f), "\n")
	rawStones := strings.Split(line, " ")
	stones := []uint64{}
	for _, rawStone := range rawStones {
		stone, err := strconv.ParseUint(rawStone, 10, 64)
		if err != nil {
			return nil, err
		}
		stones = append(stones, stone)
	}
	return stones, nil
}
