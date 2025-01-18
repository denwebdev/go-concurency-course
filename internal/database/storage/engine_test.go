package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestSetAndGet(t *testing.T) {
	logger := zap.L()
	engine := NewInMemoryEngine(logger)
	engine.Set("testKey", "testValue")
	v, err := engine.Get("testKey")
	assert.NoError(t, err)
	assert.Equal(t, "testValue", v)
}

func TestDelAndGet(t *testing.T) {
	logger := zap.L()
	engine := NewInMemoryEngine(logger)
	engine.Set("testKey", "testValue")
	engine.Del("testKey")
	v, err := engine.Get("testKey")
	assert.Error(t, err)
	assert.Equal(t, "", v)
}
