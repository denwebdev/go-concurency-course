package database

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"key-value-database/internal/database/compute"
)

func TestProcess(t *testing.T) {
	t.Run("Successful SET command", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		db := NewDatabase(parser, engine)

		request := "SET key value"
		parser.On("Parse", request).Return(compute.SetCommand, []string{"key", "value"}, nil)
		engine.On("Set", "key", "value").Return()

		result, err := db.HandleQuery(request)
		assert.NoError(t, err)
		assert.Equal(t, "OK", result)

		parser.AssertExpectations(t)
		engine.AssertExpectations(t)
	})

	t.Run("Successful GET command", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		db := NewDatabase(parser, engine)

		request := "GET key"
		parser.On("Parse", request).Return(compute.GetCommand, []string{"key"}, nil)
		engine.On("Get", "key").Return("value", nil)

		result, err := db.HandleQuery(request)
		assert.NoError(t, err)
		assert.Equal(t, "value", result)

		parser.AssertExpectations(t)
		engine.AssertExpectations(t)
	})

	t.Run("Successful DEL command", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		db := NewDatabase(parser, engine)

		request := "DEL key"
		parser.On("Parse", request).Return(compute.DelCommand, []string{"key"}, nil)
		engine.On("Del", "key").Return()

		result, err := db.HandleQuery(request)
		assert.NoError(t, err)
		assert.Equal(t, "OK", result)

		parser.AssertExpectations(t)
		engine.AssertExpectations(t)
	})

	t.Run("Invalid command", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		db := NewDatabase(parser, engine)

		request := "INVALID key"
		parser.On("Parse", request).Return("INVALID", nil, nil)

		result, err := db.HandleQuery(request)
		assert.Error(t, err)
		assert.Equal(t, "", result)
		assert.Contains(t, err.Error(), "invalid request")

		parser.AssertExpectations(t)
	})

	t.Run("Parser error", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		db := NewDatabase(parser, engine)

		request := "BROKEN"
		parser.On("Parse", request).Return("", nil, errors.New("parse error"))

		result, err := db.HandleQuery(request)
		assert.Error(t, err)
		assert.Equal(t, "", result)
		assert.Equal(t, "parse error", err.Error())

		parser.AssertExpectations(t)
	})

	t.Run("GET command with missing key", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		db := NewDatabase(parser, engine)

		request := "GET missing_key"
		parser.On("Parse", request).Return(compute.GetCommand, []string{"missing_key"}, nil)
		engine.On("Get", "missing_key").Return("", errors.New("key not found"))

		result, err := db.HandleQuery(request)
		assert.Error(t, err)
		assert.Equal(t, "", result)
		assert.Equal(t, "key not found", err.Error())

		parser.AssertExpectations(t)
		engine.AssertExpectations(t)
	})
}
