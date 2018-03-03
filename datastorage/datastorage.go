// Package datastorage is resposible for storing build results in a given database
package datastorage

import (
	"fmt"
	"reflect"

	"github.com/q231950/alethea/database"

	"github.com/apex/log"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type DataStorage interface {
	CreateTable(model interface{})
	StoreIncident(model interface{}) error
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

// StoreIncident stores the given build result in the database
func (ds AletheaDataStorage) StoreIncident(incident interface{}) error {
	log.Infof("log incident")
	err := ds.database.Insert(&incident)
	if err != nil {
		panic(err)
	}
	fmt.Println(incident)
	return err
}
