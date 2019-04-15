package main

import (
	"database/sql"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type DBAddr struct {
	Dsn string
}

func examDB_entry() {
	connectDB()
	// examDrivers()
}
func examDrivers() {
	drivers := sql.Drivers()
	fmt.Println(drivers)
}
func connectDB() {
	dbw := DBAddr{
		Dsn: "test:test123@tcp(127.0.0.1:3306)/goDB",
	}
	db, err := sql.Open("mysql", dbw.Dsn)
	if err := db.Ping(); err != nil {
		panic(err)
		return
	} else {
		println("DB ping OK :p")
	}
	if err != nil {
		panic(err)
		return
	} else {
		println("DB connected")
	}
	defer db.Close()
}
func disconnectDB() {

}
