package access

import "time"

const (
	syncTokenCount  = 32
	maxSyncTokenLen = 2 * syncTokenCount
)

const (
	accessDuration  = 30 * time.Second
	refreshDuration = 7 * 24 * time.Hour
)

const (
	tokenTypeKey = "token-type"
	syncTokenKey = "sync-token"
)
