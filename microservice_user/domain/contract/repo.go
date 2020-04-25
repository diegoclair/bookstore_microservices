package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_user/domain/entity"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Ping() PingRepo
	User() UserRepo
}

// PingRepo defines the data set for access token
type PingRepo interface{}

// UserRepo defines the data set for access token
type UserRepo interface {
	GetByID(userID int64) (*entity.User, *resterrors.RestErr)
	GetUserByStatus(status string) ([]entity.User, *resterrors.RestErr)
	Create(entity.User) (*entity.User, *resterrors.RestErr)
	Update(entity.User) (*entity.User, *resterrors.RestErr)
	Delete(userID int64) *resterrors.RestErr
	GetByEmailAndPassword(user entity.User) (*entity.User, *resterrors.RestErr)
}
