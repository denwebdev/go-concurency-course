package compute

import (
	"errors"
	"fmt"
)

const (
	errInvalidInput = "invalid input"
)

//go:generate docker run --rm -v ${PWD}:/app -w /app/internal/database/compute/ vektra/mockery --name Parser --inpackage --case=snake
type Parser interface {
	Parse(input string) (string, []string, error)
}

//go:generate docker run --rm -v ${PWD}:/app -w /app/internal/database/compute/ vektra/mockery --name Engine --inpackage --case=snake
type Engine interface {
	Set(key, value string)
	Get(key string) (string, error)
	Del(key string)
}

type Compute struct {
	parser Parser
	engine Engine
}

func NewCompute(
	parser Parser,
	engine Engine,
) *Compute {
	return &Compute{
		parser: parser,
		engine: engine,
	}
}

func (c *Compute) Process(input string) (string, error) {
	cmd, args, err := c.parser.Parse(input)
	if err != nil {
		return "", err
	}
	switch cmd {
	case SetCommand:
		c.engine.Set(args[0], args[1])
		return "OK", nil
	case GetCommand:
		v, err := c.engine.Get(args[0])
		if err != nil {
			return "", err
		}
		return v, nil
	case DelCommand:
		c.engine.Del(args[0])
		return "OK", nil
	}
	return "", errors.New(fmt.Sprintf("%s: %s", errInvalidInput, input))
}
