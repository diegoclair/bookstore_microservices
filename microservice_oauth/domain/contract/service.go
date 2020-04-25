package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_oauth/domain/entity"
)

// AccessTokenService holds access token operations
type AccessTokenService interface {
	GetByID(userID string) (*entity.AccessToken, *resterrors.RestErr)
	Create(token entity.AccessTokenRequest) (*entity.AccessToken, *resterrors.RestErr)
	UpdateExpirationTime(token entity.AccessToken) *resterrors.RestErr
}

// UserAPIService holds access token operations
type UserAPIService interface {
	LoginUser(email string, password string) (*entity.APIUser, *resterrors.RestErr)
}
