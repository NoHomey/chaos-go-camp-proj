package repo

import (
	"context"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/data/enum/level"
	"github.com/NoHomey/chaos-go-camp-proj/data/enum/priority"
	"github.com/NoHomey/chaos-go-camp-proj/data/enum/rating"
	"github.com/NoHomey/chaos-go-camp-proj/data/tag"
	"github.com/NoHomey/chaos-go-camp-proj/service/blog/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Repo is an abstraction for the blog repository.
type Repo interface {
	Save(ctx context.Context, userID uuid.UUID, data *BlogData) (primitive.ObjectID, error)
	Fetch(ctx context.Context, userID uuid.UUID, data *FetchData) ([]model.Blog, error)
}

//UseCollection returns Repo wich uses the given collection.
func UseCollection(coll *mongo.Collection) Repo {
	return repo{coll}
}

//FetchData represents data for fetching.
type FetchData struct {
	Tags  []tag.Tag
	Count uint32
	After *primitive.ObjectID
}

//BlogData is the data for saving blogs.
type BlogData struct {
	FeedURL     string
	Author      string
	Title       string
	Description string
	Rating      rating.Rating
	Level       level.Level
	Tags        []tag.Tag
	QuickNote   string
}

type repo struct {
	coll *mongo.Collection
}

func (r repo) Save(ctx context.Context, userID uuid.UUID, data *BlogData) (primitive.ObjectID, error) {
	res, err := r.coll.InsertOne(ctx, &blog{
		UserIDHiddenField: userID[:],
		FeedURLField:      data.FeedURL,
		AuthorField:       data.Author,
		TitleField:        data.Title,
		DescriptionField:  data.Description,
		RatingField:       data.Rating.Ord(),
		LevelField:        data.Level.Ord(),
		TagsField:         fromTags(data.Tags),
		QuickNoteObjField: quickNote{
			TextField:   data.QuickNote,
			PublicField: false,
		},
		SavedAtField: time.Now(),
	})
	return res.InsertedID.(primitive.ObjectID), err
}

func (r repo) Fetch(ctx context.Context, userID uuid.UUID, data *FetchData) ([]model.Blog, error) {
	cursor, err := r.findPaged(ctx, userID, data)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	return decodeLimited(ctx, data.Count, cursor)
}

func (r repo) findPaged(ctx context.Context, userID uuid.UUID, data *FetchData) (*mongo.Cursor, error) {
	opts := options.Find()
	opts.SetSort(bson.D{
		{Key: "level", Value: 1},
		{Key: "rating", Value: -1},
		{Key: "_id", Value: 1}})
	opts.SetLimit(int64(data.Count))
	filter := bson.D{
		{Key: "userID", Value: userID[:]},
		{Key: "tags", Value: bson.M{"$all": data.Tags}},
	}
	if data.After != nil {
		filter = bson.D{
			filter[0],
			{Key: "_id", Value: bson.M{"$gt": data.After}},
			filter[1],
		}
	}
	return r.coll.Find(ctx, filter, opts)
}

func decodeLimited(ctx context.Context, limit uint32, cursor *mongo.Cursor) ([]model.Blog, error) {
	blogs := make([]model.Blog, limit)
	i := 0
	for cursor.Next(ctx) {
		b := new(blog)
		cursor.Decode(b)
		blogs[i] = b
		i++
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return blogs[:i], nil
}

type blog struct {
	IDField               primitive.ObjectID `bson:"_id,omitempty"`
	UserIDHiddenField     []byte             `bson:"userID"`
	FeedURLField          string             `bson:"feedURL"`
	AuthorField           string             `bson:"author"`
	TitleField            string             `bson:"title"`
	DescriptionField      string             `bson:"descrition,omitempty"`
	RatingField           uint8              `bson:"rating"`
	LevelField            uint8              `bson:"level"`
	TagsField             []bsonTag          `bson:"tags"`
	QuickNoteObjField     quickNote          `bson:"quickNote"`
	SavedAtField          time.Time          `bson:"savedAt,omitempty"`
	StartedAtOptField     *time.Time         `bson:"startedAt,omitempty"`
	FinishedAtOptField    *time.Time         `bson:"finishedAt,omitempty"`
	LastSyncedAtOptField  *time.Time         `bson:"lastSyncedAt,omitempty"`
	LastUpdatedAtOptField *time.Time         `bson:"lastUpdatedAt,omitempty"`
}

type quickNote struct {
	TextField   string `bson:"text,omitempty"`
	PublicField bool   `bson:"public"`
}

type bsonTag struct {
	Value    string `bson:"value"`
	Priority uint8  `bson:"priority"`
}

func toTags(list []bsonTag) []tag.Tag {
	tags := make([]tag.Tag, len(list))
	for i := range list {
		tags[i] = tag.Tag{
			Value:    list[i].Value,
			Priority: priority.FromNum((list[i].Priority)),
		}
	}
	return tags
}

func fromTags(tags []tag.Tag) []bsonTag {
	list := make([]bsonTag, len(tags))
	for i := range list {
		list[i] = bsonTag{
			Value:    tags[i].Value,
			Priority: tags[i].Priority.Ord(),
		}
	}
	return list
}

func (b *blog) ID() primitive.ObjectID {
	return b.IDField
}

func (b *blog) FeedURL() string {
	return b.FeedURLField
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

func (b *blog) Tags() []tag.Tag {
	return toTags(b.TagsField)
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

func (b *blog) StartedAt() *time.Time {
	return b.StartedAtOptField
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
