package cmd

import (
	cunterService "counter/domain/counter"
	"counter/domain/counter/services"
)

type Counter struct {
	GeometryService cunterService.RequestCounterService
}

func NewCounter() *Counter { return &Counter{} }

type counterBuilder interface {
	setConfig()
	setServices()
	setHttpServer()
}

type BuildDirector struct {
	builder counterBuilder
}

func NewCounterBuilder(b counterBuilder) *BuildDirector {
	return &BuildDirector{
		builder: b,
	}
}

func (sm *BuildDirector) BuildCounter() {
	sm.builder.setConfig()
	sm.builder.setServices()
	sm.builder.setHttpServer()
}

func (sm *Counter) setConfig() {
	// _ := config.GetEnvOrDefault("ENV_FILE_PATH", ".env")
}

func (sm *Counter) setServices() {
	sm.GeometryService = services.NewCounterService()
}

func (sm *Counter) setHttpServer() {
	StartHttpServer(sm.GeometryService)
}
