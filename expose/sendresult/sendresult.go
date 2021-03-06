package sendresult

import (
	"github.com/NoHomey/chaos-go-camp-proj/expose/logger"
	"github.com/NoHomey/chaos-go-camp-proj/expose/utility/request"
	"github.com/gofiber/fiber/v2"
)

//SendAndLog sends json result and logs request end.
func SendAndLog(ctx *fiber.Ctx, status int, result interface{}, l logger.Logger) error {
	l.Response(request.FromFiberCtx(ctx), request.GetDuration(ctx), status)
	return ctx.Status(status).JSON(fiber.Map{"result": result})
}
