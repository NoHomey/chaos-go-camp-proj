package app

import (
	"os"

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
	fx.Provide(func(lc fx.Lifecycle) *fiber.App {
		app := fiber.New()
		lc.Append(fx.Hook{
			OnStart: miscfx.IgnoreContext(func() error {
				go app.Listen(":" + os.Getenv(portKey))
				return nil
			}),
			OnStop: miscfx.IgnoreContext(app.Shutdown),
		})
		return app
	}),
	fx.Provide(func(lc fx.Lifecycle) reqlogger.Logger {
		path := os.Getenv(reqLogPathKey)
		logger := reqlogger.New(logcrtr.Config(path))
		lc.Append(fx.Hook{
			OnStop: miscfx.IgnoreContext(logger.Sync),
		})
		return logger
	}),
	fx.Provide(func(lc fx.Lifecycle) *zap.Logger {
		path := os.Getenv(appLogPathKey)
		logger := logcrtr.Config(path)
		lc.Append(fx.Hook{
			OnStop: miscfx.IgnoreContext(logger.Sync),
		})
		return logger
	}),
	fx.Provide(validator.New),
	fx.Invoke(logger.Register),
)

const (
	reqLogPathKey = "REQ_LOG_PATH"
	appLogPathKey = "APP_LOG_PATH"
	portKey       = "PORT"
)
