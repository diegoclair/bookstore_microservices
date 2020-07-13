package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_items/domain/entity"
)

// ItemService holds a item service operations
type ItemService interface {
	CreateItem(item entity.Item) (*entity.Item, resterrors.RestErr)
	GetByID(ID string) (*entity.Item, resterrors.RestErr)
	Search(query entity.EsQuery) ([]entity.Item, resterrors.RestErr)
}
