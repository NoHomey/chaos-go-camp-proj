package access

import (
	"context"
	"fmt"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/model"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/repo"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

//ErrInvalSyncToken is used for reporting that a sync token is invalid.
type ErrInvalSyncToken struct {
	str string
}

func (err ErrInvalSyncToken) Error() string {
	return fmt.Sprintf("Invalid Sync token: %s", err.str)
}

func (err ErrInvalSyncToken) Unwrap() error {
	return nil
}

//Text returns human readable error text.
func (err ErrInvalSyncToken) Text() string {
	return "Invalid Sync token"
}

//Context returns error Context.
func (err ErrInvalSyncToken) Context() ctxerr.Context {
	return ctxerr.Context{
		Name: "invalid-sync-token",
	}
}

//HttpStatusCode returns http status code for the error.
func (err ErrInvalSyncToken) HttpStatusCode() int {
	return 400
}

//SyncToken is pair of PASETO token and synchronization token.
type SyncToken struct {
	Token string
	Sync  string
}

//Token is pair of access and refresh SyncTokens.
type Token struct {
	Access  SyncToken
	Refresh SyncToken
}

//TokenData is the data obtained from validating and decoding SyncToken.
type TokenData struct {
	TokenID  uuid.UUID
	UserID   uuid.UUID
	IssuedAt time.Time
}

//Service abstracts the access service.
type Service interface {
	GrantAccess(ctx context.Context, user model.User) (*Token, int64, ctxerr.Error)
	RemExpired(ctx context.Context) ctxerr.Error
	RefreshAccess(ctx context.Context, refresh SyncToken) (model.Access, *SyncToken, int64, ctxerr.Error)
	RevokeAccess(ctx context.Context, refresh SyncToken) ctxerr.Error
	DecodeAndValidateAccessToken(access SyncToken) (*TokenData, ctxerr.Error)
}

//Use creates an access Service.
func Use(repo repo.AccessRepo, logger *zap.Logger, refreshSecret []byte, accessSecret []byte) Service {
	return service{repo, logger, refreshSecret, accessSecret}
}

type service struct {
	repo          repo.AccessRepo
	logger        *zap.Logger
	refreshSecret []byte
	accessSecret  []byte
}

func (srvc service) GrantAccess(ctx context.Context, user model.User) (*Token, int64, ctxerr.Error) {
	refreshID := uuid.New()
	refreshToken, now, err := srvc.genSyncToken(&tokenGenData{
		userID:        user.ID(),
		userEmail:     user.Email(),
		tokenID:       refreshID,
		forRefresh:    true,
		tokenDuration: refreshDuration,
	})
	if err != nil {
		return nil, 0, err
	}
	accessID := uuid.New()
	accessToken, _, err := srvc.genSyncToken(&tokenGenData{
		userID:        user.ID(),
		userEmail:     user.Email(),
		tokenID:       accessID,
		forRefresh:    false,
		tokenDuration: accessDuration,
		refreshID:     &refreshID,
	})
	if err != nil {
		return nil, 0, err
	}
	cerr := srvc.repo.Create(ctx, repo.AccessData{
		UserID:   user.ID(),
		AccessID: refreshID,
		Since:    now,
	})
	if cerr != nil {
		srvc.logger.Error(
			"Failed to create Access record",
			zap.String("userEmail", user.Email()),
			zap.String("userID", user.ID().String()),
			zap.String("accessID", refreshID.String()),
		)
		return nil, 0, ctxerr.NewInternal(cerr)
	}
	srvc.logger.Info(
		"Granting access",
		zap.String("userID", user.ID().String()),
		zap.String("refreshTokenID", refreshID.String()),
		zap.Time("accessGratedAt", now),
		zap.String("accessTokenID", accessID.String()),
	)
	return &Token{Refresh: refreshToken, Access: accessToken}, accessDuration.Milliseconds(), nil
}

func (srvc service) RemExpired(ctx context.Context) ctxerr.Error {
	err := srvc.repo.RemExpired(ctx)
	if err != nil {
		srvc.logger.Error("Failed to remove expired Access records")
		return ctxerr.NewInternal(err)
	}
	srvc.logger.Info("Removed expired Access records")
	return nil
}

func (srvc service) RefreshAccess(ctx context.Context, refresh SyncToken) (model.Access, *SyncToken, int64, ctxerr.Error) {
	data, cerr := srvc.decodeRefreshToken(refresh)
	if cerr != nil {
		return nil, nil, 0, cerr
	}
	found, err := srvc.repo.Find(ctx, *data)
	if err != nil {
		srvc.logger.Error(
			"Failed to find Access record",
			zap.String("userID", data.UserID.String()),
			zap.String("accessID", data.AccessID.String()),
			zap.Time("createdAt", data.Since),
		)
		return nil, nil, 0, ctxerr.NewInternal(err)
	}
	tokenID := uuid.New()
	token, _, cerr := srvc.genSyncToken(&tokenGenData{
		userID:        data.UserID,
		userEmail:     found.UserEmail(),
		tokenID:       tokenID,
		forRefresh:    false,
		tokenDuration: accessDuration,
		refreshID:     &data.AccessID,
	})
	if cerr != nil {
		return nil, nil, 0, cerr
	}
	srvc.logger.Info(
		"Refreshing access",
		zap.String("userID", data.UserID.String()),
		zap.String("accessID", data.AccessID.String()),
		zap.Time("createdAt", data.Since),
		zap.String("tokenID", tokenID.String()),
		zap.String("userEmail", found.UserEmail()),
	)
	return found, &token, accessDuration.Milliseconds(), nil
}

func (srvc service) RevokeAccess(ctx context.Context, refresh SyncToken) ctxerr.Error {
	data, cerr := srvc.decodeRefreshToken(refresh)
	if cerr != nil {
		return cerr
	}
	err := srvc.repo.Delete(ctx, *data)
	if err != nil {
		srvc.logger.Error(
			"Failed to delete Access record",
			zap.String("userID", data.UserID.String()),
			zap.String("accessID", data.AccessID.String()),
			zap.Time("createdAt", data.Since),
			zap.Error(err),
		)
		return ctxerr.NewInternal(err)
	}
	srvc.logger.Info(
		"Access revoked",
		zap.String("userID", data.UserID.String()),
		zap.String("accessID", data.AccessID.String()),
		zap.Time("createdAt", data.Since),
	)
	return nil
}

func (srvc service) DecodeAndValidateAccessToken(access SyncToken) (*TokenData, ctxerr.Error) {
	return srvc.decodeToken(&decodeData{
		token:     access,
		isRefresh: false,
	})
}

func (srvc service) decodeRefreshToken(refresh SyncToken) (*repo.AccessData, ctxerr.Error) {
	data, err := srvc.decodeToken(&decodeData{
		token:     refresh,
		isRefresh: true,
	})
	if err != nil {
		return nil, err
	}
	return &repo.AccessData{
		UserID:   data.UserID,
		AccessID: data.TokenID,
		Since:    data.IssuedAt,
	}, nil
}
