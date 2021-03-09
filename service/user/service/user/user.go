package user

import (
	"context"
	"fmt"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/mysql/errors"
	"github.com/NoHomey/chaos-go-camp-proj/service/tmvalerrs"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/data"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/model"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/repo"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/bcrypt"
)

//Service abstracts the user service.
type Service interface {
	Register(ctx context.Context, user data.User) ctxerr.Error
	Authenticate(ctx context.Context, data data.Auth) (model.User, ctxerr.Error)
}

//Use creates a user Service.
func Use(repo repo.UserRepo, logger *zap.Logger, validate *validator.Validate) Service {
	return service{repo, logger, validate}
}

//ErrEmailTaken signals that user exists for the given email address.
type ErrEmailTaken struct {
	email   string
	wrapped error
}

func (err ErrEmailTaken) Error() string {
	return fmt.Sprintf("User already exists with email: %s", err.email)
}

func (err ErrEmailTaken) Unwrap() error {
	return err.wrapped
}

//Text returns human readable error text.
func (err ErrEmailTaken) Text() string {
	return fmt.Sprintf("Email address %s is taken by a user", err.email)
}

//Context returns error Context.
func (err ErrEmailTaken) Context() ctxerr.Context {
	return ctxerr.Context{
		Name: "email-taken",
		Data: err.email,
	}
}

//HttpStatusCode returns http status code for the error.
func (err ErrEmailTaken) HttpStatusCode() int {
	return 400
}

//ErrInvalCredents signals that invalid authentication credentials were given.
type ErrInvalCredents struct {
	wrapped error
}

func (err ErrInvalCredents) Error() string {
	return "Invalid authentication credentials"
}

func (err ErrInvalCredents) Unwrap() error {
	return err.wrapped
}

//Text returns human readable error text.
func (err ErrInvalCredents) Text() string {
	return "Invalid credentials"
}

//Context returns error Context.
func (err ErrInvalCredents) Context() ctxerr.Context {
	return ctxerr.Context{
		Name: "invalid-credentials",
	}
}

//HttpStatusCode returns http status code for the error.
func (err ErrInvalCredents) HttpStatusCode() int {
	return 400
}

type service struct {
	repo     repo.UserRepo
	logger   *zap.Logger
	validate *validator.Validate
}

func (srvc service) Register(ctx context.Context, user data.User) ctxerr.Error {
	err := srvc.validate.Struct(user)
	if err != nil {
		return tmvalerrs.LogAndReturnCtxErr(&tmvalerrs.Ctx{
			Err:    err.(validator.ValidationErrors),
			Logger: srvc.logger,
			Msg:    "Invalid registration data",
			Log: []zapcore.Field{
				zap.String("name", user.Name),
				zap.String("email", user.Email),
			},
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		srvc.logger.Error(
			"Could not hash password",
			zap.String("password", user.Password),
			zap.Error(err),
		)
		return ctxerr.NewInternal(err)
	}

	err = srvc.repo.Create(ctx, repo.UserData{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: hash,
	})
	if err != nil {
		srvc.logger.Error(
			"Could not create user",
			zap.String("name", user.Name),
			zap.String("email", user.Email),
			zap.Error(err),
		)
		if _, ok := err.(errors.ErrExists); ok {
			return ErrEmailTaken{user.Email, err}
		}
		return ctxerr.NewInternal(err)
	}

	srvc.logger.Info(
		"Succesful user registration",
		zap.String("name", user.Name),
		zap.String("email", user.Email),
	)
	return nil
}

func (srvc service) Authenticate(ctx context.Context, data data.Auth) (model.User, ctxerr.Error) {
	err := srvc.validate.Struct(data)
	if err != nil {
		return nil, tmvalerrs.LogAndReturnCtxErr(&tmvalerrs.Ctx{
			Err:    err.(validator.ValidationErrors),
			Logger: srvc.logger,
			Msg:    "Invalid authentication data",
			Log: []zapcore.Field{
				zap.String("email", data.Email),
			},
		})
	}

	user, err := srvc.repo.FindByEmail(ctx, data.Email)
	if err != nil {
		srvc.logger.Error(
			"Could not find user",
			zap.String("email", data.Email),
			zap.Error(err),
		)
		if _, ok := err.(errors.ErrNotFound); ok {
			return nil, ErrInvalCredents{err}
		}
		return nil, ctxerr.NewInternal(err)
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash(), []byte(data.Password))
	if err != nil {
		srvc.logger.Error(
			"Password mismatch",
			zap.String("email", data.Email),
			zap.Error(err),
		)
		return nil, ErrInvalCredents{err}
	}

	srvc.logger.Info(
		"Succesful user authentication",
		zap.String("email", data.Email),
	)
	return user, nil
}
