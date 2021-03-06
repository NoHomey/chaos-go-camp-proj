package user

import (
	"github.com/NoHomey/chaos-go-camp-proj/expose/handler/user"
	"github.com/NoHomey/chaos-go-camp-proj/expose/reqlogger"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/prime"
	"github.com/gofiber/fiber/v2"
)

//Register registers the user routes.
func Register(app *fiber.App, service prime.Service, l reqlogger.Logger) {
	router := app.Group("/user")
	handler := user.Handler{
		Service:   service,
		ReqLogger: l,
	}
	router.Post("/sign-up", handler.SignUp)
	router.Post("/sign-in", handler.SignIn)
	router.Post("/sign-out", handler.SignOut)
	router.Get("/access", handler.ObtainAccess)
}
