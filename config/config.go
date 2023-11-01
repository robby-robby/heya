package config

import (
	"os"
	"strings"
)

type config struct {
	Dsn          string
	ENV          env
	LogLevelStr  string
	OpenAIApiKey string
}
type env string

func (e env) IsDev() bool {
	return strings.ToUpper(string(e)) == "DEV"
}
func (e env) IsProd() bool {
	return strings.ToUpper(string(e)) == "PROD"
}
func NewConfig() *config {
	c := &config{
		Dsn:          getEnv("DSN", "file:heya.db"),
		ENV:          env(strings.ToUpper(getEnv("ENV", "PROD"))),
		LogLevelStr:  strings.ToUpper(getEnv("LOG_LEVEL", "DEBUG")),
		OpenAIApiKey: mustGetEnv("OPENAI_API_KEY"),
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
var Dsn = Config.Dsn
var ENV = Config.ENV
var LogLevelStr = Config.LogLevelStr
var OpenAIApiKey = Config.OpenAIApiKey
