package repository

type Repository interface {
	save(t interface{}) interface{}
}
