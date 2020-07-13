package service

import (
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

func (s *itemService) CreateItem(item entity.Item) (*entity.Item, resterrors.RestErr) {
	return s.svc.db.Item().CreateItem(item)
}

func (s *itemService) GetByID(id string) (*entity.Item, resterrors.RestErr) {
	return s.svc.db.Item().GetByID(id)
}

func (s *itemService) Search(query entity.EsQuery) ([]entity.Item, resterrors.RestErr) {
	return s.svc.db.Item().SearchItems(query)
}
