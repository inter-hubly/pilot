package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type connection struct {
	username string
	password string
	url      []string
	elastic  *elasticsearch.Client
}

// Option was a function optional pattern
type Option func(*connection)

func WithUrl(url []string) Option {
	return func(a *connection) {
		a.url = url
	}
}
func WithUsernameAndPassword(username, password string) Option {
	return func(a *connection) {
		a.username = username
		a.password = password
	}
}

func NewConn(opts ...Option) {
	onceElastic.Do(func() {

		elasticConn.url = []string{elasticDefault}
		for _, opt := range opts {
			opt(elasticConn)
		}
		client, err := elasticsearch.NewClient(elasticsearch.Config{
			Addresses: elasticConn.url,
			Username:  elasticConn.username,
			Password:  elasticConn.password,
		})
		if err != nil {
			panic(err)
		}
		elasticConn.elastic = client
	})
}
