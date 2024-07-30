package datastore

type DataStore interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Delete(key string)
}
