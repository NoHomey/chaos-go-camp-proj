package repo

import (
	"github.com/NoHomey/chaos-go-camp-proj/service/user/data"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/model"
	"github.com/google/uuid"
)

//AccessRepo is an abstraction for access repository.
type AccessRepo interface {
	Create(data data.Access) error
	FindByID(id uuid.UUID) (model.Access, error)
}
