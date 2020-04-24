package contract

import (
	"github.com/diegoclair/bookstore_oauth-api/domain/entity"
	"github.com/diegoclair/bookstore_oauth-api/utils/errors"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	AccessToken() AccessTokenRepo
}

// AccessTokenRepo defines the data set for access token
type AccessTokenRepo interface {
	GetByID(string) (*entity.AccessToken, *errors.RestErr)
	Create(token entity.AccessToken) *errors.RestErr
	UpdateExpirationTime(token entity.AccessToken) *errors.RestErr
}
