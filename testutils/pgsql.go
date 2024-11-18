package testutils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Pgsql(ctx context.Context) (string, func(context.Context) error, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	pgsqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not start container: %s", err))
	}
	host, err := pgsqlC.Host(ctx)
	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not get container host: %s", err))
	}

	port, err := pgsqlC.MappedPort(ctx, "5432")
	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not get container port: %s", err))
	}
	return fmt.Sprintf("postgres://test:test@%s:%s/testdb?sslmode=disable", host, port.Port()), pgsqlC.Terminate, nil
}

func CreateScripts(host, sqlFile string) {
	fileContent, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("failed to read file %s: %v", sqlFile, err)
	}

	db, err := sql.Open("postgres", host)
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	defer db.Close()
	_, err = db.Exec(string(fileContent))
	if err != nil {
		log.Fatalf("failed to execute query %s: %v", sqlFile, err)
	}
}
