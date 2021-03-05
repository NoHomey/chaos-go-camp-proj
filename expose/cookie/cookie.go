package cookie

import "github.com/gofiber/fiber/v2"

//Set sets a cookie.
func Set(ctx *fiber.Ctx, name string, val string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    val,
		HTTPOnly: true,
		SameSite: "Strict",
	})
}
