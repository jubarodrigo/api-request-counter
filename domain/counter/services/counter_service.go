package services

import (
	"context"
	"counter/config"
	"fmt"
	"os"

	geometryDomain "counter/domain/counter"
)

func NewCounterService(fileConfig config.FileConfig) geometryDomain.RequestCounterService {
	return &CounterService{
		fileConfig: fileConfig,
	}
}

const errorFileOperation = "service: error with file operation: %w"

type CounterService struct {
	fileConfig config.FileConfig
}

func (cr *CounterService) CountRequest(ctx context.Context) (int32, error) {
	requestCount, err := cr.loadData()
	if err != nil {
		return 0, err
	}

	requestCount++

	err = cr.saveData(requestCount)
	if err != nil {
		return 0, err
	}

	return requestCount, nil
}

func (cr *CounterService) loadData() (int32, error) {
	var requests int32

	file, err := os.Open(cr.fileConfig.FilePath)
	if err != nil {
		return 0, fmt.Errorf(errorFileOperation, err)
	}

	fmt.Fscanf(file, "%d", &requests)

	err = file.Close()
	if err != nil {
		return 0, fmt.Errorf(errorFileOperation, err)
	}

	return requests, nil
}

func (cr *CounterService) saveData(requestCount int32) error {
	file, err := os.Create(cr.fileConfig.FilePath)
	if err == nil {
		fmt.Fprintf(file, "%d", requestCount)
	}

	err = file.Close()
	if err != nil {
		return fmt.Errorf(errorFileOperation, err)
	}

	return nil
}
