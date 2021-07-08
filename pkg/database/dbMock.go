package database

import (
	"github.com/stretchr/testify/mock"
)

type DBMock struct {
	mock.Mock
}

func (m *DBMock) Upsert(doc *Document) error {
	// gets
	args := m.Called(doc)
	// returns
	return args.Error(0)
}

func (m *DBMock) Get(id string) (Document, error) {
	// gets
	args := m.Called(id)
	// returns
	return args.Get(0).(Document), args.Error(1)
}
