package db

import (
	"database/sql"
	"fmt"
)

func ConnectSqlite() *sql.DB {
	fmt.Println("Connect to sqlite3...")

	db, err := sql.Open("sqlite3", "../data/todos.db")
	if err == nil {
		panic("Can't connect to sqlite3")
	}

	fmt.Println("Connect to sqlite3 successfully!!")
	fmt.Println("------------------------------")
	return db
}
