package registry

import "sync"

// Registry is the global mapping to store resource provider
type Registry[T any] struct {
	cache sync.Map
}

// NewRegistry will create new registry for resource provider
func NewRegistry[T any]() *Registry[T] {
	return &Registry[T]{
		cache: sync.Map{},
	}
}

// Register will register the resource provider.
func (r *Registry[T]) Register(name string, t T) error {
	r.cache.Store(name, t)
	return nil
}

// Get will return resource provider by name
func (r *Registry[T]) Get(name string) (T, bool) {
	t, ok := r.cache.Load(name)
	return t, ok
}
