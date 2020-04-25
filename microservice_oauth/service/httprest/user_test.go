package httprest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/diegoclair/microservice_oauth/domain"
	"github.com/federicoleon/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	//"github.com/mercadolibre/golang-restclient/rest" this version doesn't work with golang >= 1.13, so we are use the federicoleon while wait for approve the federicoleon PR
)

func TestMain(m *testing.M) {
	fmt.Println(domain.UserRestClient.BaseURL)
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromAPI(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          fmt.Sprintf("%s/users/login", domain.UserRestClient.BaseURL),
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	userRepo := user{}

	user, err := userRepo.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Error 0011: Invalid restclient response when trying to login the user", err.Message)

}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          fmt.Sprintf("%s/users/login", domain.UserRestClient.BaseURL),
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Error 0014: Invalid user credentials","status_code": "404","error": "not_found"}`, //if the status code is now string
	})

	userRepo := user{}

	user, err := userRepo.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Error 0012: Error when trying to unmarshal the login user response", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          fmt.Sprintf("%s/users/login", domain.UserRestClient.BaseURL),
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Error 0014: Invalid user credentials","status_code": 404,"error": "not_found"}`,
	})

	userRepo := user{}

	user, err := userRepo.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "Error 0014: Invalid user credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          fmt.Sprintf("%s/users/login", domain.UserRestClient.BaseURL),
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1","first_name": "Diego Clair","last_name": "Rodrigues","email": "diego93rodrigues@gmail.com"}`,
	})

	userRepo := user{}

	user, err := userRepo.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Error 0013: Error when trying to unmarshal the login response to User struct", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          fmt.Sprintf("%s/users/login", domain.UserRestClient.BaseURL),
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1,"first_name": "Diego Clair","last_name": "Rodrigues","email": "diegotest@gmail.com"}`,
	})

	userRepo := user{}

	user, err := userRepo.LoginUser("email@gmail.com", "the-password")
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "diegotest@gmail.com", user.Email)
	assert.EqualValues(t, "Diego Clair", user.FirstName)
	assert.EqualValues(t, "Rodrigues", user.LastName)

}
