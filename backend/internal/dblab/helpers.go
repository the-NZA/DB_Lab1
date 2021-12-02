package dblab

import (
	"encoding/json"
	"net/http"

	"github.com/go-pkgz/lgr"
)

// configureLogger return new logger with debug or common settings
func configureLogger(isDebug bool) *lgr.Logger {
	if isDebug {
		return lgr.New(lgr.Msec, lgr.Debug, lgr.CallerFile, lgr.CallerFunc, lgr.LevelBraces)
	}

	return lgr.New(lgr.Msec, lgr.LevelBraces)
}

// respond method perform response with json encoding and optional data
func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// error method perform respoinse with error with wrapping it
func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
