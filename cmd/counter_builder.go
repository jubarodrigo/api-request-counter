package cmd

import (
	"counter/config"
	cunterService "counter/domain/counter"
	"counter/domain/counter/services"
)

type Counter struct {
	GeometryService cunterService.RequestCounterService
	FileConfig      config.FileConfig
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
	env := config.GetEnvOrDefault("ENV_FILE_PATH", ".env")
	config.MustParseConfig(env, &sm.FileConfig)
}

func (sm *Counter) setServices() {
	sm.GeometryService = services.NewCounterService(sm.FileConfig)
}

func (sm *Counter) setHttpServer() {
	StartHttpServer(sm.GeometryService, sm.FileConfig)
}
