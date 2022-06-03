package test

import (
	"{{ .ProjectName }}/src/repositories"
	"{{ .ProjectName }}/src/services"
	"{{ .ProjectName }}/src/infra"
)

func GetServiceContainer() services.ServiceContainer {
	infraContainer := infra.New()
	repositoryContainer := repositories.New(infraContainer)
	return services.New(repositoryContainer)
}
