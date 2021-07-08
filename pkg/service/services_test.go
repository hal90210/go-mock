package service

import (
	"errors"
	"testing"

	"github.com/scottski/di-test/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ReceiveMessage_Success(t *testing.T) {

	dbMock := &database.DBMock{}

	// Return this on any argument passed in
	dbMock.On("Upsert", mock.Anything).Return(nil)

	msgService := NewMessageService(dbMock)

	status, err := msgService.RecieveMessage("")
	assert.Nil(t, err)
	assert.Equal(t, status, 200)
}

func Test_SendMessage_Error(t *testing.T) {

	dbMock := &database.DBMock{}

	// Return an error on any argument
	dbMock.On("Upsert", mock.Anything).Return(errors.New("Bad json"))

	msgService := NewMessageService(dbMock)

	status, err := msgService.RecieveMessage("")
	assert.NotNil(t, err)
	assert.Equal(t, status, 500)
	assert.Equal(t, err.Error(), "Bad json")
}
