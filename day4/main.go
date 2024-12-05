package day4

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Symbol int

const (
	X Symbol = iota
	M
	A
	S
)

var letterToSymbolMap = map[byte]Symbol{
	'X': X,
	'M': M,
	'A': A,
	'S': S,
}

// ..X...
// .SAMX.
// .A..A.
// XMAS.S
// .X....
func CountWords(matrix []string) int {
	letterToCoord := map[Symbol][][]int{}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			currentLetter := matrix[y][x]
			letterToSymbol := letterToSymbolMap[currentLetter]
			if _, ok := letterToCoord[letterToSymbol]; !ok {
				letterToCoord[letterToSymbol] = [][]int{}
			}
			letterToCoord[letterToSymbol] = append(letterToCoord[letterToSymbol], []int{x, y})
		}
	}

	count := 0
	for _, coord := range letterToCoord[X] {
		_, directions := SearchForPoint(coord, letterToCoord[M], nil)
		for _, direction := range directions {
			coordM := []int{coord[0] + direction[0], coord[1] + direction[1]}
			isNext, _ := SearchForPoint(coordM, letterToCoord[A], direction)
			if isNext {
				coordA := []int{coordM[0] + direction[0], coordM[1] + direction[1]}
				isNext, _ := SearchForPoint(coordA, letterToCoord[S], direction)
				if isNext {
					count++
				}
			}
		}
	}
	return count
}

func SearchForPoint(current []int, points [][]int, direction []int) (bool, [][]int) {
	// 					[x, y]		[[x1, y1], [x2, y2]]
	x, y := current[0], current[1]
	var potentialDirections [][]int
	if direction == nil {
		potentialDirections = [][]int{
			{-1, -1}, // top-left
			{0, -1},  // top
			{1, -1},  // top-right
			{-1, 0},  // left
			{1, 0},   // right
			{-1, 1},  // bottom-left
			{0, 1},   // bottom
			{1, 1},   // bottom-right
		}
	} else {
		potentialDirections = [][]int{direction}
	}

	matches := [][]int{}
	for _, point := range points {
		for _, potentialDirection := range potentialDirections {
			if point[0] == (x+potentialDirection[0]) && point[1] == (y+potentialDirection[1]) {
				matches = append(matches, potentialDirection)
			}
		}
	}
	return len(matches) > 0, matches
}

func ReadInputFile(filename string) ([]string, error) {
	a := []string{}
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
			return []string{}, err
		}
		line = strings.TrimSuffix(line, "\n")
		a = append(a, line)
	}
	return a, err
}
