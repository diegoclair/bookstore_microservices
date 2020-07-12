package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_user/domain/entity"
)

// PingService holds a ping service operations
type PingService interface {
}

// UserService holds a user service operations
type UserService interface {
	GetUser(userID int64) (*entity.User, resterrors.RestErr)
	SearchUser(string) ([]entity.User, resterrors.RestErr)
	CreateUser(entity.User) (*entity.User, resterrors.RestErr)
	UpdateUser(entity.User) (*entity.User, resterrors.RestErr)
	DeleteUser(userID int64) resterrors.RestErr
	LoginUser(request entity.LoginRequest) (*entity.User, resterrors.RestErr)
}
