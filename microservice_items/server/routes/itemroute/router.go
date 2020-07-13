package itemroute

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ItemRouter holds the item handlers
type ItemRouter struct {
	ctrl   *Controller
	router *mux.Router
}

// NewRouter returns a new ItemRouter instance
func NewRouter(ctrl *Controller, router *mux.Router) *ItemRouter {
	return &ItemRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of ping requests
func (r *ItemRouter) RegisterRoutes() {
	r.router.HandleFunc("/item", r.ctrl.handleCreate).Methods(http.MethodPost)
	r.router.HandleFunc("/item/{item_id}", r.ctrl.handleGetByID).Methods(http.MethodGet)

	r.router.HandleFunc("/item/search", r.ctrl.handleSearch).Methods(http.MethodPost)
}
