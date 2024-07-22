package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for POST /api/v1/uploadFile
func TestPostUser(t *testing.T) {
	router := MakeRouter()

	w := httptest.NewRecorder()

	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	fileWriter, err := multipartWriter.CreateFormFile("file", "file.txt")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}

	file, err := os.Open("file.txt")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}

	multipartWriter.Close()

	username := "user1"
	password := "love"

	auth := username + ":" + password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	req, _ := http.NewRequest("POST", "/api/v1/uploadFile", &requestBody)
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
