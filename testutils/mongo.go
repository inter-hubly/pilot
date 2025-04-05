package testutils

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Mongo(ctx context.Context) (string, func(context.Context) error, error) {

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
	}

	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not start container: %s", err))
	}
	host, err := mongoC.Host(ctx)
	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not get container host: %s", err))
	}

	port, err := mongoC.MappedPort(ctx, "27017")
	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not get container port: %s", err))
	}
	return fmt.Sprintf("mongodb://%s:%s", host, port.Port()), func(context.Context) error {
		log.Print("Execute close mongo")
		return mongoC.Terminate(ctx)
	}, nil
}
