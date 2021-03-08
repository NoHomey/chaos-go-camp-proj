package app

import (
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/env"
	"github.com/NoHomey/chaos-go-camp-proj/expose/errhandler"
	"github.com/NoHomey/chaos-go-camp-proj/expose/middleware/logger"
	"github.com/NoHomey/chaos-go-camp-proj/expose/reqlogger"
	"github.com/NoHomey/chaos-go-camp-proj/logcrtr"
	miscfx "github.com/NoHomey/chaos-go-camp-proj/misc/fx"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

//Module bundles fx.Options for the app Fx Module.
var Module = fx.Options(
	fx.Provide(func(lc fx.Lifecycle) reqlogger.Logger {
		path := env.Get(reqLogPathKey)
		logger := reqlogger.New(logcrtr.Config(path))
		lc.Append(fx.Hook{
			OnStop: miscfx.IgnoreContext(logger.Sync),
		})
		return logger
	}),
	fx.Provide(func(lc fx.Lifecycle) *zap.Logger {
		path := env.Get(appLogPathKey)
		logger := logcrtr.Config(path)
		lc.Append(fx.Hook{
			OnStop: miscfx.IgnoreContext(logger.Sync),
		})
		return logger
	}),
	fx.Provide(validator.New),
	fx.Provide(func(lc fx.Lifecycle, logger reqlogger.Logger) *fiber.App {
		app := fiber.New(fiber.Config{
			//CaseSensitive: true,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 3 * time.Second,
			ErrorHandler: errhandler.Handler(logger),
		})
		lc.Append(fx.Hook{
			OnStart: miscfx.IgnoreContext(func() error {
				go app.Listen(":" + env.Get(portKey))
				return nil
			}),
			OnStop: miscfx.IgnoreContext(app.Shutdown),
		})
		return app
	}),
	fx.Invoke(logger.Register),
)

const (
	reqLogPathKey = "REQ_LOG_PATH"
	appLogPathKey = "APP_LOG_PATH"
	portKey       = "PORT"
)
