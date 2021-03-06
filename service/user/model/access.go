package model

import (
	"time"

	"github.com/google/uuid"
)

//Access represents the access model.
type Access interface {
	ID() uuid.UUID
	UserID() uuid.UUID
	UserEmail() string
	UserName() string
	CreatedAt() time.Time
}
