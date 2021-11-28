package main

import (
	"database/sql"
	"deblog-go/api"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

func setupDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/deblog?multiStatements=true")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("creating db")
	sqlFile, err := os.ReadFile("app.sql")
	if err != nil {
		return nil, err
	}
	exec, err := db.Exec(string(sqlFile))
	if err != nil {
		return nil, err
	}
	ra, err := exec.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println("RA:", ra)
	return db, nil
}

func main() {
	fmt.Println("deblog-go")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err := setupDatabase()
	if err != nil {
		log.Println(err)
	}

	server := api.CreateServer(db)
	if err := server.Run(); err != nil {
		log.Println(err)
	}
}
