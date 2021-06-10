package main

import (
	"fmt"
	"net/http"
	"sync"
	"text/template"
)

func (s *server) routes() {
	s.router.HandleFunc("/api/", s.handleAPI())
	s.router.HandleFunc("/about", s.handleAbout())
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
}

func prepareThing() string {
	return ""
}

type User struct {
	IsAdmin bool
}

func currentUser(r *http.Request) User {
	return User{}
}

func (*server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !currentUser(r).IsAdmin {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	}
}

func (*server) handleGreeting(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, format, "World")
	}
}

func (*server) handleAPI() http.HandlerFunc {
	thing := prepareThing()
	return func(http.ResponseWriter, *http.Request) {
		fmt.Println(thing)
	}
}

func (*server) handleAbout() http.HandlerFunc {
	return func(http.ResponseWriter, *http.Request) {
	}
}

func (*server) handleIndex() http.HandlerFunc {
	type request struct {
		Name string
	}
	type response struct {
		Greeting string `json:"greeting"`
	}
	return func(http.ResponseWriter, *http.Request) {
	}
}

func (*server) handleAdminIndex() http.HandlerFunc {
	return func(http.ResponseWriter, *http.Request) {
	}
}

func (*server) handleTemplate() http.HandlerFunc {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	files := []string{}
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, tplerr = template.ParseFiles(files...)
		})
		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(tpl)
	}
}
