package pingroute

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PingRouter holds the ping handlers
type PingRouter struct {
	ctrl   *Controller
	router *mux.Router
}

// NewRouter returns a new PingRouter instance
func NewRouter(ctrl *Controller, router *mux.Router) *PingRouter {
	return &PingRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of ping requests
func (r *PingRouter) RegisterRoutes() {
	r.router.HandleFunc("/ping", r.ctrl.handlePing).Methods(http.MethodGet)
}
