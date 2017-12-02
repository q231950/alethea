package datastorage

import (
	"fmt"

	"github.com/apex/log"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/q231950/alethea/model"
)

//
type DataStorage struct {
	database pg.DB
}

func NewDataStorage(db pg.DB) *DataStorage {
	log.Infof("NewDataStorage")
	ds := DataStorage{database: db}
	return &ds
}

// CreateIncidentsTable creates an incidents table if necessary
func (ds *DataStorage) CreateIncidentsTable() {
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

func (ds *DataStorage) LogIncident(incident model.Incident) error {
	log.Infof("log incident")
	err := ds.database.Insert(&incident)
	if err != nil {
		panic(err)
	}
	fmt.Println(incident)
	return err
}
