package main

import (
	"database/sql"
	"deblog-go/api"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
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

	//query := `SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = 'deblog'`
	//row := db.QueryRow(query)
	//var dbExists bool
	//
	//row.Scan(&dbExists)
	//fmt.Println("db exists: ", dbExists)

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

type ApiConf struct {
	Module       string   `yaml:"module"`
	BeforeCreate []string `yaml:"beforeCreate"`
	AfterCreate  []string `yaml:"afterCreate"`
}

func main() {
	fmt.Println("deblog-go")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err := setupDatabase()
	if err != nil {
		log.Println(err)
	}

	_ = db

	var apiConf []ApiConf

	apiConfFile, err := os.ReadFile("api.yml")
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal(apiConfFile, &apiConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Println(apiConf)

	//
	//ctx := map[string]interface{} {
	//	"db": db,
	//	"post": &post,
	//}
	//chn := chain.GetSavePostChain()
	//if err := chn.Exec(ctx); err != nil {
	//	log.Println(err)
	//}
	//
	//chn1 := chain.GetPostGetAllChain()
	//if err := chn1.Exec(ctx); err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(ctx["posts"])

	server := api.Server{
		DB: db,
	}

	if err := server.Run(); err != nil {
		log.Println(err)
	}
}
