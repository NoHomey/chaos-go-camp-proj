package blog

import (
	"context"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/expose/middleware/auth"
	"github.com/NoHomey/chaos-go-camp-proj/expose/reqlogger"
	"github.com/NoHomey/chaos-go-camp-proj/expose/sendresult"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog/data"
	"github.com/gofiber/fiber/v2"
)

//Handler implements blogs handler.
type Handler struct {
	Service   blog.Service
	ReqLogger reqlogger.Logger
}

//Save saves a blog.
func (h Handler) Save(ctx *fiber.Ctx) error {
	blog := new(data.Blog)
	if err := ctx.BodyParser(blog); err != nil {
		return ctxerr.NewBadFormat(err)
	}
	accessData := auth.AccessData(ctx)
	srvcCtx, cancel := context.WithTimeout(ctx.Context(), 2*time.Second)
	defer cancel()
	err := h.Service.Save(srvcCtx, accessData.UserID, blog)
	if err != nil {
		return err
	}
	return sendresult.SendAndLog(ctx, fiber.StatusCreated, true, h.ReqLogger)
}

//Fetch fetches blogs.
func (h Handler) Fetch(ctx *fiber.Ctx) error {
	data := new(data.FetchBlogs)
	if err := ctx.BodyParser(data); err != nil {
		return ctxerr.NewBadFormat(err)
	}
	accessData := auth.AccessData(ctx)
	srvcCtx, cancel := context.WithTimeout(ctx.Context(), 3*time.Second)
	defer cancel()
	blogs, err := h.Service.Fetch(srvcCtx, accessData.UserID, data)
	if err != nil {
		return err
	}
	return sendresult.SendAndLog(ctx, fiber.StatusOK, blogs, h.ReqLogger)
}
