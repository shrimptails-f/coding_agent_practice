package di

import (
	"business/internal/library/mysql"

	"go.uber.org/dig"
)

// BuildContainer wires every dependency required by the sample stack.
func BuildContainer(conn *mysql.MySQL) *dig.Container {
	container := dig.New()

	ProvideSampleDependencies(container, conn)
	ProvidePresentationDependencies(container)

	return container
}
