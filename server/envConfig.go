package server

type dbConfig struct {
	Host       string `yaml:"host"`
	Database   string `yaml:"database"`
	EntryPoint string `yaml:"entryPoint"`
}

type ampqConfig struct {
	Host string `yaml:"host"`
}
