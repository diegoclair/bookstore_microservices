package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/diegoclair/bookstore_oauth-api/domain/contract"
	"github.com/diegoclair/bookstore_oauth-api/domain/entity"
	"github.com/diegoclair/bookstore_oauth-api/utils/errors"
)

const (
	tokenExpirationTime = 24
)

type accessToken struct {
	svc     *Service
	userAPI contract.UserAPIService
}

//newAccessTokenService return a new instance of the service
func newAccessTokenService(svc *Service, api contract.UserAPIService) contract.AccessTokenService {
	return &accessToken{
		svc:     svc,
		userAPI: api,
	}
}

func (s *accessToken) GetByID(accessTokenID string) (*entity.AccessToken, *errors.RestErr) {

	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("Error 0005: Invalid or expired access token")
	}
	return s.svc.db.AccessToken().GetByID(accessTokenID)
}

func (s *accessToken) Create(request entity.AccessTokenRequest) (*entity.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	//Authenticate the user against the Users API:
	user, err := s.userAPI.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	at := GetNewAccessToken(user.ID)
	at.Generate()

	if err := s.svc.db.AccessToken().Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *accessToken) UpdateExpirationTime(token entity.AccessToken) *errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.svc.db.AccessToken().UpdateExpirationTime(token)
}

// GetNewAccessToken to get a new access token
func GetNewAccessToken(userID int64) entity.AccessToken {
	return entity.AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(tokenExpirationTime * time.Hour).Unix(),
	}
}
