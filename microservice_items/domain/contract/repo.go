package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_items/domain/entity"
	"github.com/olivere/elastic"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Item() ItemRepo
	Elastic() ElasticRepo
}

// ItemRepo defines the data set for items
type ItemRepo interface {
	CreateItem(item entity.Item) (*entity.Item, resterrors.RestErr)
	GetByID(docID string) (*entity.Item, resterrors.RestErr)
	SearchItems(entity.EsQuery) ([]entity.Item, resterrors.RestErr)
}

// ElasticRepo defines the data set for elastic functions
type ElasticRepo interface {
	CreateIndex(docIndex string, docType string, doc interface{}) (*elastic.IndexResponse, error)
	GetDocumentByID(docIndex string, docType string, docID string) (*elastic.GetResult, error)
	SearchDocuments(docIndex string, query elastic.Query) (*elastic.SearchResult, error)
}
