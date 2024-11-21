package server

import (
	"fmt"
	"os"

	"github.com/inter-hubly/pilot/hlog"
	"gopkg.in/yaml.v3"
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
	Env         Environment `yaml:"env"`
	Port        int         `yaml:"port"`
	HostName    string      `yaml:"hostname"`
	Secrettoken string      `yaml:"secrettoken"`
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
		startEnv(Environment(myEnv))
	}
}

func startEnv(envName Environment) {
	f, err := os.Open(fmt.Sprintf("config.%s.yaml", envName))
	if err != nil {
		hlog.Warn("StartEnv", "Failed to open config file", "err", err)
		// want to stop app, because don't have any environment
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&env); err != nil {
		hlog.Warn("StartEnv", "Failed to parse config file", "err", err)
		panic(err)
	}
	hlog.Info("StartEnv", fmt.Sprintf("Loading environment variables in %s environment in port %d", env.Config.Env, env.Config.Port))
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
