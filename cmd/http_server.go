package cmd

import (
	"counter/bin"
	counterService "counter/domain/counter"
	"counter/handler/rest/counter"
	"github.com/labstack/echo/v4"
)

func StartHttpServer(geometryService counterService.RequestCounterService) {
	e := echo.New()
	e.HideBanner = true

	go bin.RunTicker()

	geometryAnalysisHandle := counter.NewCounterRequestHandle(geometryService)

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/counter", func(c echo.Context) error { return geometryAnalysisHandle.GetCounterRequests(c) })

	e.Logger.Fatal(e.Start(":8090"))
}
