package tag

import (
	"regexp"

	"github.com/NoHomey/chaos-go-camp-proj/data/enum/priority"
	"github.com/go-playground/validator/v10"
)

//Tag represents tag data.
type Tag struct {
	value    string
	priority priority.Priority
}

//Tags returns tag values.
func Tags(tags []Tag) []string {
	vals := make([]string, len(tags))
	for i := range tags {
		vals[i] = tags[i].value
	}
	return vals
}

//RegisterValidator registers field validator.
func RegisterValidator(validate *validator.Validate) {
	validate.RegisterValidation("tags", func(fl validator.FieldLevel) bool {
		list := fl.Field().Interface().([]Tag)
		for i := range list {
			if !tagRegExp.MatchString(list[i].value) {
				return false
			}
		}
		return true
	})
}

var tagRegExp = regexp.MustCompile("\\w+(?:(?:-|_)\\w+)*")
