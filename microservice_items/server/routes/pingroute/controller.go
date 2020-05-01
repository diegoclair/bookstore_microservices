package pingroute

import (
	"net/http"
	"sync"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds user handlers
type Controller struct {
}

//NewController to handle requests
func NewController() *Controller {
	once.Do(func() {
		instance = &Controller{}
	})
	return instance
}

// handlePing - handle a Ping request
func (s *Controller) handlePing(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
