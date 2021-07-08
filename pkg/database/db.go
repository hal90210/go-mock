package database

type Document struct {
	ID   string
	Json string
}

// API definition
type DB interface {
	Upsert(*Document) error
	Get(id string) (*Document, error)
}
