package testutils

import (
	"context"
	"log"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func ElasticSearch(ctx context.Context) (string, func(context.Context) error, error) {
	esReq := testcontainers.ContainerRequest{
		Image:        "docker.elastic.co/elasticsearch/elasticsearch:8.10.2",
		ExposedPorts: []string{"9200/tcp"},
		Env: map[string]string{
			"discovery.type":         "single-node",
			"xpack.security.enabled": "false",
		},
		WaitingFor: wait.ForHTTP("/").WithPort("9200").WithStartupTimeout(2 * time.Minute),
	}

	esContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: esReq,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("falha ao iniciar o contêiner do Elasticsearch: %v", err)
		return "", nil, err
	}
	host, err := esContainer.Host(ctx)
	if err != nil {
		log.Fatalf("falha ao obter o host do contêiner: %v", err)
		esContainer.Terminate(ctx)
		return "", nil, err
	}
	port, err := esContainer.MappedPort(ctx, "9200")
	if err != nil {
		log.Fatalf("falha ao mapear a porta do contêiner: %v", err)
		esContainer.Terminate(ctx)
		return "", nil, err
	}
	esURL := "http://" + host + ":" + port.Port()
	return esURL, esContainer.Terminate, nil
}

func StartElastic() {

}
