package sendresult

import "github.com/gofiber/fiber/v2"

//SendRes sends json result.
func SendRes(ctx *fiber.Ctx, status int, result interface{}) error {
	return ctx.Status(status).JSON(fiber.Map{"result": result})
}
