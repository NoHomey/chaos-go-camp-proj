package feed

import (
	"context"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/expose/reqlogger"
	"github.com/NoHomey/chaos-go-camp-proj/expose/sendresult"
	"github.com/NoHomey/chaos-go-camp-proj/misc/base64url"
	"github.com/NoHomey/chaos-go-camp-proj/service/feed"
	"github.com/gofiber/fiber/v2"
)

//Param is the path param.
const Param = "encurl"

//Handler implements getting feed details request handler.
type Handler struct {
	Service   feed.Service
	ReqLogger reqlogger.Logger
}

//Details returns details for a feed url.
func (h Handler) Details(ctx *fiber.Ctx) error {
	encodedURL := ctx.Params("encurl")
	url, err := base64url.DecodeString(encodedURL)
	if err != nil {
		return ctxerr.NewBadFormat(err)
	}
	srvcCtx, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()
	details, cerr := h.Service.Details(srvcCtx, url)
	if cerr != nil {
		return cerr
	}
	return sendresult.SendAndLog(ctx, fiber.StatusOK, details, h.ReqLogger)
}
