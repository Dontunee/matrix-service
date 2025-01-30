package handlers

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {

	csvContent := "1,2,3\n4,5,6\n7,8,9"
	fileContent := new(bytes.Buffer)
	writer := multipart.NewWriter(fileContent)
	part, _ := writer.CreateFormFile("file", "test.csv")
	part.Write([]byte(csvContent))
	writer.Close()

	req, _ := http.NewRequest("GET", "/echo", fileContent)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(EchoHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("failed: got %v, want %v", rr.Code, http.StatusOK)
	}

}

func TestInvertHandler(t *testing.T) {

	csvContent := "1,2,3\n4,5,6\n7,8,9"
	fileContent := new(bytes.Buffer)
	writer := multipart.NewWriter(fileContent)
	part, _ := writer.CreateFormFile("file", "test.csv")
	part.Write([]byte(csvContent))
	writer.Close()

	req, _ := http.NewRequest("GET", "/invert", fileContent)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(InvertHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("failed: got %v, want %v", rr.Code, http.StatusOK)
	}

}

func TestFlattenHandler(t *testing.T) {

	csvContent := "1,2,3\n4,5,6\n7,8,9"
	fileContent := new(bytes.Buffer)
	writer := multipart.NewWriter(fileContent)
	part, _ := writer.CreateFormFile("file", "test.csv")
	part.Write([]byte(csvContent))
	writer.Close()

	req, _ := http.NewRequest("GET", "/flatten", fileContent)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FlattenHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("failed: got %v, want %v", rr.Code, http.StatusOK)
	}

}

func TestSumHandler(t *testing.T) {

	csvContent := "1,2,3\n4,5,6\n7,8,9"
	fileContent := new(bytes.Buffer)
	writer := multipart.NewWriter(fileContent)
	part, _ := writer.CreateFormFile("file", "test.csv")
	part.Write([]byte(csvContent))
	writer.Close()

	req, _ := http.NewRequest("GET", "/sum", fileContent)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SumHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("failed: got %v, want %v", rr.Code, http.StatusOK)
	}

}

func TestMultiplyHandler(t *testing.T) {

	csvContent := "1,2,3\n4,5,6\n7,8,9"
	fileContent := new(bytes.Buffer)
	writer := multipart.NewWriter(fileContent)
	part, _ := writer.CreateFormFile("file", "test.csv")
	part.Write([]byte(csvContent))
	writer.Close()

	req, _ := http.NewRequest("GET", "/multiply", fileContent)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MultiplyHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("failed: got %v, want %v", rr.Code, http.StatusOK)
	}

}
