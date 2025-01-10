package server

import (
	"fmt"
	"os"

	"github.com/inter-hubly/pilot/hlog"
)

// Environment is a env value
type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

type Server struct {
	config
}

type config struct {
	Env                Environment `yaml:"env"`
	Port               string      `yaml:"port"`
	HostName           string      `yaml:"hostname"`
	WhatsAppSecretWord string      `yaml:"whatsAppSecretWord"`
	HashEncrypt        string      `yaml:"hashEncrypt"`
}

type environment struct {
	Config        config     `yaml:"config"`
	Mongo         dbConfig   `yaml:"mongo"`
	Pgsql         dbConfig   `yaml:"pgsql"`
	Redis         dbConfig   `yaml:"redis"`
	ElasticSearch dbConfig   `yaml:"elasticSearch"`
	Ampq          ampqConfig `yaml:"ampq"`
}

var env environment

func FillConfigEnvironment() {
	if myEnv := os.Getenv("ENVIRONMENT"); myEnv != "" {
		startEnv()
	}
}

func startEnv() {
	env = environment{
		Config: config{
			Env:                Environment(os.Getenv("ENVIRONMENT")),
			Port:               os.Getenv("ENVIRONMENT_PORT"),
			HostName:           os.Getenv("ENVIRONMENT_HOSTNAME"),
			WhatsAppSecretWord: os.Getenv("ENVIRONMENT_WHATSAPP_SECRET"),
			HashEncrypt:        os.Getenv("ENVIRONMENT_HASH_ENCRYPT"),
		},
		Mongo: dbConfig{
			Host:     os.Getenv("ENVIRONMENT_MONGO_HOST"),
			Database: os.Getenv("ENVIRONMENT_MONGO_DATABASE"),
			Username: os.Getenv("ENVIRONMENT_MONGO_USERNAME"),
			Password: os.Getenv("ENVIRONMENT_MONGO_PASSWORD"),
		},
		Pgsql: dbConfig{
			Host:       os.Getenv("ENVIRONMENT_PGSQL_HOST"),
			Database:   os.Getenv("ENVIRONMENT_PGSQL_DATABASE"),
			Username:   os.Getenv("ENVIRONMENT_PGSQL_USERNAME"),
			Password:   os.Getenv("ENVIRONMENT_PGSQL_PASSWORD"),
			EntryPoint: os.Getenv("ENVIRONMENT_PGSQL_ENTRY_POINT"),
		},
		Redis: dbConfig{
			Host:     os.Getenv("ENVIRONMENT_REDIS_HOST"),
			Database: os.Getenv("ENVIRONMENT_REDIS_DATABASE"),
			Username: os.Getenv("ENVIRONMENT_REDIS_USERNAME"),
			Password: os.Getenv("ENVIRONMENT_REDIS_PASSWORD"),
		},
		ElasticSearch: dbConfig{
			Host:     os.Getenv("ENVIRONMENT_ELASTICSEARCH_HOST"),
			Database: os.Getenv("ENVIRONMENT_ELASTICSEARCH_DATABASE"),
			Username: os.Getenv("ENVIRONMENT_ELASTICSEARCH_USERNAME"),
			Password: os.Getenv("ENVIRONMENT_ELASTICSEARCH_PASSWORD"),
		},
		Ampq: ampqConfig{
			Host: os.Getenv("ENVIRONMENT_AMPQ_HOST"),
		},
	}
	hlog.Info("StartEnv", fmt.Sprintf("Loading environment variables in %s", env))
}

func GetEnvironment() config {
	return env.Config
}
func GetMongoConfig() dbConfig {
	return env.Mongo
}
func GetPgsqlConfig() dbConfig {
	return env.Pgsql
}
func GetRedisConfig() dbConfig {
	return env.Redis
}
func GetElasticSearch() dbConfig {
	return env.ElasticSearch
}
func GetAmpqConfig() ampqConfig {
	return env.Ampq
}
