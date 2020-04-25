package service

import (
	"strings"

	"github.com/diegoclair/microservice_user/domain"
	"github.com/diegoclair/microservice_user/domain/contract"
	"github.com/diegoclair/microservice_user/domain/entity"
	"github.com/diegoclair/microservice_user/utils/cryptoutils"
	"github.com/diegoclair/microservice_user/utils/dateutils"
	"github.com/diegoclair/microservice_user/utils/errors"
)

/* Here we have the entire business logic*/

type userService struct {
	svc *Service
}

//newUserService return a new instance of the service
func newUserService(svc *Service) contract.UserService {
	return &userService{
		svc: svc,
	}
}

func (s *userService) GetUser(userID int64) (*entity.User, *errors.RestErr) {
	user := &entity.User{
		ID: userID,
	}

	user, err := s.svc.db.User().GetByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) SearchUser(status string) ([]entity.User, *errors.RestErr) {
	return s.svc.db.User().GetUserByStatus(status) // The two functions return the same values
}

func (s *userService) CreateUser(user entity.User) (*entity.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = domain.StatusActive
	user.CreatedAt = dateutils.GetCompleteDateNowDBLayout()
	user.Password = cryptoutils.GetMd5(user.Password)

	newUser, err := s.svc.db.User().Create(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) UpdateUser(user entity.User) (*entity.User, *errors.RestErr) {

	// To not update with "" others fields that we don't send in the request and to return  this others fields,
	// like the created_at in the response, if we don't do this, the field created_at, will be show with the value = ""
	currentUser, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if user.FirstName != "" {
		currentUser.FirstName = strings.TrimSpace(user.FirstName)
	}
	if user.LastName != "" {
		currentUser.LastName = strings.TrimSpace(user.LastName)
	}
	if user.Email != "" {
		currentUser.Email = strings.TrimSpace(user.Email)
	}

	updatedUser, err := s.svc.db.User().Update(*currentUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *userService) DeleteUser(userID int64) *errors.RestErr {
	return s.svc.db.User().Delete(userID)
}

func (s *userService) LoginUser(request entity.LoginRequest) (*entity.User, *errors.RestErr) {
	user := &entity.User{
		Email:    request.Email,
		Password: cryptoutils.GetMd5(request.Password),
	}

	user, err := s.svc.db.User().GetByEmailAndPassword(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
