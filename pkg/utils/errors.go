package utils

import "errors"

var (
	ErrEmptyMatrix          = errors.New("empty matrix")
	ErrInvalidMatrix        = errors.New("invalid matrix")
	ErrNonRectangularMatrix = errors.New("non-rectangular matrix")
	ErrNonNumericValues     = errors.New("matrix contains non-numeric values")
)
