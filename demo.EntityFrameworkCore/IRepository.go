package repository

//IRepository IRepository
type IRepository[T any] interface {
	FirstOrDefault() T
}
