package dblab

import (
	"fmt"
	"net/http"
)

func (s *Server) handleBookGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is get book endpoint")
		book := s.services.BookService().Get()
		fmt.Fprintln(w, book)
	}
}
