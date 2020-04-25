package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_oauth/domain/entity"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	AccessToken() AccessTokenRepo
}

// AccessTokenRepo defines the data set for access token
type AccessTokenRepo interface {
	GetByID(string) (*entity.AccessToken, *resterrors.RestErr)
	Create(token entity.AccessToken) *resterrors.RestErr
	UpdateExpirationTime(token entity.AccessToken) *resterrors.RestErr
}
