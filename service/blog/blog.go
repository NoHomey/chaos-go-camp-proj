package blog

import (
	"context"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog/data"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog/model"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog/repo"
	"github.com/NoHomey/chaos-go-camp-proj/service/tmvalerrs"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Service is an abstraction for the Blog service.
type Service interface {
	Save(ctx context.Context, userID uuid.UUID, data *data.Blog) ctxerr.Error
	Fetch(ctx context.Context, userID uuid.UUID, data *data.FetchBlogs) ([]model.Blog, ctxerr.Error)
}

//Use returns a Service instance wich uses the given arguments.
func Use(blogRepo repo.Repo, logger *zap.Logger, validate *validator.Validate) Service {
	return &service{blogRepo, logger, validate}
}

type service struct {
	blogRepo repo.Repo
	logger   *zap.Logger
	validate *validator.Validate
}

func (srvc *service) Save(ctx context.Context, userID uuid.UUID, data *data.Blog) ctxerr.Error {
	err := srvc.validate.Struct(data)
	if err != nil {
		return tmvalerrs.LogAndReturnCtxErr(&tmvalerrs.Ctx{
			Err:    err.(validator.ValidationErrors),
			Logger: srvc.logger,
			Msg:    "Invalid blog save data",
			Log: []zapcore.Field{
				zap.String("userID", userID.String()),
			},
		})
	}
	err = srvc.blogRepo.Save(ctx, userID, data)
	if err != nil {
		srvc.logger.Error(
			"Could not save blog",
			zap.Error(err),
		)
		return ctxerr.NewInternal(err)
	}
	srvc.logger.Info(
		"Saved blog",
		zap.String("userID", userID.String()),
		zap.String("feedURL", data.FeedURL),
	)
	return nil
}

func (srvc *service) Fetch(ctx context.Context, userID uuid.UUID, data *data.FetchBlogs) ([]model.Blog, ctxerr.Error) {
	err := srvc.validate.Struct(data)
	if err != nil {
		return nil, tmvalerrs.LogAndReturnCtxErr(&tmvalerrs.Ctx{
			Err:    err.(validator.ValidationErrors),
			Logger: srvc.logger,
			Msg:    "Invalid blog fetch data",
			Log: []zapcore.Field{
				zap.String("userID", userID.String()),
			},
		})
	}
	var after *primitive.ObjectID
	if len(data.After) > 0 {
		*after, err = primitive.ObjectIDFromHex(data.After)
		if err != nil {
			srvc.logger.Error(
				"Invalid ObjectID for the After field of FetchBlogs data",
				zap.String("userID", userID.String()),
				zap.String("after", data.After),
			)
			return nil, ctxerr.NewInvalData(err, map[string]string{"after": "objectID"})
		}
	}
	blogs, err := srvc.blogRepo.Fetch(ctx, userID, &repo.FetchData{
		Tags:  data.Tags,
		Count: data.Count,
		After: after,
	})
	if err != nil {
		srvc.logger.Error(
			"Could not fetch blogs",
			zap.String("userID", userID.String()),
			zap.Strings("tags", data.Tags),
			zap.Uint32("count", data.Count),
			zap.Reflect("after", data.After),
			zap.Error(err),
		)
		return nil, ctxerr.NewInternal(err)
	}
	srvc.logger.Info(
		"Fetched bloogs",
		zap.String("userID", userID.String()),
		zap.Int("blogsLen", len(blogs)),
	)
	return blogs, nil
}
