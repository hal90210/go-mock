# go-mock

## Usage: *Mocks need to mock interfaces*


```// API definition
type DB interface {
	Upsert(*Document) error
	Get(id string) (*Document, error)
}
```

#### Create the mock

```
type DBMock struct {
	mock.Mock
}

func (m *DBMock) Upsert(doc *Document) error {
	// gets
	args := m.Called(doc)
	// returns
	return args.Error(0)
}

func (m *DBMock) Get(id string) (*Document, error) {
	// gets
	args := m.Called(id)
	// returns
	return args.Get(0).(*Document), args.Error(1)
}
```

#### In the test, create an instance of your mock

```	dbMock := &database.DBMock{}```

#### In test, specify func return values

```
// Return this on any argument passed in
dbMock.On("Upsert", mock.Anything).Return(nil)
```
