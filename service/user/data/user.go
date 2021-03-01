package data

import "github.com/go-playground/validator/v10"

//User represents the user data.
type User struct {
	Name     string `json:"name" validate:"min=3,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,max=32,password"`
}

//RegisterPasswordValidator registers validator for password.
func RegisterPasswordValidator(validate *validator.Validate) {
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return isPasswordValid(fl.Field().String())
	})
}

func isPasswordValid(s string) bool {
	digit := false
	lowerCase := false
	upperCase := false
	for _, c := range s {
		if inRange(c, '0', '9') {
			digit = true
			continue
		}
		if inRange(c, 'a', 'z') {
			lowerCase = true
			continue
		}
		if inRange(c, 'A', 'Z') {
			upperCase = true
		}
	}
	return digit && lowerCase && upperCase
}

func inRange(x, a, b rune) bool {
	return x >= a && x <= b
}
