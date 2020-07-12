package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_items/domain/entity"
	"github.com/olivere/elastic"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Item() ItemRepo
}

// ItemRepo defines the data set for items
type ItemRepo interface {
	Index(index string, esType string, doc interface{}) (*elastic.IndexResponse, error)
	Save(i entity.Item) (item entity.Item, restErr *resterrors.RestErr)
}
