package itemroute

import (
	"net/http"
	"sync"

	"github.com/diegoclair/bookstore_oauth-go/oauth"
	"github.com/diegoclair/microservice_items/domain/contract"
	"github.com/diegoclair/microservice_items/domain/entity"
	"github.com/diegoclair/microservice_items/utils/httputils"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds item handlers
type Controller struct {
	itemService contract.ItemService
}

//NewController to handle the requests
func NewController(itemService contract.ItemService) *Controller {
	once.Do(func() {
		instance = &Controller{
			itemService: itemService,
		}
	})
	return instance
}

func (s *Controller) handleCreate(w http.ResponseWriter, r *http.Request) {

	if err := oauth.AuthenticateRequest(r); err != nil {
		httputils.RespondError(w, *err)
		return
	}

	item := entity.Item{
		Seller: oauth.GetCallerID(r),
	}

	response, err := s.itemService.Create(item)
	if err != nil {
		httputils.RespondError(w, *err)
		return
	}

	httputils.RespondJSON(w, http.StatusCreated, response)

}

func (s *Controller) handleGetByID(w http.ResponseWriter, r *http.Request) {

}
