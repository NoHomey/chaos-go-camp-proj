package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/mysql/errors"
	sqluuid "github.com/NoHomey/chaos-go-camp-proj/mysql/types/uuid"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/model"
	"github.com/google/uuid"
)

//UserData is the data needed to create new user.
type UserData struct {
	Name         string
	Email        string
	PasswordHash []byte
}

//UserRepo is an abstraction for user repository.
type UserRepo interface {
	Create(ctx context.Context, data UserData) error
	FindByEmail(ctx context.Context, email string) (model.User, error)
}

//UserRepoForDB returns UserRepo for the given db.
func UserRepoForDB(db *sql.DB) UserRepo {
	return userRepo{db}
}

type userRepo struct {
	db *sql.DB
}

func (repo userRepo) Create(ctx context.Context, data UserData) error {
	sql := "INSERT INTO User(name, email, password) VALUES (?, ?, ?)"
	args := []interface{}{data.Name, data.Email, data.PasswordHash}
	_, err := repo.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, sql, args)
	}
	return nil
}

func (repo userRepo) FindByEmail(ctx context.Context, email string) (model.User, error) {
	sql := "SELECT id, name, email, password, registered_at, last_modified_at FROM User WHERE email = ?"
	args := []interface{}{email}
	var user user
	read := []interface{}{
		&user.id,
		&user.name,
		&user.email,
		&user.passwordHash,
		&user.registeredAt,
		&user.lastModifiedAt,
	}
	err := repo.db.QueryRowContext(ctx, sql, args...).Scan(read...)
	if err != nil {
		return nil, errors.Wrap(err, sql, args)
	}
	return &user, nil
}

//compile time contract check
var _ model.User = &user{}

type user struct {
	id             sqluuid.UUID
	name           string
	email          string
	passwordHash   []byte
	registeredAt   time.Time
	lastModifiedAt time.Time
}

func (u *user) ID() uuid.UUID {
	return u.id.UUID
}

func (u *user) Name() string {
	return u.name
}

func (u *user) Email() string {
	return u.email
}

func (u *user) PasswordHash() []byte {
	return u.passwordHash
}

func (u *user) RegisteredAt() time.Time {
	return u.registeredAt
}

func (u *user) LastModifiedAt() time.Time {
	return u.lastModifiedAt
}
