package day6

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Coords struct {
	X int
	Y int
}

func ReadMap(filename string) ([][]bool, []int, error) {
	mapData, pos := [][]bool{}, make([]int, 2)
	f, err := os.Open(filename)
	if err != nil {
		return mapData, pos, err
	}

	reader := bufio.NewReader(f)
	col := 0
	for {
		line, err := reader.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return [][]bool{}, []int{}, err
		}
		line = strings.TrimSuffix(line, "\n")
		mapData = append(mapData, make([]bool, len(line)))
		for i, str := range line {
			switch str {
			case '#':
				mapData[col][i] = true
			case '^':
				pos[0] = i
				pos[1] = col
			default:
				mapData[col][i] = false
			}
		}
		col++
	}
	return mapData, pos, err
}

type Key struct {
	X int
	Y int
}

func CountSteps(mapData [][]bool, position []int) int {
	uniquePlaces := make(map[Key]struct{})
	direction := []int{0, -1}
	for position[0] >= 0 && position[0] < len(mapData[0]) && position[1] >= 0 && position[1] < len(mapData) {
		uniquePlaces[Key{X: position[0], Y: position[1]}] = struct{}{}
		newPosition := []int{position[0] + direction[0], position[1] + direction[1]}
		if newPosition[0] >= 0 && newPosition[0] < len(mapData[0]) && newPosition[1] >= 0 && newPosition[1] < len(mapData) {
			if val := mapData[newPosition[1]][newPosition[0]]; val {
				// rotation 90 deg clockwise
				oldDir := direction[0]
				direction[0] = -direction[1]
				direction[1] = oldDir
			}
		}
		position[0] += direction[0]
		position[1] += direction[1]
	}
	return len(uniquePlaces)
}

func PrintMap(mapData [][]bool, locations map[Key]struct{}, player []int) string {
	projection := ""
	for y, col := range mapData {
		for x, isObstacle := range col {
			if _, ok := locations[Key{X: x, Y: y}]; ok {
				projection += "X"
			} else if isObstacle {
				projection += "#"
			} else if x == player[0] && y == player[1] {
				projection += "^"
			} else {
				projection += "."
			}
		}
		projection += "\n"
	}
	return projection
}
