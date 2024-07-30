package commands

import "errors"

type HSetCommand struct {
	*BaseCommand
}

func (c *HSetCommand) Execute(args []string) (interface{}, error) {
	if len(args) != 3 {
		return nil, errors.New("HSET command requires exactly three arguments")
	}
	// Implementation left as an exercise
	return nil, errors.New("HSET command not implemented")
}

// Add more hash commands (HGET, HDEL) here
