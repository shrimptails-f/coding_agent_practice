package di

import (
	"business/internal/library/mysql"
	"business/internal/sample/application"
	"business/internal/sample/infrastructure"

	"go.uber.org/dig"
)

// ProvideSampleDependencies registers the domain/application/infrastructure bindings.
func ProvideSampleDependencies(container *dig.Container, conn *mysql.MySQL) {
	_ = container.Provide(func() *infrastructure.SampleRepository {
		return infrastructure.NewRepository(conn.DB)
	})

	_ = container.Provide(func(
		repo *infrastructure.SampleRepository,
	) application.UseCase {
		return application.NewUseCase(repo)
	})
}
