package oauthroute

import (
	"net/http"
	"sync"

	"github.com/diegoclair/bookstore_oauth-api/domain/contract"
	"github.com/diegoclair/bookstore_oauth-api/domain/entity"
	"github.com/diegoclair/bookstore_oauth-api/utils/errors"
	"github.com/gin-gonic/gin"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds access token handlers
type Controller struct {
	accessTokenService contract.AccessTokenService
}

//NewController to handle requests
func NewController(accessTokenService contract.AccessTokenService) *Controller {
	once.Do(func() {
		instance = &Controller{
			accessTokenService: accessTokenService,
		}
	})
	return instance
}

// handleGetByID to handle a get by id request
func (s *Controller) handleGetByID(c *gin.Context) {
	accessToken, err := s.accessTokenService.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

// handleCreate to handle a create request
func (s *Controller) handleCreate(c *gin.Context) {
	var request entity.AccessTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("Error 0010: Invalid Json Body Request")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	accessToken, err := s.accessTokenService.Create(request)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}
