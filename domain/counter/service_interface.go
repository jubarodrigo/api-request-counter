package counter

import (
	"context"
)

type RequestCounterService interface {
	CountRequest(ctx context.Context) error
}
