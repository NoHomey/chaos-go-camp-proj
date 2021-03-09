package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/expose/cookie"
	"github.com/NoHomey/chaos-go-camp-proj/expose/middleware/auth"
	"github.com/NoHomey/chaos-go-camp-proj/expose/reqlogger"
	"github.com/NoHomey/chaos-go-camp-proj/expose/sendresult"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/data"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/access"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/prime"
	"github.com/gofiber/fiber/v2"
)

//Handler implements request handlers for the user.
type Handler struct {
	Service   prime.Service
	ReqLogger reqlogger.Logger
}

//SignUp signs user up.
func (h Handler) SignUp(ctx *fiber.Ctx) error {
	user := new(data.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctxerr.NewBadFormat(err)
	}
	srvcCtx, cancel := context.WithTimeout(ctx.Context(), 2*time.Second)
	defer cancel()
	err := h.Service.SignUp(srvcCtx, *user)
	if err != nil {
		return err
	}
	return sendresult.SendAndLog(ctx, fiber.StatusCreated, true, h.ReqLogger)
}

//SignIn signs user in.
func (h Handler) SignIn(ctx *fiber.Ctx) error {
	authData := new(data.Auth)
	if err := ctx.BodyParser(authData); err != nil {
		return ctxerr.NewBadFormat(err)
	}
	srvcCtx, cancel := context.WithTimeout(ctx.Context(), 3*time.Second)
	defer cancel()
	user, token, err := h.Service.SignIn(srvcCtx, *authData)
	if err != nil {
		return err
	}
	cookie.Set(ctx, refreshSyncTokenKey, token.Refresh.Sync)
	cookie.Set(ctx, auth.AccessTokenKey, token.Access.Token)
	return sendresult.SendAndLog(ctx, fiber.StatusOK, accessRes{
		Name:            user.Name(),
		Email:           user.Email(),
		RefreshToken:    token.Refresh.Token,
		AccessSyncToken: token.Access.Sync,
	}, h.ReqLogger)
}

//SignOut signs user out.
func (h Handler) SignOut(ctx *fiber.Ctx) error {
	token, err := extractRefreshToken(ctx)
	if err != nil {
		return err
	}
	srvcCtx, cancel := context.WithTimeout(ctx.Context(), 2*time.Second)
	defer cancel()
	err = h.Service.SignOut(srvcCtx, *token)
	if err != nil {
		return err
	}
	return sendresult.SendAndLog(ctx, fiber.StatusOK, true, h.ReqLogger)
}

//ObtainAccess sends new access token.
func (h Handler) ObtainAccess(ctx *fiber.Ctx) error {
	token, err := extractRefreshToken(ctx)
	if err != nil {
		return err
	}
	srvcCtx, cancel := context.WithTimeout(ctx.Context(), 2*time.Second)
	defer cancel()
	access, accessToken, err := h.Service.ObtainAccess(srvcCtx, *token)
	if err != nil {
		return err
	}
	cookie.Set(ctx, auth.AccessTokenKey, accessToken.Token)
	return sendresult.SendAndLog(ctx, fiber.StatusOK, accessRes{
		Name:            access.UserName(),
		Email:           access.UserEmail(),
		AccessSyncToken: accessToken.Sync,
	}, h.ReqLogger)
}

func extractRefreshToken(ctx *fiber.Ctx) (*access.SyncToken, ctxerr.Error) {
	header := ctx.Get(authHeader)
	l := len(authHeaderScheme)
	if len(header) > l+1 && header[:l] == authHeaderScheme {
		return &access.SyncToken{
			Token: header[l+1:],
			Sync:  ctx.Cookies(refreshSyncTokenKey),
		}, nil
	}
	return nil, errInvalAuthHeader{header}
}

type errInvalAuthHeader struct {
	header string
}

func (err errInvalAuthHeader) Error() string {
	return fmt.Sprintf("Invalid Auth header: %s", err.header)
}

func (err errInvalAuthHeader) Unwrap() error {
	return nil
}

//Text returns human readable error text.
func (err errInvalAuthHeader) Text() string {
	return "Invalid Auth header"
}

//Context returns error Context.
func (err errInvalAuthHeader) Context() ctxerr.Context {
	return ctxerr.Context{
		Name: "invalid-auth-header",
	}
}

//HttpStatusCode returns http status code for the error.
func (err errInvalAuthHeader) HttpStatusCode() int {
	return fiber.StatusBadRequest
}

type accessRes struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	RefreshToken    string `json:"refreshToken,omitempty"`
	AccessSyncToken string `json:"accessSyncToken"`
}

const refreshSyncTokenKey = "refresh-sync-token"

const (
	authHeaderScheme = "PASETO"
	authHeader       = "Authorization"
)
