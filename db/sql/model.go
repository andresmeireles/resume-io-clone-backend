package sql

type Model interface {
	create(data ...interface{}) (Model, error)
	getAll() ([]Model, error)
	getById(id int) (Model, error)
	update(id int, data ...interface{}) (Model, error)
	delete(id int) error
}
