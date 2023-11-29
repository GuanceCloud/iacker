package singleton

import (
	"reflect"
	"sync"
)

var cache sync.Map

type Singleton[T any] struct {
	v T
}

func NewSingleton[T any]() (t *Singleton[T]) {
	id := reflect.TypeOf(t).String()
	v, ok := cache.Load(id)
	if ok {
		return v.(*Singleton[T])
	}
	v = new(*Singleton[T])
	v, _ = cache.LoadOrStore(id, v)
	return v.(*Singleton[T])
}
