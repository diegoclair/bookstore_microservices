package contract

import (
	"github.com/diegoclair/microservice_oauth/domain/entity"
	"github.com/diegoclair/microservice_oauth/utils/errors"
)

// AccessTokenService holds access token operations
type AccessTokenService interface {
	GetByID(userID string) (*entity.AccessToken, *errors.RestErr)
	Create(token entity.AccessTokenRequest) (*entity.AccessToken, *errors.RestErr)
	UpdateExpirationTime(token entity.AccessToken) *errors.RestErr
}

// UserAPIService holds access token operations
type UserAPIService interface {
	LoginUser(email string, password string) (*entity.APIUser, *errors.RestErr)
}
