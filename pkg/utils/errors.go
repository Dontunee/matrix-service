package utils

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrEmptyMatrix          = errors.New("empty matrix")
	ErrInvalidMatrix        = errors.New("invalid matrix")
	ErrNonRectangularMatrix = errors.New("non-rectangular matrix")
	ErrNonNumericValues     = errors.New("matrix contains non-numeric values")
)

// RespondWithError sends a JSON-formatted error response
func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), statusCode)
}
