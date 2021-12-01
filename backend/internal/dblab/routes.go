package dblab

import (
	"fmt"
	"net/http"
)

func (s *Server) handleIndexPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is index page")
	}
}
