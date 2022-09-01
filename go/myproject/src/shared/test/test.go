package test

import (
	"{{ .ProjectName }}/src/infra"
	"{{ .ProjectName }}/src/repositories"
	"{{ .ProjectName }}/src/services"
)

type ContextString string

const DB ContextString = "db"

func GetServiceContainer() services.ServiceContainer {
	infraContainer := infra.New()
	repositoryContainer := repositories.New(infraContainer)
	return services.New(repositoryContainer)
}
