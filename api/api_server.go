package api

import (
	"database/sql"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
)

type Server struct {
	DB *sql.DB
}

func CreateServer(db *sql.DB) *Server {

	var apiConf []Conf
	apiConfFile, err := os.ReadFile("api.yml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(apiConfFile, &apiConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println(apiConf)

	server := Server{
		DB: db,
	}

	return &server
}

func (s *Server) Run() error {

	http.HandleFunc("/api/v1/list", s.listHandler)
	http.HandleFunc("/api/v1/summary", s.summaryHandler)
	http.HandleFunc("/api/v1/create", s.createHandler)
	http.HandleFunc("/api/v1/delete", s.deleteHandler)
	http.HandleFunc("/api/v1/update", s.updateHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return err
	}
	return nil
}
