package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func GetDbConnetion() *sql.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"))
	//fmt.Println("host is",os.Getenv("DB_HOST"))
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

