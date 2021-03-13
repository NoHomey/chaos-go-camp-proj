package data

import (
	"github.com/NoHomey/chaos-go-camp-proj/data/enum/level"
	"github.com/NoHomey/chaos-go-camp-proj/data/enum/rating"
	"github.com/NoHomey/chaos-go-camp-proj/data/tag"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Blog is the blog data.
type Blog struct {
	FeedURL     string        `json:"feedURL" validate:"url"`
	Author      string        `json:"author" valiadate:"required"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Rating      rating.Rating `json:"rating"`
	Level       level.Level   `json:"level"`
	Tags        []tag.Tag     `json:"tags"`
	QuickNote   string        `json:"qickNote"`
}

//FetchBlogs is data for fetching blogs.
type FetchBlogs struct {
	Rating rating.Rating `json:"rating"`
	Level  level.Level   `json:"level"`
	Tags   []tag.Tag     `json:"tags"`
	Count  uint32        `json:"count" validate:"min=10"`
	After  string        `json:"after" validate:"optObjectID"`
}

//RegisterValidators registers field validators
func RegisterValidators(validate *validator.Validate) {
	validate.RegisterValidation("optObjectID", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		if len(val) == 0 {
			return true
		}
		_, err := primitive.ObjectIDFromHex(val)
		return err == nil
	})
}
