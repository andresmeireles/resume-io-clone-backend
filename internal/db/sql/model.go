package sql

type Model interface {
	TableName() string
	InsertFields() map[string]interface{}
	GetId() *uint
	Content() map[string]interface{}
}
