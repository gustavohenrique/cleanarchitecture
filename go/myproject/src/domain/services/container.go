package services

import (
	"{{ .ProjectName }}/src/domain/services/auth"
	"{{ .ProjectName }}/src/domain/services/todo"
	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/interfaces"
)

type ServiceContainer struct {
	config      *conf.Config
	todoService interfaces.ITodoService
	authService interfaces.IAuthService
}

func With(config *conf.Config) interfaces.IService {
	return &ServiceContainer{
		config: config,
	}
}

func (c *ServiceContainer) Inject(gateways interfaces.IGateway) interfaces.IService {
	c.todoService = todo.New(gateways)
	c.authService = auth.With(c.config).New(gateways)
	return c
}

func (c *ServiceContainer) GetTodoService() interfaces.ITodoService {
	return c.todoService
}

func (c *ServiceContainer) SetTodoService(service interfaces.ITodoService) {
	c.todoService = service
}

func (c *ServiceContainer) GetAuthService() interfaces.IAuthService {
	return c.authService
}

func (c *ServiceContainer) SetAuthService(service interfaces.IAuthService) {
	c.authService = service
}
