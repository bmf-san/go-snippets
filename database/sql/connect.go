package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	dataSourceName := fmt.Sprintf("user:password@tcp(host:port/database")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connected DB:", err)
	}
	defer db.Close()
}
