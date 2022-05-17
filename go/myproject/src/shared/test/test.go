package test

import (
	"myproject/src/repositories"
	"myproject/src/services"
	"myproject/src/infra"
)

func GetServiceContainer() services.ServiceContainer {
	infraContainer := infra.New()
	repositoryContainer := repositories.New(infraContainer)
	return services.New(repositoryContainer)
}
