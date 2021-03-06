package request

import (
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/expose/reqlogger"
	"github.com/gofiber/fiber/v2"
)

//FromFiberCtx constructs Request value.
func FromFiberCtx(ctx *fiber.Ctx) reqlogger.Request {
	return reqlogger.Request{
		Method: ctx.Method(),
		URL:    ctx.OriginalURL(),
		IP:     ctx.IP(),
	}
}

//SetTime sets current time to locals.
func SetTime(ctx *fiber.Ctx) {
	ctx.Locals(localKey, time.Now())
}

//GetDuration returns duration since the setted time.
func GetDuration(ctx *fiber.Ctx) time.Duration {
	now := time.Now()
	start := ctx.Locals(localKey, now).(time.Time)
	return now.Sub(start)
}

const localKey = "$__request-start-time__$"
