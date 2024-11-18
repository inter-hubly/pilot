package server

type dbConfig struct {
	Host       string `yaml:"host"`
	Database   string `yaml:"database"`
	EntryPoint string `yaml:"entryPoint"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}

type ampqConfig struct {
	Host string `yaml:"host"`
}
