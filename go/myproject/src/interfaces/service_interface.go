package interfaces

type IService interface {
	Inject(gateways IGateway) IService

	GetTodoService() ITodoService
	SetTodoService(s ITodoService)

	GetAuthService() IAuthService
	SetAuthService(s IAuthService)
}
