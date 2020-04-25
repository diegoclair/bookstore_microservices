package contract

import (
	"github.com/diegoclair/microservice_user/domain/entity"
	"github.com/diegoclair/microservice_user/utils/errors"
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
	GetByID(userID int64) (*entity.User, *errors.RestErr)
	GetUserByStatus(status string) ([]entity.User, *errors.RestErr)
	Create(entity.User) (*entity.User, *errors.RestErr)
	Update(entity.User) (*entity.User, *errors.RestErr)
	Delete(userID int64) *errors.RestErr
	GetByEmailAndPassword(user entity.User) (*entity.User, *errors.RestErr)
}
