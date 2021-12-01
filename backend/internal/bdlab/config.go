package bdlab

import "fmt"

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

func NewConfig() *Config {
	return &Config{
		AppDomain: "",
		AppPort:   ":8080",
		DBType:    "",
		DBURL:     "",
		LogDebug:  true,
	}
}
