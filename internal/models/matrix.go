package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Dontunee/matrix-service/pkg/utils"
)

type Matrix struct {
	numbers [][]string
}

func NewMatrix(input [][]string) *Matrix {
	return &Matrix{numbers: input}
}

// ToString returns the numbers in the matrix as a string
func (m *Matrix) ToString() string {
	return convertToString(m.numbers)
}

// Invert transposes the numbers in the matrix and returns a string
func (m *Matrix) Invert() string {
	rowCount := len(m.numbers)

	if rowCount == 0 {
		return ""
	}

	columnCount := len(m.numbers[0])
	inverted := make([][]string, columnCount)

	for i := 0; i < columnCount; i++ {
		inverted[i] = make([]string, rowCount)
		for j := 0; j < rowCount; j++ {
			inverted[i][j] = m.numbers[j][i]
		}
	}
	return convertToString(inverted)
}

// Flatten returns the numbers in the matrix as a 1 line string
func (m *Matrix) Flatten() string {
	var response []string
	for _, row := range m.numbers {
		response = append(response, row...)
	}

	return strings.Join(response, ",")
}

// Sum adds up all numbers in the given matrix and returns the sum
func (m *Matrix) Sum() (int, error) {
	if len(m.numbers) == 0 || len(m.numbers[0]) == 0 {
		return 0, utils.ErrEmptyMatrix
	}

	numCols := len(m.numbers[0])
	for i := 1; i < len(m.numbers); i++ {
		if len(m.numbers[i]) != numCols {
			return 0, utils.ErrNonRectangularMatrix
		}
	}

	sum := 0
	for i := 0; i < len(m.numbers); i++ {
		for j := 0; j < len(m.numbers[i]); j++ {
			val, err := strconv.Atoi(m.numbers[i][j])
			if err != nil {
				return 0, utils.ErrNonNumericValues
			}
			sum += val
		}
	}

	return sum, nil
}

// Sum multiplies all numbers in the given matrix and returns the value
func (m *Matrix) Multiply() (int, error) {
	if len(m.numbers) == 0 || len(m.numbers[0]) == 0 {
		return 0, utils.ErrEmptyMatrix
	}

	numCols := len(m.numbers[0])
	for i := 1; i < len(m.numbers); i++ {
		if len(m.numbers[i]) != numCols {
			return 0, utils.ErrNonRectangularMatrix
		}
	}

	product := 1
	for i := 0; i < len(m.numbers); i++ {
		for j := 0; j < len(m.numbers[i]); j++ {
			val, err := strconv.Atoi(m.numbers[i][j])
			if err != nil {
				return 0, utils.ErrNonNumericValues
			}
			product *= val
		}
	}

	return product, nil
}

func convertToString(input [][]string) string {
	var response string
	for _, row := range input {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}
