package data

import (
	"regexp"

	"github.com/NoHomey/chaos-go-camp-proj/service/blog/enum/level"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog/enum/rating"
	"github.com/go-playground/validator/v10"
)

//Blog is the blog data.
type Blog struct {
	FeedURL     string   `json:"feedURL" validate:"url"`
	Author      string   `json:"author" valiadate:"required"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Rating      uint8    `json:"rating" validate:"rating"`
	Level       uint8    `json:"level" validate:"level"`
	Tags        []string `json:"tags" validate:"tags"`
	QuickNote   string   `json:"qickNote"`
}

//FetchBlogs is data for fetching blogs.
type FetchBlogs struct {
	Tags  []string `json:"tags" validate:"tags"`
	Count uint32   `json:"count" validate:"min=10"`
	After string   `json:"after" validate:"hexadecimal"`
}

//RegisterValidators registers field validators
func RegisterValidators(validate *validator.Validate) {
	validate.RegisterValidation("rating", func(fl validator.FieldLevel) bool {
		return fl.Field().Interface().(uint8) <= rating.MaxNum
	})
	validate.RegisterValidation("level", func(fl validator.FieldLevel) bool {
		return fl.Field().Interface().(uint8) <= level.MaxNum
	})
	validate.RegisterValidation("tags", func(fl validator.FieldLevel) bool {
		list := fl.Field().Interface().([]string)
		for i := range list {
			if !tagRegExp.MatchString(list[i]) {
				return false
			}
		}
		return true
	})
}

var tagRegExp = regexp.MustCompile("\\w+(?:(?:-|_)\\w+)*")
