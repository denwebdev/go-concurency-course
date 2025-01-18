package compute

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	t.Run("Successful SET command", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		compute := NewCompute(parser, engine)

		input := "SET key value"
		parser.On("Parse", input).Return(SetCommand, []string{"key", "value"}, nil)
		engine.On("Set", "key", "value").Return()

		result, err := compute.Process(input)
		assert.NoError(t, err)
		assert.Equal(t, "OK", result)

		parser.AssertExpectations(t)
		engine.AssertExpectations(t)
	})

	t.Run("Successful GET command", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		compute := NewCompute(parser, engine)

		input := "GET key"
		parser.On("Parse", input).Return(GetCommand, []string{"key"}, nil)
		engine.On("Get", "key").Return("value", nil)

		result, err := compute.Process(input)
		assert.NoError(t, err)
		assert.Equal(t, "value", result)

		parser.AssertExpectations(t)
		engine.AssertExpectations(t)
	})

	t.Run("Successful DEL command", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		compute := NewCompute(parser, engine)

		input := "DEL key"
		parser.On("Parse", input).Return(DelCommand, []string{"key"}, nil)
		engine.On("Del", "key").Return()

		result, err := compute.Process(input)
		assert.NoError(t, err)
		assert.Equal(t, "OK", result)

		parser.AssertExpectations(t)
		engine.AssertExpectations(t)
	})

	t.Run("Invalid command", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		compute := NewCompute(parser, engine)

		input := "INVALID key"
		parser.On("Parse", input).Return("INVALID", nil, nil)

		result, err := compute.Process(input)
		assert.Error(t, err)
		assert.Equal(t, "", result)
		assert.Contains(t, err.Error(), "invalid input")

		parser.AssertExpectations(t)
	})

	t.Run("Parser error", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		compute := NewCompute(parser, engine)

		input := "BROKEN"
		parser.On("Parse", input).Return("", nil, errors.New("parse error"))

		result, err := compute.Process(input)
		assert.Error(t, err)
		assert.Equal(t, "", result)
		assert.Equal(t, "parse error", err.Error())

		parser.AssertExpectations(t)
	})

	t.Run("GET command with missing key", func(t *testing.T) {
		parser := new(MockParser)
		engine := new(MockEngine)
		compute := NewCompute(parser, engine)

		input := "GET missing_key"
		parser.On("Parse", input).Return(GetCommand, []string{"missing_key"}, nil)
		engine.On("Get", "missing_key").Return("", errors.New("key not found"))

		result, err := compute.Process(input)
		assert.Error(t, err)
		assert.Equal(t, "", result)
		assert.Equal(t, "key not found", err.Error())

		parser.AssertExpectations(t)
		engine.AssertExpectations(t)
	})
}
