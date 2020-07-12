package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_oauth/domain"
	"github.com/diegoclair/microservice_oauth/utils/cryptoutils"
)

// AccessTokenRequest struct
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	//Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	//Used for client_credentials grant type
	ClientID     int64  `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

//Validate checks if the access token request is valid
func (at *AccessTokenRequest) Validate() resterrors.RestErr {
	switch at.GrantType {
	case domain.GrantTypePassowrd:
		break

	case domain.GrantTypeClientCredentials:
		break

	default:
		errCode := "Error 0014: "
		return resterrors.NewBadRequestError(fmt.Sprintf("%sInvalid grant_type parameter", errCode))
	}

	return nil
}

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

//Validate checks if the token is valid
func (at *AccessToken) Validate() resterrors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return resterrors.NewBadRequestError("Error 0006: Invalid access token")
	}
	if at.UserID <= 0 {
		return resterrors.NewBadRequestError("Error 0007: Invalid user id")
	}
	if at.ClientID <= 0 {
		return resterrors.NewBadRequestError("Error 0008: Invalid client id")
	}
	if at.Expires <= 0 {
		return resterrors.NewBadRequestError("Error 0009: Invalid expiration time")
	}

	return nil
}

// IsExpired checks if the access token is expired
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = cryptoutils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
