package prime

import (
	"context"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/data"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/model"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/access"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/user"
)

//Service abstracts the main service.
type Service interface {
	SignUp(ctx context.Context, user data.User) ctxerr.Error
	SignIn(ctx context.Context, data data.Auth) (model.User, *access.Token, ctxerr.Error)
	SignOut(ctx context.Context, refresh access.SyncToken) ctxerr.Error
	CleanExpired(ctx context.Context) ctxerr.Error
	ObtainAccess(ctx context.Context, refresh access.SyncToken) (model.Access, *access.SyncToken, ctxerr.Error)
}

//Use returns a Service implementation wich uses the proviced user.Service and access.Service.
func Use(usrSrvc user.Service, acsSrvc access.Service) Service {
	return service{usrSrvc, acsSrvc}
}

type service struct {
	userService   user.Service
	accessService access.Service
}

func (srvc service) SignUp(ctx context.Context, user data.User) ctxerr.Error {
	return srvc.userService.Register(ctx, user)
}

func (srvc service) SignIn(ctx context.Context, data data.Auth) (model.User, *access.Token, ctxerr.Error) {
	authCtx, authCancle := context.WithTimeout(ctx, time.Second)
	defer authCancle()
	user, err := srvc.userService.Authenticate(authCtx, data)
	if err != nil {
		return nil, nil, err
	}
	accessCtx, accessCancel := context.WithTimeout(ctx, 2*time.Second)
	defer accessCancel()
	token, err := srvc.accessService.GrantAccess(accessCtx, user)
	return user, token, err
}

func (srvc service) SignOut(ctx context.Context, refresh access.SyncToken) ctxerr.Error {
	return srvc.accessService.RevokeAccess(ctx, refresh)
}

func (srvc service) CleanExpired(ctx context.Context) ctxerr.Error {
	return srvc.accessService.RemExpired(ctx)
}

func (srvc service) ObtainAccess(ctx context.Context, refresh access.SyncToken) (model.Access, *access.SyncToken, ctxerr.Error) {
	return srvc.accessService.RefreshAccess(ctx, refresh)
}
