package dblab

import "github.com/go-pkgz/lgr"

// configureLogger return new logger with debug or common settings
func configureLogger(isDebug bool) *lgr.Logger {
	if isDebug {
		return lgr.New(lgr.Msec, lgr.Debug, lgr.CallerFile, lgr.CallerFunc, lgr.LevelBraces)
	}

	return lgr.New(lgr.Msec, lgr.LevelBraces)
}
