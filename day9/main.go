package day9

import (
	"os"
)

func DecodeLine(line string) ([]int, []int, error) {
	disk := []int{}
	freeSpaceIdxs := []int{}
	counter := 0
	for i, char := range line {
		num := int(char - '0')

		if i%2 == 0 {
			for j := 0; j < num; j++ {
				disk = append(disk, counter)
			}
			counter++
		} else {
			for j := 0; j < num; j++ {
				disk = append(disk, -1)
				freeSpaceIdxs = append(freeSpaceIdxs, len(disk)-1)
			}
		}
	}
	return disk, freeSpaceIdxs, nil
}

func ReformatDisk(disk []int, freeSpaceIdxs []int) []int {
	for i := len(disk) - 1; i >= 0; i-- {
		val := disk[i]

		if val == -1 {
			continue
		}

		idx := freeSpaceIdxs[0]
		if i <= idx {
			break
		}

		disk[idx] = val
		disk[i] = -1

		freeSpaceIdxs = freeSpaceIdxs[1:]

		if len(freeSpaceIdxs) == 0 {
			break
		}
	}

	return disk
}

func CalcChecksum(disk []int) uint64 {
	sum := uint64(0)
	for i := 0; i < len(disk)-1; i++ {
		val := disk[i]
		if val == -1 {
			break
		}

		sum += uint64(i * val)
	}
	return sum
}

func ReadLine(filename string) (string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(f), nil
}
