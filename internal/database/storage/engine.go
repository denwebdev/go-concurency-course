package storage

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
)

const (
	errNotExists = "value by key not exists"
)

type InMemoryEngine struct {
	logger *zap.Logger
	data   map[string]string
}

func NewInMemoryEngine(logger *zap.Logger) *InMemoryEngine {
	return &InMemoryEngine{
		logger: logger,
		data:   make(map[string]string),
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
	_, exists := e.data[key]
	if !exists {
		e.logger.Info(fmt.Sprintf("value by key: %s not exists", key))
		return
	}
	delete(e.data, key)
}
