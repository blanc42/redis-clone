package commands

import "redis-clone/internal/datastore"

type Commander interface {
	Execute(args []string) (interface{}, error)
}

type BaseCommand struct {
	store datastore.DataStore
}

func NewBaseCommand(store datastore.DataStore) *BaseCommand {
	return &BaseCommand{store: store}
}
