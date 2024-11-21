package testutils

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func RabbitMq(ctx context.Context) (string, func(context.Context) error, error) {
	rabbitmqContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "rabbitmq:3-management",
			ExposedPorts: []string{"5672/tcp", "15672/tcp"},
			WaitingFor:   wait.ForLog("Server startup complete"),
		},
		Started: true,
	})

	if err != nil {
		log.Fatalf("falha ao iniciar o contêiner do Elasticsearch: %v", err)
		return "", nil, err
	}
	host, err := rabbitmqContainer.Host(ctx)
	if err != nil {
		log.Fatalf("falha ao obter o host do contêiner: %v", err)
		rabbitmqContainer.Terminate(ctx)
		return "", nil, err
	}
	port, err := rabbitmqContainer.MappedPort(ctx, "5672")
	if err != nil {
		log.Fatalf("falha ao mapear a porta do contêiner: %v", err)
		rabbitmqContainer.Terminate(ctx)
		return "", nil, err
	}
	amqpURL := fmt.Sprintf("amqp://guest:guest@%s:%s/", host, port.Port())
	return amqpURL, rabbitmqContainer.Terminate, nil
}
