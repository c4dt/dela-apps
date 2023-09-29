package models

import (
	"github.com/c4dt/dela/core/access"
)

// NewIdentity creates a new identity
func NewIdentity(id string) *Identity {
	return &Identity{
		buffer: []byte(id),
	}
}

// Identity represents an identity
//
// - implements arc.Identity
type Identity struct {
	access.Identity
	buffer []byte
}

// MarshalText implements arc.Identity
func (i Identity) MarshalText() ([]byte, error) {
	return i.buffer, nil
}

// String implements arc.Identity
func (i Identity) String() string {
	return string(i.buffer)
}
