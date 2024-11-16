package elasticsearch

import (
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	onceElastic    sync.Once
	elasticConn    *connection
	elasticDefault = "http://localhost:9200"
)

type ElasticConn interface {
	GetConnection() *elasticsearch.Client
}

func GetConnection() *connection {
	return elasticConn
}
