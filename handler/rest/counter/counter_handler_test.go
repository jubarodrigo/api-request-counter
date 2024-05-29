package counter

import (
	"counter/config"
	"counter/domain/counter/services"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandle_GetCounterRequests(t *testing.T) {
	// Preparação
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	tempFile, _ := ioutil.TempFile(os.TempDir(), "prefix")
	defer os.Remove(tempFile.Name())

	fileConfig := config.FileConfig{FilePath: tempFile.Name()}
	service := services.NewCounterService(fileConfig)

	handler := NewCounterRequestHandle(service)

	// Execução
	if assert.NoError(t, handler.GetCounterRequests(c)) {
		// Verificação
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `"Requests Total: 1"`, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
