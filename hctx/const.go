package hctx

import "context"

var (
	Tenant Value[string] = "tenant"
)

type Value[K any] string

func (v Value[K]) New(value K) context.Context {
	return context.WithValue(context.Background(), v, value)
}

func (v Value[K]) Get(ctx context.Context) K {
	res := ctx.Value(v)
	if value, ok := res.(K); ok {
		return value
	}

	return *new(K)
}
