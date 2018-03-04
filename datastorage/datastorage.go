// Package datastorage is resposible for storing build results in a given database
package datastorage

import (
	"reflect"

	"github.com/q231950/alethea/database"

	"github.com/apex/log"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type DataStorage interface {
	CreateTable(model interface{})
	StoreCIBuild(model interface{}) error
}

// DataStorage has a database where it stores build results in
type AletheaDataStorage struct {
	database pg.DB
}

// New creates a new DataStorage. It needs a database to store build results in
func New() DataStorage {
	log.Infof("New DataStorage")
	return AletheaDataStorage{database: database.PostgresqlDatabase()}
}

// CreateTable creates a table for the given model if none exists yet
func (ds AletheaDataStorage) CreateTable(model interface{}) {
	log.Debugf("Creating table for database (%s) if necessary", ds.database)
	ds.database.CreateTable(model, &orm.CreateTableOptions{
		Temp: false,
	})

	var info []struct {
		ColumnName string
		DataType   string
	}
	_, err := ds.database.Query(&info, `
			SELECT column_name, data_type
			FROM information_schema.columns
			WHERE table_name = 'models'
		`)
	if err != nil {
		panic(err)
	}

	log.Infof("Created table for model %s", reflect.TypeOf(model))
}

// StoreCIBuild stores the given CI build in the database
func (ds AletheaDataStorage) StoreCIBuild(build interface{}) error {
	log.Infof("Store CI build")
	err := ds.database.Insert(build)
	if err == nil {
		log.Infof("Stored build %s", build)
	}
	return err
}
