package httprest

import (
	"encoding/json"
	"fmt"

	"github.com/diegoclair/bookstore_oauth-api/domain"
	"github.com/diegoclair/bookstore_oauth-api/domain/contract"
	"github.com/diegoclair/bookstore_oauth-api/domain/entity"
	"github.com/diegoclair/bookstore_oauth-api/utils/errors"
)

type user struct {
}

//NewUserAPI return a new instance of UserAPI
func NewUserAPI() contract.UserAPIService {
	return &user{}
}

func (s *user) LoginUser(email, password string) (*entity.APIUser, *errors.RestErr) {

	body := entity.APIUserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := domain.UserRestClient.Post("/users/login", body)
	if response == nil || response.Response == nil {
		errCode := "Error 0011: "
		return nil, errors.NewInternalServerError(fmt.Sprintf("%sInvalid restclient response when trying to login the user", errCode))
	}

	if response.StatusCode > 299 {
		errCode := "Error 0012: "
		var restErr errors.RestErr

		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError(fmt.Sprintf("%sError when trying to unmarshal the login user response", errCode))
		}
		return nil, &restErr
	}
	var user entity.APIUser
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		errCode := "Error 0013: "
		return nil, errors.NewInternalServerError(fmt.Sprintf("%sError when trying to unmarshal the login response to User struct", errCode))
	}
	return &user, nil
}
