package userroute

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/diegoclair/bookstore_oauth-go/oauth"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_user/domain/contract"
	"github.com/diegoclair/microservice_user/domain/entity"
	"github.com/gin-gonic/gin"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds user handlers
type Controller struct {
	userService contract.UserService
}

//NewController to handle requests
func NewController(userService contract.UserService) *Controller {
	once.Do(func() {
		instance = &Controller{
			userService: userService,
		}
	})
	return instance
}

// handleGetUser to handle a get user request
func (s *Controller) handleGetUser(c *gin.Context) {

	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	userID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	user, getErr := s.userService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	fmt.Println(oauth.GetCallerID(c.Request), user.ID)
	if oauth.GetCallerID(c.Request) == user.ID {
		//If the caller is the user, so we can show a private data
		c.JSON(http.StatusOK, s.limitedJSON(*user, false))
		return
	}

	c.JSON(http.StatusOK, s.limitedJSON(*user, oauth.IsPublic(c.Request)))
}

// handleCreateUser to handle a create user request
func (s *Controller) handleCreateUser(c *gin.Context) {

	var user entity.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	result, createErr := s.userService.CreateUser(user)
	if createErr != nil {
		c.JSON(createErr.StatusCode, createErr)
		return
	}

	c.JSON(http.StatusCreated, s.limitedJSON(*result, c.GetHeader("X-Public") == "true"))
}

// handleUpdateUser to handle a update user request
func (s *Controller) handleUpdateUser(c *gin.Context) {
	var user entity.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	userID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	user.ID = userID

	resUser, updateErr := s.userService.UpdateUser(user)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}

	c.JSON(http.StatusOK, s.limitedJSON(*resUser, c.GetHeader("X-Public") == "true"))
}

// handleDeleteUser to handle a delete user request
func (s *Controller) handleDeleteUser(c *gin.Context) {

	userID, errID := s.getIDParameter(c.Param("id"))
	if errID != nil {
		c.JSON(errID.StatusCode, errID)
		return
	}

	deleteErr := s.userService.DeleteUser(userID)
	if deleteErr != nil {
		c.JSON(deleteErr.StatusCode, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

// handleLogin to handle a user login request
func (s *Controller) handleLogin(c *gin.Context) {
	var credentials = entity.LoginRequest{}

	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	resUser, loginErr := s.userService.LoginUser(credentials)
	if loginErr != nil {
		c.JSON(loginErr.StatusCode, loginErr)
		return
	}

	c.JSON(http.StatusOK, s.limitedJSON(*resUser, c.GetHeader("X-Public") == "true"))
}

// handleSearch - handle a Search request and return all users by some parameter
func (s *Controller) handleSearch(c *gin.Context) {
	status := c.Query("status")

	users, getErr := s.userService.SearchUser(status)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}

	c.JSON(http.StatusOK, s.listLimitedJSON(users, c.GetHeader("X-Public") == "true"))
}

//listLimitedJSON check if the return will be a public or private list interface of user
func (s *Controller) listLimitedJSON(users []entity.User, isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = s.limitedJSON(user, isPublic)
	}
	return result
}

//limitedJSON check if the return will be a public or private interface of user
func (s *Controller) limitedJSON(user entity.User, isPublic bool) interface{} {
	//We have two option to return the map of User to Private or Public user

	// 1 - if we have some json that have different key:
	// Ex: in User{} we have the `json:"id"` and in the PrivateUser{} we have the key as `json:"user_id"`
	// so we need to do one by one like below
	if isPublic {
		return entity.PublicUser{
			ID:        user.ID,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
		}
	}

	// 2 - if the keys are the same so we can do just a json.Marshal and after a json.Unmashal to fill the colunms
	userJSON, _ := json.Marshal(user)
	var privateUser entity.PrivateUser
	json.Unmarshal(userJSON, &privateUser)

	return privateUser
}

func (s *Controller) getIDParameter(userParamID string) (int64, *resterrors.RestErr) {
	userID, userErr := strconv.ParseInt(userParamID, 10, 64)
	if userErr != nil {
		return 0, resterrors.NewBadRequestError("User id should be a number")
	}

	return userID, nil
}
