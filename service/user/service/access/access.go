package access

import (
	"context"
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/secrand"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/model"
	"github.com/NoHomey/chaos-go-camp-proj/service/user/repo"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"go.uber.org/zap"
)

type SyncToken struct {
	Token string
	Sync  string
}

type Token struct {
	Access  SyncToken
	Refresh SyncToken
}

type Service interface {
	GrantAccess(ctx context.Context, user model.User) (*Token, ctxerr.Error)
	RemExpired(ctx context.Context) ctxerr.Error
	RefreshAccess(ctx context.Context, refresh SyncToken) (*SyncToken, model.Access, ctxerr.Error)
	RevokeAccess(ctx context.Context, refresh SyncToken) ctxerr.Error
}

func Use(repo repo.AccessRepo, logger *zap.Logger, refreshSecret []byte, accessSecret []byte) Service {
	return service{repo, logger, refreshSecret, accessSecret}
}

type service struct {
	repo          repo.AccessRepo
	logger        *zap.Logger
	refreshSecret []byte
	accessSecret  []byte
}

func (srvc service) GrantAccess(ctx context.Context, user model.User) (*Token, ctxerr.Error) {
	refreshID := uuid.New()
	refreshToken, now, err := srvc.genSyncToken(tokenData{
		userID:        user.ID(),
		userEmail:     user.Email(),
		tokenID:       refreshID,
		tokenType:     "refresh",
		tokenDuration: refreshDuration,
	})
	if err != nil {
		return nil, err
	}
	accessID := uuid.New()
	accessToken, _, err := srvc.genSyncToken(tokenData{
		userID:        user.ID(),
		userEmail:     user.Email(),
		tokenID:       accessID,
		tokenType:     "access",
		tokenDuration: accessDuration,
		refreshID:     &refreshID,
	})
	if err != nil {
		return nil, err
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
		return nil, ctxerr.NewInternal(cerr)
	}
	srvc.logger.Info(
		"Grating access",
		zap.String("userID", user.ID().String()),
		zap.String("refreshTokenID", refreshID.String()),
		zap.Time("accessGratedAt", now),
		zap.String("accessTokenID", accessID.String()),
	)
	return &Token{Refresh: refreshToken, Access: accessToken}, nil
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

func (srvc service) RefreshAccess(ctx context.Context, refresh SyncToken) (*SyncToken, model.Access, ctxerr.Error) {
	data, cerr := srvc.decodeRefreshToken(refresh)
	if cerr != nil {
		return nil, nil, cerr
	}
	found, err := srvc.repo.Find(ctx, *data)
	if err != nil {
		srvc.logger.Error(
			"Failed to find Access record",
			zap.String("userID", data.UserID.String()),
			zap.String("accessID", data.AccessID.String()),
			zap.Time("createdAt", data.Since),
		)
		return nil, nil, ctxerr.NewInternal(err)
	}
	tokenID := uuid.New()
	token, _, cerr := srvc.genSyncToken(tokenData{
		userID:        data.UserID,
		userEmail:     found.UserEmail(),
		tokenID:       tokenID,
		tokenType:     "access",
		tokenDuration: accessDuration,
		refreshID:     &data.AccessID,
	})
	if cerr != nil {
		return nil, nil, cerr
	}
	srvc.logger.Info(
		"Refreshing access",
		zap.String("userID", data.UserID.String()),
		zap.String("accessID", data.AccessID.String()),
		zap.Time("createdAt", data.Since),
		zap.String("tokenID", tokenID.String()),
		zap.String("userEmail", found.UserEmail()),
	)
	return &token, found, nil
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

func (srvc service) decodeRefreshToken(refresh SyncToken) (*repo.AccessData, ctxerr.Error) {
	var jsonToken paseto.JSONToken
	err := paseto.NewV2().Decrypt(refresh.Token, srvc.refreshSecret, &jsonToken, nil)
	if err != nil {
		srvc.logger.Error(
			"Failed to Decrypt refresh token",
			zap.String("token", refresh.Token),
		)
		return nil, ctxerr.NewInternal(err)
	}
	tokenType, syncToken := jsonToken.Get(tokenTypeKey), jsonToken.Get(syncTokenKey)
	if tokenType != "refresh" || syncToken != refresh.Sync || jsonToken.Expiration.Before(time.Now()) {
		srvc.logger.Error(
			"Failed to recognize refresh token",
			zap.String("tokenAudience", jsonToken.Audience),
			zap.String("JTI", jsonToken.Jti),
			zap.String("Subject", jsonToken.Subject),
			zap.Time("expires", jsonToken.Expiration),
			zap.Time("issuedAt", jsonToken.IssuedAt),
			zap.String("tokenType", tokenType),
			zap.String("syncToken", syncToken),
		)
		return nil, ctxerr.NewNotAuthed(nil)
	}
	accessID, err := uuid.Parse(jsonToken.Jti)
	if err != nil {
		srvc.logger.Error(
			"Failed to parse UUID string",
			zap.String("uudi str", jsonToken.Jti),
		)
		return nil, ctxerr.NewInternal(err)
	}
	userID, err := uuid.Parse(jsonToken.Subject)
	if err != nil {
		srvc.logger.Error(
			"Failed to parse UUID string",
			zap.String("uudi str", jsonToken.Subject),
		)
		return nil, ctxerr.NewInternal(err)
	}
	return &repo.AccessData{
		AccessID: accessID,
		UserID:   userID,
		Since:    jsonToken.IssuedAt,
	}, nil
}

type tokenData struct {
	userID        uuid.UUID
	userEmail     string
	tokenID       uuid.UUID
	tokenType     string
	tokenDuration time.Duration
	refreshID     *uuid.UUID
}

func (srvc service) genSyncToken(data tokenData) (SyncToken, time.Time, ctxerr.Error) {
	now := time.Now()
	sync, err := secrand.RandString(32)
	if err != nil {
		srvc.logger.Error(
			"Failed to generate secure random string",
			zap.Error(err),
		)
		return SyncToken{}, now, ctxerr.NewInternal(err)
	}
	jsonToken := initToken(data, now)
	jsonToken.Set(tokenTypeKey, data.tokenType)
	jsonToken.Set(syncTokenKey, sync)
	if data.refreshID != nil {
		jsonToken.Set("refresh-token", data.refreshID.String())
	}
	secret := srvc.accessSecret
	if data.tokenType == "refresh" {
		secret = srvc.refreshSecret
	}
	token, err := paseto.NewV2().Encrypt(secret, jsonToken, nil)
	if err != nil {
		srvc.logger.Error(
			"Failed to generate PASETO token",
			zap.String("userEmail", data.userEmail),
			zap.String("userID", data.userID.String()),
			zap.Error(err),
		)
		return SyncToken{}, now, ctxerr.NewInternal(err)
	}
	return SyncToken{Token: token, Sync: sync}, now, nil
}

func initToken(data tokenData, now time.Time) paseto.JSONToken {
	return paseto.JSONToken{
		Audience:   "User: " + data.userEmail,
		Issuer:     "AuthService",
		Jti:        data.tokenID.String(),
		Subject:    data.userID.String(),
		Expiration: now.Add(data.tokenDuration),
		IssuedAt:   now,
		NotBefore:  now,
	}
}

const (
	accessDuration  = 5 * time.Minute
	refreshDuration = 7 * 24 * time.Hour
)

const (
	tokenTypeKey = "token-type"
	syncTokenKey = "sync-token"
)
