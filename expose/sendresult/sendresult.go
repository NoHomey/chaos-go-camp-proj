package sendresult

import (
	"github.com/NoHomey/chaos-go-camp-proj/expose/reqlogger"
	"github.com/NoHomey/chaos-go-camp-proj/expose/utility/request"
	"github.com/gofiber/fiber/v2"
)

//SendAndLog sends json result and logs request end.
func SendAndLog(ctx *fiber.Ctx, status int, result interface{}, l reqlogger.Logger) error {
	l.Response(request.FromFiberCtx(ctx), request.GetDuration(ctx), status)
	return ctx.Status(status).JSON(fiber.Map{"result": result})
}
