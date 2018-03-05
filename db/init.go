package db

import (
	"fmt"

	"../entities"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var models = []interface{}{
	&entities.Commodity{},
	&entities.Location{},
	&entities.Parity{},
	&entities.Price{},
	&entities.User{}}

func createSchema(db *pg.DB) error {

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			return err
		}
		fmt.Printf("Created Table: ")
		fmt.Printf("%+v\n", model)
	}
	return nil
}

func destroySchema(db *pg.DB) error {

	for _, model := range models {
		err := db.DropTable(model, &orm.DropTableOptions{IfExists: true})
		if err != nil {
			return err
		}
		fmt.Printf("Deleted Table: ")
		fmt.Printf("%+v\n", model)
	}
	return nil
}

func dbConnect() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "commonwand",
		Password: "commonwand",
		Database: "commonwand"})
	return db
}

// InsertModel interts a new price object
func InsertModel(model interface{}) {
	db := dbConnect()
	err := db.Insert(model)
	if err != nil {
		panic(err)
	}
}

// Init initialses the DB
func Init() *pg.DB {
	fmt.Println("Initialising DB Connection")
	// connStr := "user=commonwand password=commonwand dbname=commonwand  sslmode=disable"
	db := dbConnect()

	fmt.Printf("%+v\n", db)

	err := createSchema(db)
	if err != nil {
		panic(err)
	}
	return db
}
