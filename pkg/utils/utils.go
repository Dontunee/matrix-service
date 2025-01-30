package utils

import (
	"encoding/csv"
	"net/http"
)

func ReadCSVFromRequest(r *http.Request, fieldName string) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
