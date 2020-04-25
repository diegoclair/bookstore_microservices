package service

import (
	"testing"
	"time"

	"github.com/diegoclair/microservice_oauth/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, tokenExpirationTime, 24, "Expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(1)
	assert.False(t, at.IsExpired(), "Brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "New access token should not have defined access token id")
	assert.True(t, at.UserID == 0, "New access token should not have an associeated user id")

}

func TestAccessTokenIsExpired(t *testing.T) {
	at := entity.AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Access token expiring three hours from now should NOT be expired byu default")
}
