package database

import (
	"errors"
	"fmt"

	"key-value-database/internal/database/compute"
)

const (
	errInvalidRequestMsg = "invalid request"
)

//go:generate docker run --rm -v ${PWD}:/app -w /app/internal/database/ vektra/mockery --name Parser --inpackage --case=snake
type Parser interface {
	Parse(input string) (string, []string, error)
}

//go:generate docker run --rm -v ${PWD}:/app -w /app/internal/database/ vektra/mockery --name Engine --inpackage --case=snake
type Engine interface {
	Set(key, value string)
	Get(key string) (string, error)
	Del(key string)
}

type Database struct {
	parser Parser
	engine Engine
}

func NewDatabase(parser Parser, engine Engine) *Database {
	return &Database{
		parser: parser,
		engine: engine,
	}
}

func (d *Database) HandleQuery(request string) (string, error) {
	cmd, args, err := d.parser.Parse(request)
	if err != nil {
		return "", err
	}
	switch cmd {
	case compute.SetCommand:
		d.engine.Set(args[0], args[1])
		return "OK", nil
	case compute.GetCommand:
		v, err := d.engine.Get(args[0])
		if err != nil {
			return "", err
		}
		return v, nil
	case compute.DelCommand:
		d.engine.Del(args[0])
		return "OK", nil
	}
	return "", errors.New(fmt.Sprintf("%s: %s", errInvalidRequestMsg, request))
}
