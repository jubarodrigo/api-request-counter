package services

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	geometryDomain "counter/domain/counter"
)

func NewCounterService() geometryDomain.RequestCounterService {

	return &CounterService{}
}

type CounterService struct{}

const dataFile = "data.txt"

var (
	requestCount  int
	mu            sync.RWMutex
	lastTimestamp int64
)

func (cr *CounterService) CountRequest(ctx context.Context) error {
	now := time.Now().Unix()
	if now-lastTimestamp >= 60 {
		requestCount = 0
		lastTimestamp = now
	}

	requestCount++

	return nil
}

func loadData() {
	file, err := os.Open(dataFile)
	if err == nil {
		fmt.Fscanf(file, "%d %d", &requestCount, &lastTimestamp)
	}
	file.Close()
}

func saveData() {
	file, err := os.Create(dataFile)
	if err == nil {
		fmt.Fprintf(file, "%d %d", requestCount, lastTimestamp)
	}
	file.Close()
}
