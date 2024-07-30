package commands

import "errors"

type GetCommand struct {
	*BaseCommand
}

func (c *GetCommand) Execute(args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, errors.New("GET command requires exactly one argument")
	}
	value, exists := c.store.Get(args[0])
	if !exists {
		return nil, nil // Redis returns nil for non-existent keys
	}
	return value, nil
}

type SetCommand struct {
	*BaseCommand
}

func (c *SetCommand) Execute(args []string) (interface{}, error) {
	if len(args) != 2 {
		return nil, errors.New("SET command requires exactly two arguments")
	}
	c.store.Set(args[0], args[1])
	return "OK", nil
}
