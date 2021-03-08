package repo

import (
	"context"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/service/blog/data"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog/enum/level"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog/enum/rating"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repo is an abstraction for the blog repository.
type Repo interface {
	Save(ctx context.Context, data data.Blog) error
}

//UseCollection returns Repo wich uses the given collection.
func UseCollection(coll *mongo.Collection) Repo {
	return repo{coll}
}

type repo struct {
	coll *mongo.Collection
}

func (r repo) Save(ctx context.Context, data data.Blog) error {
	_, err := r.coll.InsertOne(ctx, &blog{
		FeedURLField:     data.FeedURL,
		AuthorField:      data.Author,
		TitleField:       data.Title,
		DescriptionField: data.Description,
		RatingField:      data.Rating,
		LevelField:       data.Level,
		TagsField:        data.Tags,
		QuickNoteObjField: quickNote{
			TextField:   data.QuickNote,
			PublicField: false,
		},
		SavedAtField: time.Now(),
	})
	return err
}

type blog struct {
	IDField               primitive.ObjectID `bson:"_id,omitempty"`
	FeedURLField          string             `bson:"feedURL"`
	WebsiteField          string             `bson:"website"`
	AuthorField           string             `bson:"author"`
	TitleField            string             `bson:"title"`
	DescriptionField      string             `bson:"descrition"`
	RatingField           uint8              `bson:"rating"`
	LevelField            uint8              `bson:"level"`
	TagsField             []string           `bson:"tags"`
	QuickNoteObjField     quickNote          `bson:"quickNote"`
	SavedAtField          time.Time          `bson:"savedAt"`
	StartedAtOptField     *time.Time         `bson:"startedAt"`
	FinishedAtOptField    *time.Time         `bson:"finishedAt"`
	LastSyncedAtOptField  *time.Time         `bson:"lastSyncedAt"`
	LastUpdatedAtOptField *time.Time         `bson:"lastUpdatedAt"`
}

type quickNote struct {
	TextField   string `bson:"text"`
	PublicField bool   `bson:"public"`
}

func (b *blog) ID() primitive.ObjectID {
	return b.IDField
}

func (b *blog) FeedURL() string {
	return b.FeedURLField
}

func (b *blog) Website() string {
	return b.WebsiteField
}

func (b *blog) Author() string {
	return b.AuthorField
}

func (b *blog) Title() string {
	return b.TitleField
}

func (b *blog) Description() string {
	return b.DescriptionField
}

func (b *blog) Rating() rating.Rating {
	return rating.FromNum(b.RatingField)
}

func (b *blog) Level() level.Level {
	return level.FromNum(b.LevelField)
}

func (b *blog) Tags() []string {
	return b.TagsField
}

func (b *blog) QuickNote() string {
	return b.QuickNoteObjField.TextField
}

func (b *blog) IsQickNotePublic() bool {
	return b.QuickNoteObjField.PublicField
}

func (b *blog) SavedAt() time.Time {
	return b.SavedAtField
}

func (b *blog) FinishedAt() *time.Time {
	return b.FinishedAtOptField
}

func (b *blog) LastSyncedAt() *time.Time {
	return b.LastSyncedAtOptField
}

func (b *blog) LastUpdatedAt() *time.Time {
	return b.LastUpdatedAtOptField
}
