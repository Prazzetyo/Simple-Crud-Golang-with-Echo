package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var err error

func Init()  {
	// conf := config.GetConfig()

	// connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@" + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	
	// Local DB
	connectionString := "root:@/pegawai"
	
	// Online DB
	// connectionString := "bertandur:1sampaibertandur@tcp(db-2.badabes.my.id)/bertandur"
	
	db, err = sql.Open("mysql", connectionString)
	
	if err != nil {
		panic("connectionString Error")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func CreateCon() *sql.DB {
	return db
}