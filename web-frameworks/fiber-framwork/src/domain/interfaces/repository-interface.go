package interfaces

type Repository[T any] interface {
	Create(entity T) (T, error)
	FindAll() ([]T, error)
	FindByID(id string) (T, error)
	Update(id string, entity T) (T, error)
	Delete(id string) error
}
