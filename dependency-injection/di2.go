package dependencyInjection

import (
	"log"
	"os"
)

//MongoDataStore access to mongodb database
type MongoDataStore struct{}

//DataStore interface for datastore
type DataStore interface{}

//PersonResource - person resource
type PersonResource struct {
	Store  *DataStore
	Logger *log.Logger
}

//NewPersonResource -for new resource
func NewPersonResource(store *DataStore, logger *log.Logger) *PersonResource {
	return &PersonResource{Store: store, Logger: logger}
}

// func mainMethod() {
// 	e := InitializePersonResource()
// 	e.createItem("")
// }

func (pr *PersonResource) createItem(string string) {
}

//NewLogger new logger
func NewLogger() *log.Logger {
	return log.New(os.Stderr, "exp", log.LstdFlags|log.Lshortfile)
}

//NewDatastore Create for new data store
func NewDatastore() *DataStore {
	var datastore DataStore = &MongoDataStore{}
	return &datastore
}
