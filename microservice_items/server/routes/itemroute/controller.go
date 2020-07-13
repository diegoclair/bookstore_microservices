package itemroute

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/diegoclair/go_oauth-lib/oauth"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_items/domain/contract"
	"github.com/diegoclair/microservice_items/domain/entity"
	"github.com/diegoclair/microservice_items/utils/httputils"
	"github.com/gorilla/mux"
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
		httputils.RespondError(w, err)
		return
	}

	sellerID := oauth.GetCallerID(r)

	if sellerID == 0 {
		restErr := resterrors.NewUnauthorizedError("Invalid access_token")
		httputils.RespondError(w, restErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid request body")
		httputils.RespondError(w, restErr)
		return
	}
	defer r.Body.Close()

	var item entity.Item
	err = json.Unmarshal(requestBody, &item)
	if err != nil {
		log.Println(err)
		restErr := resterrors.NewBadRequestError("Invalid item json body")
		httputils.RespondError(w, restErr)
		return
	}

	item.Seller = sellerID

	response, createErr := s.itemService.CreateItem(item)
	if createErr != nil {
		httputils.RespondError(w, createErr)
		return
	}

	httputils.RespondJSON(w, http.StatusCreated, response)
}

func (s *Controller) handleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["item_id"]

	response, err := s.itemService.GetByID(itemID)
	if err != nil {
		httputils.RespondError(w, err)
		return
	}

	httputils.RespondJSON(w, http.StatusOK, response)
}

func (s *Controller) handleSearch(w http.ResponseWriter, r *http.Request) {

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Invalid json body: ", err)
		httputils.RespondError(w, resterrors.NewInternalServerError("Invalid json body"))
		return
	}
	defer r.Body.Close()

	var query entity.EsQuery
	err = json.Unmarshal(bytes, &query)
	if err != nil {
		logger.Error("Invalid unmarshal json body: ", err)
		httputils.RespondError(w, resterrors.NewInternalServerError("Invalid to unmarshal json body"))
		return
	}

	response, apiErr := s.itemService.Search(query)
	if apiErr != nil {
		httputils.RespondError(w, apiErr)
		return
	}

	httputils.RespondJSON(w, http.StatusOK, response)
}
