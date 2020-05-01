package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_items/domain/entity"
)

// ItemService holds a item service operations
type ItemService interface {
	Create(item entity.Item) (retVal *entity.Item, err *resterrors.RestErr)
	GetByID(ID string) (retVal *entity.Item, err *resterrors.RestErr)
}
