package access

import (
	"time"

	"github.com/NoHomey/chaos-go-camp-proj/ctxerr"
	"github.com/NoHomey/chaos-go-camp-proj/misc/base64url"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"go.uber.org/zap"
)

type decodeData struct {
	token     SyncToken
	isRefresh bool
}

func (srvc service) decodeToken(data *decodeData) (*TokenData, ctxerr.Error) {
	if !isValidSyncToken(data.token.Sync) {
		srvc.logger.Error(
			"Got invalid sync token",
			zap.String("syncToken", data.token.Sync),
		)
		return nil, ErrInvalSyncToken{data.token.Sync}
	}
	var jsonToken paseto.JSONToken
	err := paseto.NewV2().Decrypt(data.token.Token, srvc.obtainSecret(data.isRefresh), &jsonToken, nil)
	if err != nil {
		srvc.logger.Error(
			"Failed to Decrypt refresh token",
			zap.String("token", data.token.Token),
		)
		return nil, ctxerr.NewInternal(err)
	}
	tokenType, syncToken := jsonToken.Get(tokenTypeKey), jsonToken.Get(syncTokenKey)
	if tokenType != stringTokenType(data.isRefresh) || syncToken != data.token.Sync || jsonToken.Expiration.Before(time.Now()) {
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
	tokenID, err := uuid.Parse(jsonToken.Jti)
	if err != nil {
		srvc.logger.Error(
			"Failed to parse PASETO Token ID UUID string",
			zap.String("uuidString", jsonToken.Jti),
		)
		return nil, ctxerr.NewInternal(err)
	}
	userID, err := uuid.Parse(jsonToken.Subject)
	if err != nil {
		srvc.logger.Error(
			"Failed to parse PASETO Subject UUID string",
			zap.String("uuidString", jsonToken.Subject),
		)
		return nil, ctxerr.NewInternal(err)
	}
	return &TokenData{
		TokenID:  tokenID,
		UserID:   userID,
		IssuedAt: jsonToken.IssuedAt,
	}, nil
}

func isValidSyncToken(s string) bool {
	if len(s) > maxSyncTokenLen || !base64url.Test(s) {
		return false
	}
	return true
}
