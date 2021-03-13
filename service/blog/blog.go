package blog

import (
	"context"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/data/enum/level"
	"github.com/NoHomey/chaos-go-camp-proj/data/enum/rating"
	"github.com/NoHomey/chaos-go-camp-proj/data/tag"
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

//Blog is the expoxed blog data.
type Blog struct {
	ID            string        `json:"id"`
	FeedURL       string        `json:"feedURL"`
	Author        string        `json:"author"`
	Title         string        `json:"title"`
	Description   string        `json:"descrition,omitempty"`
	Rating        rating.Rating `json:"rating"`
	Level         level.Level   `json:"level"`
	Tags          []tag.Tag     `json:"tags"`
	QuickNote     QuickNote     `json:"quickNote"`
	SavedAt       time.Time     `json:"savedAt"`
	StartedAt     *time.Time    `json:"startedAt,omitempty"`
	FinishedAt    *time.Time    `json:"finishedAt,omitempty"`
	LastSyncedAt  *time.Time    `json:"lastSyncedAt,omitempty"`
	LastUpdatedAt *time.Time    `json:"lastUpdatedAt,omitempty"`
}

//QuickNote is the expoxed quick note data.
type QuickNote struct {
	Text   string `json:"text,omitempty"`
	Public bool   `json:"public"`
}

//Service is an abstraction for the Blog service.
type Service interface {
	Save(ctx context.Context, userID uuid.UUID, data *data.Blog) (primitive.ObjectID, ctxerr.Error)
	Fetch(ctx context.Context, userID uuid.UUID, data *data.FetchBlogs) ([]*Blog, ctxerr.Error)
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

func (srvc *service) Save(ctx context.Context, userID uuid.UUID, data *data.Blog) (primitive.ObjectID, ctxerr.Error) {
	err := srvc.validate.Struct(data)
	if err != nil {
		return primitive.ObjectID{}, tmvalerrs.LogAndReturnCtxErr(&tmvalerrs.Ctx{
			Err:    err.(validator.ValidationErrors),
			Logger: srvc.logger,
			Msg:    "Invalid blog save data",
			Log: []zapcore.Field{
				zap.String("userID", userID.String()),
			},
		})
	}
	id, err := srvc.blogRepo.Save(ctx, userID, &repo.BlogData{
		FeedURL:     data.FeedURL,
		Author:      data.Author,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Level:       data.Level,
		Tags:        data.Tags,
		QuickNote:   data.QuickNote,
	})
	if err != nil {
		srvc.logger.Error(
			"Could not save blog",
			zap.Error(err),
		)
		return primitive.ObjectID{}, ctxerr.NewInternal(err)
	}
	srvc.logger.Info(
		"Saved blog",
		zap.String("userID", userID.String()),
		zap.String("feedURL", data.FeedURL),
	)
	return id, nil
}

func (srvc *service) Fetch(ctx context.Context, userID uuid.UUID, data *data.FetchBlogs) ([]*Blog, ctxerr.Error) {
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
		Tags:   data.Tags,
		Count:  data.Count,
		Rating: data.Rating,
		Level:  data.Level,
		After:  after,
	})
	if err != nil {
		srvc.logger.Error(
			"Could not fetch blogs",
			zap.String("userID", userID.String()),
			zap.Strings("tags", tag.Tags(data.Tags)),
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
	return transform(blogs), nil
}

func transform(data []model.Blog) []*Blog {
	res := make([]*Blog, len(data))
	for i := range data {
		res[i] = toData(data[i])
	}
	return res
}

func toData(m model.Blog) *Blog {
	return &Blog{
		ID:          m.ID().Hex(),
		FeedURL:     m.FeedURL(),
		Author:      m.Author(),
		Title:       m.Title(),
		Description: m.Description(),
		Rating:      m.Rating(),
		Level:       m.Level(),
		Tags:        m.Tags(),
		QuickNote: QuickNote{
			Text:   m.QuickNote(),
			Public: m.IsQickNotePublic(),
		},
		SavedAt:       m.SavedAt(),
		StartedAt:     m.StartedAt(),
		FinishedAt:    m.FinishedAt(),
		LastSyncedAt:  m.LastSyncedAt(),
		LastUpdatedAt: m.LastUpdatedAt(),
	}
}
