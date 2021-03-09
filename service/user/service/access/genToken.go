package access

import (
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/secrand"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"go.uber.org/zap"
)

type tokenGenData struct {
	userID        uuid.UUID
	userEmail     string
	tokenID       uuid.UUID
	forRefresh    bool
	tokenDuration time.Duration
	refreshID     *uuid.UUID
}

func (srvc service) genSyncToken(data *tokenGenData) (SyncToken, time.Time, ctxerr.Error) {
	now := time.Now()
	sync, err := secrand.RandString(syncTokenCount)
	if err != nil {
		srvc.logger.Error(
			"Failed to generate secure random string",
			zap.Error(err),
		)
		return SyncToken{}, now, ctxerr.NewInternal(err)
	}
	jsonToken := initToken(data, now)
	jsonToken.Set(tokenTypeKey, stringTokenType(data.forRefresh))
	jsonToken.Set(syncTokenKey, sync)
	if data.refreshID != nil {
		jsonToken.Set("refresh-token", data.refreshID.String())
	}
	token, err := paseto.NewV2().Encrypt(srvc.obtainSecret(data.forRefresh), jsonToken, nil)
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

func stringTokenType(isRefresh bool) string {
	if isRefresh {
		return "refresh"
	}
	return "access"
}

func (srvc service) obtainSecret(isRefresh bool) []byte {
	if isRefresh {
		return srvc.refreshSecret
	}
	return srvc.accessSecret
}

func initToken(data *tokenGenData, now time.Time) paseto.JSONToken {
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
