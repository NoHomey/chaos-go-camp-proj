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
	Since    time.Time
}

//AccessRepo is an abstraction for access repository.
type AccessRepo interface {
	Create(ctx context.Context, data AccessData) error
	Find(ctx context.Context, data AccessData) (model.Access, error)
	Delete(ctx context.Context, data AccessData) error
	RemExpired(ctx context.Context) error
}

//AccessRepoForDB returns AccessRepo for the given db.
func AccessRepoForDB(db *sql.DB) AccessRepo {
	return accessRepo{db}
}

type accessRepo struct {
	db *sql.DB
}

func (repo accessRepo) Create(ctx context.Context, data AccessData) error {
	sql := "INSERT INTO Access(id, userID, created_at) VALUES (?, ?, ?)"
	args := args.Args(sqluuid.Wrap(data.AccessID), sqluuid.Wrap(data.UserID), data.Since.UTC())
	_, err := repo.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, sql, args)
	}
	return nil
}

func (repo accessRepo) Find(ctx context.Context, info AccessData) (model.Access, error) {
	sql := `
	SELECT Access.id, Access.userID, Access.created_at, User.email, User.name
	FROM Access JOIN User ON Access.userID = User.id
	WHERE Access.id = ? AND Access.userID = ?`
	args := args.Args(sqluuid.Wrap(info.AccessID), sqluuid.Wrap(info.UserID))
	var data access
	read := []interface{}{&data.id, &data.userID, &data.createdAt, &data.userEmail, &data.userName}
	err := repo.db.QueryRowContext(ctx, sql, args...).Scan(read...)
	if err != nil {
		return nil, errors.Wrap(err, sql, args)
	}
	return &data, nil
}

func (repo accessRepo) Delete(ctx context.Context, data AccessData) error {
	sql := "DELETE FROM Access WHERE id = ? AND userID = ?"
	args := args.Args(sqluuid.Wrap(data.AccessID), sqluuid.Wrap(data.UserID))
	_, err := repo.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, sql, args)
	}
	return nil
}

func (repo accessRepo) RemExpired(ctx context.Context) error {
	sql := "DELETE FROM Access WHERE DATE_ADD(created_at, INTERVAL 1 WEEK) <= CURRENT_TIME"
	_, err := repo.db.ExecContext(ctx, sql)
	if err != nil {
		return errors.Wrap(err, sql, nil)
	}
	return nil
}

type access struct {
	id        sqluuid.UUID
	userID    sqluuid.UUID
	userEmail string
	userName  string
	createdAt time.Time
}

func (data *access) ID() uuid.UUID {
	return data.id.UUID
}

func (data *access) UserID() uuid.UUID {
	return data.userID.UUID
}

func (data *access) UserEmail() string {
	return data.userEmail
}

func (data *access) UserName() string {
	return data.userName
}

func (data *access) CreatedAt() time.Time {
	return data.createdAt
}
