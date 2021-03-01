package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/mysql/args"
	"github.com/NoHomey/chaos-go-camp-proj/mysql/errors"
	sqluuid "github.com/NoHomey/chaos-go-camp-proj/mysql/types/uuid"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/model"
	"github.com/google/uuid"
)

//AccessData is the data for creating access record.
type AccessData struct {
	UserID   uuid.UUID
	AccessID uuid.UUID
}

//AccessRepo is an abstraction for access repository.
type AccessRepo interface {
	Create(ctx context.Context, data AccessData) error
	FindByID(ctx context.Context, id uuid.UUID) (model.Access, error)
}

//AccessRepoForDB returns AccessRepo for the given db.
func AccessRepoForDB(db *sql.DB) AccessRepo {
	return accessRepo{db}
}

type accessRepo struct {
	db *sql.DB
}

func (repo accessRepo) Create(ctx context.Context, data AccessData) error {
	sql := "INSERT INTO Access(id, userID) VALUES (?, ?)"
	args := args.Args(sqluuid.Wrap(data.UserID), sqluuid.Wrap(data.AccessID))
	_, err := repo.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, sql, args)
	}
	return nil
}

func (repo accessRepo) FindByID(ctx context.Context, id uuid.UUID) (model.Access, error) {
	sql := "SELECT id, userID, created_at FROM Access WHERE id = ?"
	args := args.Args(id)
	var data access
	read := []interface{}{&data.id, &data.userID, &data.createdAt}
	err := repo.db.QueryRowContext(ctx, sql, args...).Scan(read...)
	if err != nil {
		return nil, errors.Wrap(err, sql, args)
	}
	return &data, nil
}

type access struct {
	id        sqluuid.UUID
	userID    sqluuid.UUID
	createdAt time.Time
}

func (data *access) ID() uuid.UUID {
	return data.id.UUID
}

func (data *access) UserID() uuid.UUID {
	return data.userID.UUID
}

func (data *access) CreatedAt() time.Time {
	return data.createdAt
}
