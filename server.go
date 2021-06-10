package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	db     *sql.DB
	router *mux.Router
	email  EmailSender
}

type EmailSender struct {
}

func start(args ...string) int {
	srv := &server{
		router: mux.NewRouter(),
	}
	srv.routes()
	if err := srv.run(); err != nil {
		return 1
	}
	return 0
}

func (s *server) run() error {
	return http.ListenAndServe("/", s.router)
}
