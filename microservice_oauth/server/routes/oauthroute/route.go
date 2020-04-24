package oauthroute

import (
	"github.com/gin-gonic/gin"
)

// OAuthRouter holds the user handlers
type OAuthRouter struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRouter returns a new OAuthRouter instance
func NewRouter(ctrl *Controller, router *gin.Engine) *OAuthRouter {
	return &OAuthRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of access token request
func (r *OAuthRouter) RegisterRoutes() {
	r.router.GET("/oauth/accesstoken/:access_token_id", r.ctrl.handleGetByID)
	r.router.POST("/oauth/accesstoken", r.ctrl.handleCreate)
}
