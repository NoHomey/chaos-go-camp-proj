package tag

import (
	"regexp"

	"github.com/NoHomey/chaos-go-camp-proj/data/enum/priority"
	"github.com/go-playground/validator/v10"
)

//Tag represents tag data.
type Tag struct {
	Value    string            `json:"value" bson:"value" validate:"tag"`
	Priority priority.Priority `json:"priority" bson:"priority"`
}

//Tags returns tag values.
func Tags(tags []Tag) []string {
	vals := make([]string, len(tags))
	for i := range tags {
		vals[i] = tags[i].Value
	}
	return vals
}

//RegisterValidator registers field validator.
func RegisterValidator(validate *validator.Validate) {
	validate.RegisterValidation("tag", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		return tagRegExp.MatchString(val)
	})
}

var tagRegExp = regexp.MustCompile("\\w+(?:(?:-|_)\\w+)*")
