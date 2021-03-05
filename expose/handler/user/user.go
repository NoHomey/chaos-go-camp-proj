package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/expose/cookie"
	"github.com/NoHomey/chaos-go-camp-proj/expose/sendresult"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/data"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/access"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/prime"
	"github.com/gofiber/fiber/v2"
)

//SignUp signs user up.
func SignUp(service prime.Service) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user := new(data.User)
		if err := ctx.BodyParser(user); err != nil {
			return ctxerr.NewBadFormat(err)
		}
		srvcCtx, cancel := context.WithTimeout(ctx.Context(), 2*time.Second)
		defer cancel()
		err := service.SignUp(srvcCtx, *user)
		if err != nil {
			return err
		}
		return sendresult.SendRes(ctx, fiber.StatusCreated, true)
	}
}

//SignIn signs user in.
func SignIn(service prime.Service) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		auth := new(data.Auth)
		if err := ctx.BodyParser(auth); err != nil {
			return ctxerr.NewBadFormat(err)
		}
		srvcCtx, cancel := context.WithTimeout(ctx.Context(), 3*time.Second)
		defer cancel()
		user, token, err := service.SignIn(srvcCtx, *auth)
		if err != nil {
			return err
		}
		cookie.Set(ctx, refreshSyncTokenKey, token.Refresh.Sync)
		cookie.Set(ctx, accessTokenKey, token.Access.Token)
		return sendresult.SendRes(ctx, fiber.StatusOK, accessRes{
			Name:            user.Name(),
			Email:           user.Email(),
			RefreshToken:    token.Refresh.Token,
			AccessSyncToken: token.Access.Sync,
		})
	}
}

//SignOut signs user out.
func SignOut(service prime.Service) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		token, err := extractRefreshToken(ctx)
		if err != nil {
			return err
		}
		srvcCtx, cancel := context.WithTimeout(ctx.Context(), 2*time.Second)
		defer cancel()
		err = service.SignOut(srvcCtx, *token)
		if err != nil {
			return err
		}
		return sendresult.SendRes(ctx, fiber.StatusOK, true)
	}
}

//ObtainAccess sends new access token.
func ObtainAccess(service prime.Service) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		token, err := extractRefreshToken(ctx)
		if err != nil {
			return err
		}
		srvcCtx, cancel := context.WithTimeout(ctx.Context(), 2*time.Second)
		defer cancel()
		access, accessToken, err := service.ObtainAccess(srvcCtx, *token)
		if err != nil {
			return err
		}
		cookie.Set(ctx, accessTokenKey, accessToken.Token)
		return sendresult.SendRes(ctx, fiber.StatusOK, accessRes{
			Name:            access.UserName(),
			Email:           access.UserEmail(),
			AccessSyncToken: accessToken.Sync,
		})
	}
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

const (
	refreshSyncTokenKey = "refresh-sync-token"
	accessTokenKey      = "access-token"
)

const (
	authHeaderScheme = "PASETO"
	authHeader       = "Authorization"
)
