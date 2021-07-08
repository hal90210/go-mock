package externalDB1

import "github.com/scottski/di-test/pkg/database"

// Composer of the DB interface
type DynamoDB struct {
}

// Constructor
func NewDB() DynamoDB {
	return DynamoDB{}
}

func (db *DynamoDB) Upsert(*database.Document) error {
	return nil
}

func (db *DynamoDB) Get(id string) (*database.Document, error) {
	return &database.Document{}, nil
}
