package server

import (
	"github.com/diegoclair/bookstore_users-api/server/routes/pingroute"
	"github.com/diegoclair/bookstore_users-api/server/routes/userroute"
	"github.com/diegoclair/bookstore_users-api/service"
	"github.com/gin-gonic/gin"
)

type controller struct {
	pingController *pingroute.Controller
	userController *userroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	svm := service.NewServiceManager()
	srv := gin.Default()

	return setupRoutes(srv, &controller{
		pingController: pingroute.NewController(),
		userController: userroute.NewController(svm.UserService(svc)),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, s *controller) *gin.Engine {

	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()
	userroute.NewRouter(s.userController, srv).RegisterRoutes()

	return srv
}
