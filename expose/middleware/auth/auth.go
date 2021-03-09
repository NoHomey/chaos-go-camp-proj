package auth

import (
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/access"
	"github.com/gofiber/fiber/v2"
)

//AccessTokenKey is the cookie key for the access token.
const AccessTokenKey = "access-token"

//Middleware returns Auth middleware when given access service.
func Middleware(service access.Service) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		accessToken := ctx.Cookies(AccessTokenKey)
		syncToken := ctx.Get(accessSyncTokenHeaderKey)
		data, err := service.DecodeAndValidateAccessToken(access.SyncToken{
			Token: accessToken,
			Sync:  syncToken,
		})
		if err != nil {
			return err
		}
		ctx.Locals(accessDataKey, data)
		return ctx.Next()
	}
}

//AccessData returns access data from fiber's context.
func AccessData(ctx *fiber.Ctx) *access.TokenData {
	return ctx.Locals(accessDataKey).(*access.TokenData)
}

const accessDataKey = "$__access-data__$"

const accessSyncTokenHeaderKey = "X-Access-Sync-Token"
