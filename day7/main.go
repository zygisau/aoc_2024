package day7

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	Result   uint64
	Operands []uint64
}

func NewOperation(result uint64, operands []uint64) *Operation {
	return &Operation{Result: result, Operands: operands}
}

func NewOperationSpread(result uint64, operands ...uint64) *Operation {
	return &Operation{Result: result, Operands: operands}
}

type OperationFn func(a uint64, b uint64) uint64

func Sum(a uint64, b uint64) uint64 {
	return a + b
}

func Mult(a uint64, b uint64) uint64 {
	return a * b
}

func IsOperationCorrect(operation *Operation) bool {
	guessFns := []OperationFn{Sum, Mult}
	placesCount := len(operation.Operands) - 1
	guessCount := math.Pow(2, float64(placesCount))
	for i := 0; i < int(guessCount); i++ {
		flagBinaryRepresentation := fmt.Sprintf("%%%03db", placesCount)
		flags := fmt.Sprintf(flagBinaryRepresentation, i)
		guess := make([]OperationFn, placesCount)
		for idx, f := range flags {
			opIdx := int(f - '0')
			guess[idx] = guessFns[opIdx]
		}

		answer := uint64(operation.Operands[0])
		for j := 1; j < len(operation.Operands); j++ {
			operand := operation.Operands[j]
			answer = guess[j-1](answer, operand)
		}

		if answer == operation.Result {
			return true
		}
	}
	return false
}

func ReadFromFile(filename string) ([]*Operation, error) {
	operations := []*Operation{}
	f, err := os.Open(filename)
	if err != nil {
		return operations, err
	}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return []*Operation{}, err
		}
		line = strings.TrimSuffix(line, "\n")
		parts := strings.Split(line, ": ")
		resultStr := parts[0]
		val, err := strconv.ParseInt(resultStr, 10, 64)
		if err != nil {
			return []*Operation{}, err
		}
		result := uint64(val)

		operandsStr := strings.Split(parts[1], " ")
		operands := make([]uint64, len(operandsStr))
		for i, operandStr := range operandsStr {
			val, err := strconv.ParseInt(operandStr, 10, 64)
			if err != nil {
				return []*Operation{}, err
			}
			operands[i] = uint64(val)
		}
		operations = append(operations, NewOperationSpread(result, operands...))
	}
	return operations, nil
}
