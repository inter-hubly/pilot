package pgsql

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var (
	oncePgsql       sync.Once
	pgConn          *connection
	pgsqlDefaultUrl = "jdbc:postgresql://localhost:5432/postgres"
)

type PConnInterface interface {
	GetConnection() *sql.DB
}

func GetConnection() *connection {
	return pgConn
}
