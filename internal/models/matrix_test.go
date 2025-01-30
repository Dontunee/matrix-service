package models

import (
	"testing"

	"github.com/Dontunee/matrix-service/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {

	tests := []struct {
		name    string
		matrix  [][]string
		want    int
		wantErr error
	}{
		{
			"success",
			[][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			45,
			nil,
		},
		{
			"empty matrix",
			[][]string{},
			0,
			utils.ErrEmptyMatrix,
		},
		{
			"non rectangular matrix",
			[][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "9"},
			},
			0,
			utils.ErrNonRectangularMatrix,
		},
		{
			"invalid matrix",
			[][]string{
				{"1", "b", "3"},
				{"4", "ijkl", "6"},
				{"7", "gh"},
			},
			0,
			utils.ErrInvalidMatrix,
		},
	}

	for _, tt := range tests {
		got, _ := NewMatrix(tt.matrix).Sum()
		assert.Equal(t, tt.want, got)
	}
}

func TestMultiply(t *testing.T) {

	tests := []struct {
		matrix [][]string
		want   int
	}{
		{
			matrix: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			want: 362880,
		},
	}

	for _, tt := range tests {
		got, _ := NewMatrix(tt.matrix).Multiply()
		assert.Equal(t, tt.want, got)
	}
}
