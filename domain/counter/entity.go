package counter

import "sync"

type RequestCount struct {
	requestCount  int
	mu            sync.RWMutex
	lastTimestamp int64
}
