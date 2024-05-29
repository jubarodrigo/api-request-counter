package services

import (
	"context"
	"counter/config"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterService_CountRequest(t *testing.T) {
	tempFile, _ := ioutil.TempFile(os.TempDir(), "prefix")
	defer os.Remove(tempFile.Name())

	fileConfig := config.FileConfig{FilePath: tempFile.Name()}
	service := NewCounterService(fileConfig)

	count, err := service.CountRequest(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, int32(1), count)

	content, _ := ioutil.ReadFile(tempFile.Name())
	assert.Equal(t, "1", string(content))
}
