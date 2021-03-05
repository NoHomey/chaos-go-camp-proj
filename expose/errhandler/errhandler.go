package errhandler

import (
	"fmt"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

//Error is the type of the error data.
type Error struct {
	Message string      `json:"message"`
	Name    string      `json:"name"`
	Data    interface{} `json:"data,omitempty"`
}

//Handler is the error handler creator.
func Handler(logger *zap.Logger) func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusBadRequest
		var send Error
		if e, ok := err.(ctxerr.Error); ok {
			if coder, cOk := e.(interface{ HttpStatusCode() int }); cOk {
				code = coder.HttpStatusCode()
			} else {
				code = statusCodeFromCtxError(e)
			}
			errCtx := e.Context()
			send.Message = e.Text()
			send.Name = errCtx.Name
			send.Data = errCtx.Data
		} else {
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				send.Message = e.Message
				send.Name = "fiber"
			} else {
				send.Message = "Unknown error"
				send.Name = "unknown"
			}
		}
		logger.Error(
			"Sending error",
			zap.String("name", send.Name),
			zap.String("message", send.Message),
			zap.Reflect("data", send.Data),
			zap.Reflect("status", code),
		)
		return ctx.Status(code).JSON(fiber.Map{"error": send})
	}
}

func statusCodeFromCtxError(err ctxerr.Error) int {
	switch err.(type) {
	case ctxerr.InvalData:
		return fiber.StatusBadRequest
	case ctxerr.NotAuthed:
		return fiber.StatusUnauthorized
	case ctxerr.Internal:
		return fiber.StatusInternalServerError
	case ctxerr.BadFormat:
		return fiber.StatusBadRequest
	default:
		panic(fmt.Errorf("Error %T should implement HttpStatusCode method", err))
	}
}
