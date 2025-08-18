package tests

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/dimitar728/virtual-showroom/backend/internal/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
	app := fiber.New()
	// mount routes directly for test
	app.Post("/api/showrooms", contro	llers.CreateShowroom)
	return app
}

func TestUploadValidModel(t *testing.T) {
	app := setupApp()

	// prepare test file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, _ := os.Open("testdata/valid.glb")
	defer file.Close()
	part, _ := writer.CreateFormFile("model", filepath.Base(file.Name()))
	_, _ = bytes.NewReader([]byte("dummy")).WriteTo(part)

	_ = writer.WriteField("name", "Test Showroom")
	writer.Close()

	req := httptest.NewRequest("POST", "/api/showrooms", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)
}

func TestUploadInvalidModel(t *testing.T) {
	app := setupApp()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, _ := os.Open("testdata/invalid.txt")
	defer file.Close()
	part, _ := writer.CreateFormFile("model", filepath.Base(file.Name()))
	_, _ = bytes.NewReader([]byte("dummy")).WriteTo(part)

	_ = writer.WriteField("name", "Invalid Showroom")
	writer.Close()

	req := httptest.NewRequest("POST", "/api/showrooms", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, _ := app.Test(req)

	assert.Equal(t, 400, resp.StatusCode)
}
