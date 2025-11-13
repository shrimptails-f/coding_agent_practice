package di

import (
	"business/internal/app/presentation"
	"business/internal/sample/application"

	"go.uber.org/dig"
)

// ProvidePresentationDependencies wires controllers that depend on application-layer services.
func ProvidePresentationDependencies(container *dig.Container) {
	_ = container.Provide(func(useCase application.UseCase) *presentation.SampleController {
		return presentation.NewSampleController(useCase)
	})
}
