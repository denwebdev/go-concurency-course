package storage

import (
	"errors"
)

const (
	errNotExists = "value by key not exists"
)

type InMemoryEngine struct {
	data map[string]string
}

func NewInMemoryEngine() *InMemoryEngine {
	return &InMemoryEngine{
		data: make(map[string]string),
	}
}

func (e *InMemoryEngine) Set(key, value string) {
	e.data[key] = value
}

func (e *InMemoryEngine) Get(key string) (string, error) {
	v, exists := e.data[key]
	if !exists {
		return "", errors.New(errNotExists)
	}
	return v, nil
}

func (e *InMemoryEngine) Del(key string) {
	delete(e.data, key)
}
