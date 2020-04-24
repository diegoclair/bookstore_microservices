package contract

import (
	"github.com/diegoclair/bookstore_users-api/domain/entity"
	"github.com/diegoclair/bookstore_users-api/utils/errors"
)

// PingService holds access token operations
type PingService interface {
}

// UserService holds access token operations
type UserService interface {
	GetUser(userID int64) (*entity.User, *errors.RestErr)
	SearchUser(string) ([]entity.User, *errors.RestErr)
	CreateUser(entity.User) (*entity.User, *errors.RestErr)
	UpdateUser(entity.User) (*entity.User, *errors.RestErr)
	DeleteUser(userID int64) *errors.RestErr
	LoginUser(request entity.LoginRequest) (*entity.User, *errors.RestErr)
}
