package storage

import "github.com/c4dt/dela/serde"

// KeyValue defines a simple key value storage
type KeyValue interface {
	Store(key []byte, value serde.Message) error
	Read(key []byte) (serde.Message, error)
}
