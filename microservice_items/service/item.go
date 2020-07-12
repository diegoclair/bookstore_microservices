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

func (s *itemService) Create(item entity.Item) (*entity.Item, resterrors.RestErr) {

	result, err := s.svc.db.Item().Save(item)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *itemService) GetByID(ID string) (retVal *entity.Item, err resterrors.RestErr) {

	return retVal, resterrors.NewRestError("Implement me", http.StatusNotImplemented, "not_implemented")
}
