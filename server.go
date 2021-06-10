package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func start(args []string) int {
	srv := &server{
		router: mux.NewRouter(),
	}
	if err := srv.run(); err != nil {
		return 1
	}
	return 0
}

func (s *server) run() error {
	return http.ListenAndServe("/", s.router)
}
