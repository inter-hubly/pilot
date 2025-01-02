package server

import (
	"fmt"
	"os"

	"github.com/inter-hubly/pilot/hlog"
	"gopkg.in/yaml.v3"
)

type MockEnvironment struct {
	PgsqlHost       string
	PgsqlDatabase   string
	PgsqlEntryPoint string

	RedisHost       string
	RedisDatabase   string
	RedisEntryPoint string

	ElasticSearchHost string
}

func NewMockEnvironment(mock MockEnvironment) {
	env.Config.HostName = "https://test.io"
	env.Pgsql = dbConfig{
		Host:       mock.PgsqlHost,
		Database:   mock.PgsqlDatabase,
		EntryPoint: mock.PgsqlEntryPoint,
	}
	env.Redis = dbConfig{
		Host:       mock.RedisHost,
		Database:   mock.RedisDatabase,
		EntryPoint: mock.RedisEntryPoint,
	}
}

func MockStartEnv(baseRoot string) {
	f, err := os.Open(fmt.Sprintf("%sconfig.test.yaml", baseRoot))
	if err != nil {
		currentDir, _ := os.Getwd()
		hlog.Warn("MockStartEnv", fmt.Sprintf("Failed to open config file: %s with error: %s", currentDir, err))
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&env); err != nil {
		hlog.Warn("MockStartEnv", "Failed to parse config file", "err", err)
		panic(err)
	}
	hlog.Info("MockStartEnv", fmt.Sprintf("Loading environment variables in %s environment in port %d", env.Config.Env, env.Config.Port))
}
