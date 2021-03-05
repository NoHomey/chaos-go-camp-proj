package user

import (
	"github.com/NoHomey/chaos-go-camp-proj/expose/handler/user"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/prime"
	"github.com/gofiber/fiber/v2"
)

//Register registers the user routes.
func Register(app *fiber.App, service prime.Service) {
	router := app.Group("/user")
	router.Post("/sign-up", user.SignUp(service))
	router.Post("/sign-in", user.SignIn(service))
	router.Post("/sign-out", user.SignOut(service))
	router.Get("/access", user.ObtainAccess(service))
}
