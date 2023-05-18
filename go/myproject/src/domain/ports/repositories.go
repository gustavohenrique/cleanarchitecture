package ports

type Repositories interface {
	TodoRepository() TodoRepository
}
