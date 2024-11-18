package testutils

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Redis(ctx context.Context) (string, func(context.Context) error, error) {

	req := testcontainers.ContainerRequest{
		Image:        "redis:alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForListeningPort("6379/tcp"),
	}

	redisContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not start container: %s", err))
	}
	host, err := redisContainer.Host(ctx)
	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not get container host: %s", err))
	}

	port, err := redisContainer.MappedPort(ctx, "6379")
	if err != nil {
		return "", nil, errors.New(fmt.Sprintf("Could not get container port: %s", err))
	}
	return fmt.Sprintf("%s:%s", host, port.Port()), redisContainer.Terminate, nil
}
