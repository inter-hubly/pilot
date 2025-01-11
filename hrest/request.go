package hrest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/inter-hubly/pilot/hlog"
	"github.com/pkg/errors"
)

type Request struct {
	Headers      map[string]string
	Url          string
	Body         any
	responseBody []byte
}

type Option func(*Request)

func WithBody(body any) Option {
	return func(r *Request) {
		r.Body = body
	}
}

type Pair[K string, V string] struct {
	Key   K
	Value V
}

func WithHeader(pair []Pair[string, string]) Option {
	return func(r *Request) {
		for _, p := range pair {
			r.Headers[p.Key] = p.Value
		}
	}
}

func NewRequest(url string, options ...Option) *Request {
	req := &Request{
		Headers: make(map[string]string),
	}

	for _, opt := range options {
		opt(req)
	}

	if req.Headers == nil {
		req.Headers = map[string]string{
			"Content-Type": "application/json",
		}
	}

	req.Url = url

	return req
}

func (r *Request) CreateRequest(ctx context.Context, httpMethod string) error {
	body, err := json.Marshal(r.Body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}
	req, err := http.NewRequest(httpMethod, r.Url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	r.responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed (%d): %s", resp.StatusCode, string(r.responseBody))
	}

	return nil
}

func (r *Request) GetBody(ctx context.Context, result any) error {
	hlog.Debug(ctx, "Request.GetBody", fmt.Sprintf("GetBody :%s", r.Body))
	if err := json.Unmarshal(r.responseBody, &result); err != nil {
		return errors.New("can't get body")
	}

	return nil
}
