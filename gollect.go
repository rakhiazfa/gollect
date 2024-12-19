package gollect

type Collection[T interface{}] struct {
	items  []T
	result []T
}

func Collect[T interface{}](items []T) *Collection[T] {
	result := make([]T, len(items))

	copy(result, items)

	return &Collection[T]{items, result}
}

func (collection *Collection[T]) Push(items ...T) *Collection[T] {
	collection.result = append(collection.result, items...)

	return collection
}

func (collection *Collection[T]) Map(callable func(T) T) *Collection[T] {
	for idx, item := range collection.result {
		collection.result[idx] = callable(item)
	}

	return collection
}

func (collection *Collection[T]) Filter(callable func(T) bool) *Collection[T] {
	var r []T

	for _, item := range collection.result {
		if callable(item) {
			r = append(r, item)
		}
	}

	collection.result = r

	return collection
}

func (collection *Collection[T]) Value() []T {
	return collection.result
}
