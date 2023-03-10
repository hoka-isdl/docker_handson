package repository

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Opendb() {

	db_name, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/dockerlec_database?")

	if err != nil {
		panic(err.Error())
	}

	db = db_name
}

func Register(username string,password string){
	fmt.Print(1)
	Opendb()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO Students(username,password) VALUES(?, ?)")

	if err != nil {
		log.Fatal(err.Error())
	}
	
	insert.Exec(username,password)
}