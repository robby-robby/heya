package config

import (
	"os"
	"strings"
)

type config struct {
	Dsn         string
	ENV         ENV
	LogLevelStr string
}
type ENV string

func (e ENV) IsDev() bool {
	return strings.ToUpper(string(e)) == "DEV"
}
func (e ENV) IsProd() bool {
	return strings.ToUpper(string(e)) == "PROD"
}
func NewConfig() *config {
	c := &config{
		Dsn:         getEnv("DSN", "file:heya.db"),
		ENV:         ENV(strings.ToUpper(getEnv("ENV", "PROD"))),
		LogLevelStr: strings.ToUpper(getEnv("LOG_LEVEL", "DEBUG")),
	}
	return c
}

func getEnv(key string, defV string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defV
}

// func getPath(rawurl string) string {
// 	u, err := url.Parse(rawurl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return u.Path
// }

func mustGetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic("env var " + key + " not found")
}

var Config = NewConfig()
