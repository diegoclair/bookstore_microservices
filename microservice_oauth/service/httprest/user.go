package httprest

import (
	"encoding/json"
	"fmt"

	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_oauth/domain"
	"github.com/diegoclair/microservice_oauth/domain/contract"
	"github.com/diegoclair/microservice_oauth/domain/entity"
)

type user struct {
}

//NewUserAPI return a new instance of UserAPI
func NewUserAPI() contract.UserAPIService {
	return &user{}
}

func (s *user) LoginUser(email, password string) (*entity.APIUser, *resterrors.RestErr) {

	body := entity.APIUserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := domain.UserRestClient.Post("/users/login", body)
	if response == nil || response.Response == nil {
		errCode := "Error 0011: "
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sInvalid restclient response when trying to login the user", errCode))
	}

	if response.StatusCode > 299 {
		errCode := "Error 0012: "
		var restErr resterrors.RestErr

		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sError when trying to unmarshal the login user response", errCode))
		}
		return nil, &restErr
	}
	var user entity.APIUser
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		errCode := "Error 0013: "
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("%sError when trying to unmarshal the login response to User struct", errCode))
	}
	return &user, nil
}
