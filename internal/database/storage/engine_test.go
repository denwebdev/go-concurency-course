package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAndGet(t *testing.T) {
	engine := NewInMemoryEngine()
	engine.Set("testKey", "testValue")
	v, err := engine.Get("testKey")
	assert.NoError(t, err)
	assert.Equal(t, "testValue", v)
}

func TestDelAndGet(t *testing.T) {
	engine := NewInMemoryEngine()
	engine.Set("testKey", "testValue")
	engine.Del("testKey")
	v, err := engine.Get("testKey")
	assert.Error(t, err)
	assert.Equal(t, "", v)
}
