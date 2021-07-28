package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/scottski/di-test/pkg/database"
)

// API
type MessageService interface {
	RecieveMessage(msg string) (httpCode int, err error)
}

// Do not export, force clients to use New to get instance
type messageReciver struct {
	db database.DBLayer
}

// Constructor
func NewMessageService(dbInstance database.DBLayer) messageReciver {
	return messageReciver{db: dbInstance}
}

// Handle http message sent from client - save to database
// returns http status code (int) and error
func (ms *messageReciver) RecieveMessage(msg string) (int, error) {

	fmt.Println("Recived Message....")
	fmt.Println(msg)

	doc := database.Document{ID: uuid.New().String(), Json: msg}

	// apply validation/business logic

	err := ms.db.Upsert(&doc)

	if err != nil {
		// Log error here or let caller log
		return 500, err
	}

	fmt.Println(">> Done")
	return 200, nil
}
