package database

type Document struct {
	ID   string
	Json string
}

// API definition
type DBLayer interface {
	Upsert(*Document) error
	Get(id string) (*Document, error)
}
