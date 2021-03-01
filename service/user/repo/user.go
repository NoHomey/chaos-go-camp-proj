package repo

import (
	"github.com/NoHomey/chaos-go-camp-proj/service/user/data"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/model"
)

//UserRepo is an abstraction for user repository.
type UserRepo interface {
	Create(data data.User) error
	Find(email string, password string) (model.User, error)
}
