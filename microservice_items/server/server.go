package server

import (
	"github.com/diegoclair/microservice_items/server/routes/itemroute"
	"github.com/diegoclair/microservice_items/server/routes/pingroute"
	"github.com/diegoclair/microservice_items/service"
	"github.com/gorilla/mux"
)

type controller struct {
	itemController *itemroute.Controller
	pingController *pingroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *mux.Router {
	svm := service.NewServiceManager()
	srv := mux.NewRouter()

	return setupRoutes(srv, &controller{
		itemController: itemroute.NewController(svm.ItemService(svc)),
		pingController: pingroute.NewController(),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *mux.Router, s *controller) *mux.Router {

	itemroute.NewRouter(s.itemController, srv).RegisterRoutes()
	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()

	return srv
}
