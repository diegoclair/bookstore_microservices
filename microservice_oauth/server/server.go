package server

import (
	"github.com/diegoclair/bookstore_oauth-api/server/routes/oauthroute"
	"github.com/diegoclair/bookstore_oauth-api/service"
	"github.com/gin-gonic/gin"
)

type controller struct {
	oauthController *oauthroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	svm := service.NewServiceManager()
	srv := gin.Default()

	return setupRoutes(srv, &controller{
		oauthController: oauthroute.NewController(svm.AccessTokenService(svc)),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, s *controller) *gin.Engine {

	oauthroute.NewRouter(s.oauthController, srv).RegisterRoutes()

	return srv
}
