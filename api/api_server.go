package api

import (
	"database/sql"
	"net/http"
)

type Server struct {
	DB *sql.DB
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
