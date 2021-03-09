package blog

import (
	"context"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/env"
	routes "github.com/NoHomey/chaos-go-camp-proj/expose/routes/blog"
	service "github.com/NoHomey/chaos-go-camp-proj/service/blog"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog/data"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog/repo"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

//Module bundles fx.Options for the feed Fx Module.
var Module = fx.Options(
	fx.Provide(func(lc fx.Lifecycle, logger *zap.Logger, validate *validator.Validate) (service.Service, error) {
		opts := options.Client().ApplyURI(env.Get(dbURIKey))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, opts)
		if err != nil {
			return nil, err
		}
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return client.Ping(ctx, nil)
			},
			OnStop: client.Disconnect,
		})
		coll := client.Database(env.Get(dbNameKey)).Collection(env.Get(collectionKey))
		return service.Use(repo.UseCollection(coll), logger, validate), nil
	}),
	fx.Invoke(data.RegisterValidators),
	fx.Invoke(routes.Register),
)

const (
	dbURIKey      = "BLOG_DB_URI"
	dbNameKey     = "BLOG_DB_NAME"
	collectionKey = "BLOG_DB_COLLECTION"
)
