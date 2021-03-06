package request

import (
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/expose/reqlogger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//FromFiberCtx constructs Request value.
func FromFiberCtx(ctx *fiber.Ctx) reqlogger.Request {
	id := ctx.Locals(localKeyForID).(string)
	return reqlogger.Request{
		Method: ctx.Method(),
		URL:    ctx.OriginalURL(),
		IP:     ctx.IP(),
		ID:     id,
	}
}

//SetTimeAndID sets current time and ads tracking id to locals.
func SetTimeAndID(ctx *fiber.Ctx) {
	ctx.Locals(localKeyForTime, time.Now())
	ctx.Locals(localKeyForID, uuid.NewString())
}

//GetDuration returns duration since the setted time.
func GetDuration(ctx *fiber.Ctx) time.Duration {
	now := time.Now()
	start := ctx.Locals(localKeyForTime).(time.Time)
	return now.Sub(start)
}

const (
	localKeyForTime = "$__request-start-time__$"
	localKeyForID   = "$__request-id__$"
)
