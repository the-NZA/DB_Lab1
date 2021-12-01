package dblab

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Config represents application config structure
type Config struct {
	AppDomain string `json:"app_domain"`
	AppPort   string `json:"app_port"`
	DBType    string `json:"db_type"`
	DBURL     string `json:"db_url"`
	LogDebug  bool   `json:"log_debug"`
}

// GetBindAddress concatenates application domain with port
func (c Config) GetBindAddress() string {
	return fmt.Sprintf("%s:%s", c.AppDomain, c.AppPort)
}

// ReadFromFile open file by passed path and trying to parse it
func (c *Config) ReadFromFile(p string) error {
	// Trying to open file
	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()

	// Trying to read bytes from opened file
	bytes, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	// Trying to parse config into 'c' structure
	err = json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}

	return nil
}

// NewConfig just return empty config
// with enabled debug level and :8080 port
func NewConfig() *Config {
	return &Config{
		AppDomain: "",
		AppPort:   "8080",
		DBType:    "",
		DBURL:     "",
		LogDebug:  true,
	}
}
