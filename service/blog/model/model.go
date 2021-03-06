package model

import (
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/data/enum/level"
	"github.com/NoHomey/chaos-go-camp-proj/data/enum/rating"
	"github.com/NoHomey/chaos-go-camp-proj/data/tag"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Blog abstracts blog entry.
type Blog interface {
	ID() primitive.ObjectID
	FeedURL() string
	Author() string
	Title() string
	Description() string
	Rating() rating.Rating
	Level() level.Level
	Tags() []tag.Tag
	QuickNote() string
	IsQickNotePublic() bool
	SavedAt() time.Time
	StartedAt() *time.Time
	FinishedAt() *time.Time
	LastSyncedAt() *time.Time
	LastUpdatedAt() *time.Time
}
