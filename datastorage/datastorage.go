// Package datastorage is resposible for storing build results in a given database
package datastorage

import (
	"fmt"

	"github.com/q231950/alethea/database"

	"github.com/apex/log"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/q231950/alethea/model"
)

type DataStorage interface {
	CreateIncidentsTable()
	StoreIncident(incident model.Incident) error
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

// CreateIncidentsTable creates an incidents table if necessary
func (ds AletheaDataStorage) CreateIncidentsTable() {
	log.Infof("%s", ds.database)
	incident := model.Incident{}
	log.Infof("%s", incident)
	ds.database.CreateTable(&incident, &orm.CreateTableOptions{
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
	fmt.Println(info)

	log.Infof("CreateIncidentsTable")
}

// StoreIncident stores the given build result in the database
func (ds AletheaDataStorage) StoreIncident(incident model.Incident) error {
	log.Infof("log incident")
	err := ds.database.Insert(&incident)
	if err != nil {
		panic(err)
	}
	fmt.Println(incident)
	return err
}
