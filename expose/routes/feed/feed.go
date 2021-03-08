package feed

import (
	feedh "github.com/NoHomey/chaos-go-camp-proj/expose/handler/feed"
	"github.com/NoHomey/chaos-go-camp-proj/expose/reqlogger"
	"github.com/NoHomey/chaos-go-camp-proj/service/feed"
	"github.com/gofiber/fiber/v2"
)

//Register registers the user routes.
func Register(app *fiber.App, service feed.Service, l reqlogger.Logger) {
	handler := feedh.Handler{
		Service:   service,
		ReqLogger: l,
	}
	path := "/feed/details/:" + feedh.Param
	app.Get(path, handler.Details)
}
