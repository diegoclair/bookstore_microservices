package service

import (
	"net/http"

	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_items/domain/contract"
	"github.com/diegoclair/microservice_items/domain/entity"
)

type itemService struct {
	svc *Service
}

//newItemService return a new instance of the service
func newItemService(svc *Service) contract.ItemService {
	return &itemService{
		svc: svc,
	}
}

func (itemService) Create(item entity.Item) (retVal *entity.Item, err *resterrors.RestErr) {

	return retVal, resterrors.NewRestError("Implement me", http.StatusNotImplemented, "not_implemented")
}

func (itemService) GetByID(ID string) (retVal *entity.Item, err *resterrors.RestErr) {

	return retVal, resterrors.NewRestError("Implement me", http.StatusNotImplemented, "not_implemented")
}
