package model

import (
	"time"

	"github.com/google/uuid"
)

//User represents the user model.
type User interface {
	ID() uuid.UUID
	Name() string
	Email() string
	PasswordHash() []byte
	RegisteredAt() time.Time
	LastModifiedAt() time.Time
}
