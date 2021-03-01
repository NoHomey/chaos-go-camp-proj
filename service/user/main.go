package main

import (
	"fmt"

	"github.com/NoHomey/chaos-go-camp-proj/misc/validator/valerrs"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/data"
	"github.com/go-playground/validator/v10"
)

func main() {
	user := data.User{
		Name:     "No",
		Email:    "ivo@test.com",
		Password: "password",
	}
	validate := validator.New()
	data.RegisterPasswordValidator(validate)
	err := validate.Struct(user).(validator.ValidationErrors)
	fmt.Println(valerrs.Fields(err))
}
