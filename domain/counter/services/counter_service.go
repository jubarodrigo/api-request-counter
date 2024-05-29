package services

import (
	"context"
	"fmt"
	"os"

	geometryDomain "counter/domain/counter"
)

func NewCounterService() geometryDomain.RequestCounterService {

	return &CounterService{}
}

type CounterService struct{}

const dataFile = "./storage/files/data.txt"

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

	file, err := os.Open(dataFile)
	if err != nil {
		return 0, err
	}

	fmt.Fscanf(file, "%d", &requests)

	err = file.Close()
	if err != nil {
		return 0, err
	}

	return requests, nil
}

func (cr *CounterService) saveData(requestCount int32) error {
	fmt.Println("Request to save: ", requestCount)
	file, err := os.Create(dataFile)
	if err == nil {
		fmt.Fprintf(file, "%d", requestCount)
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
