package commands

import "errors"

type LPushCommand struct {
	*BaseCommand
}

func (c *LPushCommand) Execute(args []string) (interface{}, error) {
	if len(args) < 2 {
		return nil, errors.New("LPUSH command requires at least two arguments")
	}
	// Implementation left as an exercise
	return nil, errors.New("LPUSH command not implemented")
}

// Add more list commands (RPUSH, LRANGE) here
