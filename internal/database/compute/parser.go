package compute

import (
	"errors"
	"strings"
)

var (
	errEmptyCommand   = errors.New("empty command")
	errEmptyArgs      = errors.New("empty arguments")
	errInvalidArgs    = errors.New("invalid arguments")
	errUnknownCommand = errors.New("unknown command")
)

type CommandParser struct{}

func (p *CommandParser) Parse(input string) (string, []string, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", nil, errEmptyCommand
	}
	if len(parts) == 1 {
		return "", nil, errEmptyArgs
	}

	cmd := parts[0]
	args := parts[1:]
	switch cmd {
	case SetCommand:
		if len(args) != 2 {
			return "", nil, errInvalidArgs
		}
		return cmd, args, nil
	case GetCommand, DelCommand:
		if len(args) != 1 {
			return "", nil, errInvalidArgs
		}
		return cmd, args, nil
	default:
		return "", nil, errUnknownCommand
	}
}
