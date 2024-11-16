package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/inter-hubly/pilot/hlog"
	"github.com/inter-hubly/pilot/server"
	"github.com/pkg/errors"
)

// Option was a function optional pattern
type Option func(*connection)

func WithUrl(url string) Option {
	return func(a *connection) {
		a.url = url
	}
}

type connection struct {
	url   string
	pgsql *sql.DB
}

func NewConnection(opts ...Option) {
	oncePgsql.Do(func() {

		pgConn.url = pgsqlDefaultUrl
		for _, opt := range opts {
			opt(pgConn)
		}

		pgConn.pgsql = pgConn.conn()
		if server.GetPgsqlConfig().EntryPoint != "" {
			pgConn.CreateScripts(server.GetPgsqlConfig().EntryPoint)
		}
	})
}

func (c *connection) conn() *sql.DB {
	ctx := context.TODO()
	host, port, dbname, user, password, err := extractDBDetails(c.url)
	if err != nil {
		hlog.Error("NewConnPgsql", fmt.Sprintf("Error opening database: %q", err))
	}
	sprintf := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", sprintf)
	if err != nil {
		hlog.Error("NewConnPgsql", "Error opening database: %q", err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			hlog.Error("NewConnPgsql", "Connection attempt timed out")
		} else {
			hlog.Error("NewConnPgsql", fmt.Sprintf("Error connecting to the database: %v", err))
		}
	}

	return db
}

func extractDBDetails(jdbcURL string) (string, string, string, string, string, error) {
	re := regexp.MustCompile(`^postgres://(.+):(.+)@([^:/?#]+):(\d+)/([^/?#]+)`)

	match := re.FindStringSubmatch(jdbcURL)

	if len(match) != 6 {
		hlog.Error("extractDBDetails", "Error parsing connection string")
		return "", "", "", "", "", errors.New("error parsing connection string")
	}

	user := match[1]
	password := match[2]
	host := match[3]
	port := match[4]
	dbname := match[5]

	return host, port, dbname, user, password, nil
}

func (c *connection) CreateScripts(sqlFile string) {
	fileContent, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("failed to read file %s: %v", sqlFile, err)
	}

	db, err := sql.Open("postgres", c.url)
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	defer db.Close()
	_, err = db.Exec(string(fileContent))
	if err != nil {
		log.Fatalf("failed to execute query %s: %v", sqlFile, err)
	}
}
