package cache

import (
	"halo-suster/internal/delivery/http/v1/response"
	"time"
)

type AccessToken struct {
	JWTClaim response.UserAccessToken
	Expired  time.Time
}

type (
	cacheAccessToken map[string]AccessToken
)

var (
	accessToken *cacheAccessToken
)
