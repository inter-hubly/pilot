package elasticsearch

import (
	"sync"
)

var (
	onceElastic    sync.Once
	elasticConn    *connection
	elasticDefault = "http://localhost:9200"
)

func GetConnection() *connection {
	return elasticConn
}
