package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func GetDbConnetion() *sql.DB {
	dataSourceName := fmt.Sprintf("host=localhost port=5432 dbname=spear user=postgres password=spear_db_") //fmt.Println("host is",os.Getenv("DB_HOST"))
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to conect to db"))
		panic(err)
	}
	log.Println("connected to db ")

	//test connection
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		log.Fatal("cannot ping db")
		panic(err)
	}
	log.Println("pinged db")
	return db
}
