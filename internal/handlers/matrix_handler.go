package handlers

import (
	"fmt"
	"github.com/Dontunee/matrix-service/internal/models"
	"github.com/Dontunee/matrix-service/pkg/utils"
	"net/http"
)

// EchoHandler prints the matrix as a string in matrix format.
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	records, err := utils.ReadCSVFromRequest(r, "file")

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Fprint(w, models.NewMatrix(records).ToString())
}

// InvertHandler prints the matrix as a string in matrix format where the columns and rows are inverted
func InvertHandler(w http.ResponseWriter, r *http.Request) {
	records, err := utils.ReadCSVFromRequest(r, "file")

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Fprint(w, models.NewMatrix(records).Invert())
}

// FlattenHandler prints the numbers in the matrix on one line
func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	records, err := utils.ReadCSVFromRequest(r, "file")

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Fprint(w, models.NewMatrix(records).Flatten())
}

// SumHandler prints the sum of the integers in the matrix
func SumHandler(w http.ResponseWriter, r *http.Request) {
	records, err := utils.ReadCSVFromRequest(r, "file")

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	response, err := models.NewMatrix(records).Sum()
	fmt.Println(response)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Fprint(w, response)
}

// MultiplyHandler prints the product of the integers in the matrix
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	records, err := utils.ReadCSVFromRequest(r, "file")

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	response, err := models.NewMatrix(records).Multiply()

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Fprint(w, response)
}
