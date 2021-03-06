package logger

import (
	"github.com/NoHomey/chaos-go-camp-proj/expose/logger"
	"github.com/NoHomey/chaos-go-camp-proj/expose/utility/request"
	"github.com/gofiber/fiber/v2"
)

//Register registers the middleware.
func Register(app *fiber.App, l logger.Logger) {
	app.Use(func(ctx *fiber.Ctx) error {
		request.SetTime(ctx)
		l.Request(request.FromFiberCtx(ctx))
		return ctx.Next()
	})
}
