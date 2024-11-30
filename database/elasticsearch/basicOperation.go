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
	Update(ctx context.Context, elasticIndex string, query map[string]interface{}) (*Response, error)
	FindAll(ctx context.Context, elasticIndex string, query map[string]interface{}) (*Response, error)
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

	if res.IsError() {
		all, err := io.ReadAll(res.Body)
		if err != nil {
			hlog.Error("ElasticConn.Create", fmt.Sprintf("Error when create document: %v", err))
		}
		hlog.Error("ElasticConn.Create", fmt.Sprintf("Error when create document: %v", string(all)))
		return nil, errors.New("error when create document")
	}
	return c.decodeElasticResponse(ctx, res.Body)
}

func (c *connection) FindById(ctx context.Context, elasticIndex, id string) (*Response, error) {
	res, err := c.elastic.Get(elasticIndex, id)

	if err != nil {
		hlog.Error("ElasticConn.Create", fmt.Sprintf("ElasticIndex: %s -> Erro when decodify body: %v", elasticIndex, err))
		return nil, err
	}

	return c.decodeElasticResponse(ctx, res.Body)
}

func (c *connection) Update(ctx context.Context, elasticIndex string, query map[string]interface{}) (*Response, error) {

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		hlog.Error("ElasticConn.Update", fmt.Sprintf("Error when serialize query: %v", err))
	}

	res, err := c.elastic.UpdateByQuery(
		[]string{elasticIndex},
		c.elastic.UpdateByQuery.WithBody(&buf),
		c.elastic.UpdateByQuery.WithContext(context.Background()),
	)

	if err != nil {
		hlog.Error("ElasticConn.Update", fmt.Sprintf("Error when update document: %v", err))
	}

	if res.IsError() {
		hlog.Error("ElasticConn.Update", fmt.Sprintf("Error when update document: %v", res.String()))
		return nil, err
	}

	return c.decodeElasticResponse(ctx, res.Body)
}

func (c *connection) FindAll(ctx context.Context, elasticIndex string, query map[string]interface{}) (*Response, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		hlog.Error("ElasticConn.FindAll", fmt.Sprintf("Error when serialize query: %v", err))
	}

	res, err := c.elastic.Search(
		c.elastic.Search.WithContext(context.Background()),
		c.elastic.Search.WithIndex(elasticIndex),
		c.elastic.Search.WithBody(&buf),
	)

	if err != nil {
		hlog.Error("ElasticConn.FindAll", fmt.Sprintf("Error when update document: %v", err))
	}

	if res.IsError() {
		hlog.Error("ElasticConn.FindAll", fmt.Sprintf("Error when update document: %v", res.String()))
		return nil, err
	}

	return c.decodeElasticResponse(ctx, res.Body)
}

func (c *connection) decodeElasticResponse(ctx context.Context, res io.ReadCloser) (*Response, error) {
	defer res.Close()
	returnResponse := &Response{}
	if err := json.NewDecoder(res).Decode(returnResponse); err != nil {
		hlog.Error("ElasticConn.Create", fmt.Sprintf("Erro when decodify body: %v", err))
		return nil, err
	}
	return returnResponse, nil
}
