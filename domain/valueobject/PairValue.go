package valueobject

type Pair[T, U any] struct {
	Key   T `json:"key"`
	Value U `json:"value"`
}
