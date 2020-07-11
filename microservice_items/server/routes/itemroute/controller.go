package itemroute

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/diegoclair/bookstore_oauth-go/oauth"
	"github.com/diegoclair/go_utils-lib/resterrors"
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

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid request body")
		httputils.RespondError(w, *restErr)
		return
	}
	defer r.Body.Close()

	var item entity.Item
	err = json.Unmarshal(requestBody, &item)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid item json body")
		httputils.RespondError(w, *restErr)
		return
	}

	item.Seller = oauth.GetCallerID(r)

	response, createErr := s.itemService.Create(item)
	if err != nil {
		httputils.RespondError(w, *createErr)
		return
	}

	httputils.RespondJSON(w, http.StatusCreated, response)
}

func (s *Controller) handleGetByID(w http.ResponseWriter, r *http.Request) {

}
