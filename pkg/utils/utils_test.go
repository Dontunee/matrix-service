package utils

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSVFromRequest(t *testing.T) {

	csvContent := "1,2,3\n4,5,6\n7,8,9"
	fileContent := new(bytes.Buffer)
	writer := multipart.NewWriter(fileContent)
	part, _ := writer.CreateFormFile("file", "sample.csv")
	part.Write([]byte(csvContent))
	writer.Close()

	req, _ := http.NewRequest("GET", "/echo", fileContent)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	records, err := ReadCSVFromRequest(req, "file")
	assert.Nil(t, err)

	expectedRecords := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	assert.ObjectsAreEqual(records, expectedRecords)
}
