package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct{}

func (s *Server) stage1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "stage1")
}

func (s *Server) stage2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "stage2: ", mux.Vars(r)["serial"])
}
