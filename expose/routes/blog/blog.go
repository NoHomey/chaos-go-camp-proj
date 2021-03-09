package blog

import (
	blogh "github.com/NoHomey/chaos-go-camp-proj/expose/handler/blog"
	"github.com/NoHomey/chaos-go-camp-proj/expose/middleware/auth"
	"github.com/NoHomey/chaos-go-camp-proj/expose/reqlogger"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/access"
	"github.com/gofiber/fiber/v2"
)

//Register registers the blog routes.
func Register(
	app *fiber.App,
	accessService access.Service,
	service blog.Service,
	logger reqlogger.Logger,
) {
	handler := blogh.Handler{
		Service:   service,
		ReqLogger: logger,
	}
	authMiddleware := auth.Middleware(accessService)
	app.Post("/blog", authMiddleware, handler.Save)
	app.Post("/blogs", authMiddleware, handler.Fetch)
}
