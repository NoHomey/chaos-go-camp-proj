package data

import (
	"github.com/NoHomey/chaos-go-camp-proj/data/tag"
	"github.com/go-playground/validator/v10"
)

//RegisterValidators registers field validators
func RegisterValidators(validate *validator.Validate) {
	tag.RegisterValidator(validate)
}
