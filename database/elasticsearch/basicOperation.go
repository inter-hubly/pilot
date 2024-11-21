package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/inter-hubly/pilot/hlog"
	"github.com/pkg/errors"
)

type ElasticConn interface {
	Create(ctx context.Context, elasticIndex string, v any) (*Response, error)
	FindById(ctx context.Context, elasticIndex string, id string) (*Response, error)
}

func (c *connection) Create(ctx context.Context, elasticIndex string, doc any) (*Response, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		hlog.Error("ElasticConn.Create", fmt.Sprintf("error when serialize doc: %v", err))
	}

	req := esapi.IndexRequest{
		Index:   elasticIndex,
		Body:    &buf,
		Refresh: "true",
	}

	res, err := req.Do(ctx, c.elastic)
	if err != nil {
		hlog.Error("ElasticConn.Create", fmt.Sprintf("Error when create document: %v", err))
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		all, err := io.ReadAll(res.Body)
		if err != nil {
			hlog.Error("ElasticConn.Create", fmt.Sprintf("Error when create document: %v", err))
		}
		hlog.Error("ElasticConn.Create", fmt.Sprintf("Error when create document: %v", all))
		return nil, errors.New("error when create document")
	}
	var resBody Response
	if err = json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		hlog.Error("ElasticConn.Create", fmt.Sprintf("Erro when decodify body: %v", err))
	}

	return &resBody, nil
}

func (c *connection) FindById(ctx context.Context, elasticIndex, id string) (*Response, error) {
	res, err := c.elastic.Get(elasticIndex, id)

	if err != nil {
		hlog.Error("ElasticConn.Create", fmt.Sprintf("ElasticIndex: %s -> Erro when decodify body: %v", elasticIndex, err))
		return nil, err
	}

	defer res.Body.Close()
	returnResponse := &Response{}
	if err = json.NewDecoder(res.Body).Decode(returnResponse); err != nil {
		hlog.Error("ElasticConn.Create", fmt.Sprintf("ElasticIndex: %s -> Erro when decodify body: %v", elasticIndex, err))
		return nil, err
	}
	return returnResponse, nil
}
