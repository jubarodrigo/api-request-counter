package counter

import (
	"context"
	"fmt"
	"net/http"

	counterService "counter/domain/counter"
	"github.com/labstack/echo/v4"
)

type RequestHandle struct {
	counterService counterService.RequestCounterService
}

func NewCounterRequestHandle(geometryService counterService.RequestCounterService) *RequestHandle {
	return &RequestHandle{
		counterService: geometryService,
	}
}

func (sfh *RequestHandle) GetCounterRequests(c echo.Context) error {
	ctx := context.Background()

	counter, err := sfh.counterService.CountRequest(ctx)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}

	requests := fmt.Sprintf("Requests Total: %d", counter)

	return c.JSON(http.StatusOK, requests)
}
