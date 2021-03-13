package data

import (
	"github.com/NoHomey/chaos-go-camp-proj/data/enum/level"
	"github.com/NoHomey/chaos-go-camp-proj/data/enum/rating"
	"github.com/go-playground/validator/v10"
)

//RegisterValidators registers field validators
func RegisterValidators(validate *validator.Validate) {
	level.RegisterValidator(validate)
	rating.RegisterValidator(validate)
}
