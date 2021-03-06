package user

import (
	"context"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/env"
	routes "github.com/NoHomey/chaos-go-camp-proj/expose/routes/user"
	miscfx "github.com/NoHomey/chaos-go-camp-proj/misc/fx"
	"github.com/NoHomey/chaos-go-camp-proj/mysql/open"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/data"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/repo"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/access"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/prime"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/service/user"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

//Module bundles fx.Options for the user Fx Module.
var Module = fx.Options(
	fx.Provide(func(lc fx.Lifecycle, logger *zap.Logger, validate *validator.Validate) (access.Service, prime.Service, error) {
		db, err := open.Open(open.DB{
			User: env.Get(usernameKey),
			Pass: env.Get(passwordKey),
			Name: env.Get(dbNameKey),
		})
		if err != nil {
			return nil, nil, err
		}
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return open.Ping(ctx, db)
			},
			OnStop: miscfx.IgnoreContext(db.Close),
		})
		userService := user.Use(
			repo.UserRepoForDB(db),
			logger,
			validate,
		)
		accessService := access.Use(
			repo.AccessRepoForDB(db),
			logger,
			[]byte(env.Get(refreshSecretKey)),
			[]byte(env.Get(accessSecretKey)),
		)
		lc.Append(miscfx.CronJob(func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			accessService.RemExpired(ctx)
		}, time.Minute))
		return accessService, prime.Use(userService, accessService), nil
	}),
	fx.Invoke(data.RegisterPasswordValidator),
	fx.Invoke(routes.Register),
)

const (
	usernameKey      = "USER_DB_USERNAME"
	passwordKey      = "USER_DB_PASSWORD"
	dbNameKey        = "USER_DB_NAME"
	refreshSecretKey = "REFRESH_TOKEN_SECRET"
	accessSecretKey  = "ACCESS_TOKEN_SECRET"
)
