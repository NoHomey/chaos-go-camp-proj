package sendresult

import "github.com/gofiber/fiber/v2"

//SendJSON sends json result.
func SendJSON(ctx *fiber.Ctx, status int, result interface{}) error {
	return ctx.Status(status).JSON(fiber.Map{"result": result})
}
