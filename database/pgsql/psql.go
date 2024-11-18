package pgsql

import (
	"sync"

	_ "github.com/lib/pq"
)

var (
	oncePgsql       sync.Once
	pgConn          *connection
	pgsqlDefaultUrl = "jdbc:postgresql://localhost:5432/postgres"
)

func GetConnection() *connection {
	return pgConn
}
