package compute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		wantCmd  string
		wantArgs []string
		wantErr  error
	}{
		{
			name:     "on empty input",
			input:    "",
			wantCmd:  "",
			wantArgs: nil,
			wantErr:  errEmptyCommand,
		},
		{
			name:     "on space input",
			input:    " ",
			wantCmd:  "",
			wantArgs: nil,
			wantErr:  errEmptyCommand,
		},
		{
			name:     "on empty args",
			input:    "CMD",
			wantCmd:  "",
			wantArgs: nil,
			wantErr:  errEmptyArgs,
		},
		{
			name:     "on unknown command",
			input:    "CMD arg",
			wantCmd:  "",
			wantArgs: nil,
			wantErr:  errUnknownCommand,
		},
		{
			name:     "on SET command with 1 arguments",
			input:    "SET arg",
			wantCmd:  "",
			wantArgs: nil,
			wantErr:  errInvalidArgs,
		},
		{
			name:     "on SET command with 2 arguments",
			input:    "SET arg1 arg2",
			wantCmd:  "SET",
			wantArgs: []string{"arg1", "arg2"},
			wantErr:  nil,
		},
		{
			name:     "on SET command with 3 arguments",
			input:    "SET arg1 arg2 arg3",
			wantCmd:  "",
			wantArgs: nil,
			wantErr:  errInvalidArgs,
		},
		{
			name:     "on GET command with 2 arguments",
			input:    "GET arg1 arg2",
			wantCmd:  "",
			wantArgs: nil,
			wantErr:  errInvalidArgs,
		},
		{
			name:     "on GET command with 1 argument",
			input:    "GET arg1",
			wantCmd:  "GET",
			wantArgs: []string{"arg1"},
			wantErr:  nil,
		},
		{
			name:     "on DEL command with 2 arguments",
			input:    "DEL arg1 arg2",
			wantCmd:  "",
			wantArgs: nil,
			wantErr:  errInvalidArgs,
		},
		{
			name:     "on DEL command with 1 argument",
			input:    "DEL arg1",
			wantCmd:  "DEL",
			wantArgs: []string{"arg1"},
			wantErr:  nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			parser := &CommandParser{}
			cmd, args, err := parser.Parse(c.input)
			assert.Equal(t, c.wantCmd, cmd)
			assert.Equal(t, c.wantArgs, args)
			if c.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, c.wantErr, err)
			}
		})
	}
}
