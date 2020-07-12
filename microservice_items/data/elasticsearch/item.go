package elasticsearch

import (
	"context"
	"fmt"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_items/domain/entity"
	"github.com/olivere/elastic"
)

const (
	indexItem = "items"
	esType    = "item"
)

type itemRepo struct {
	db *elastic.Client
}

// newItemRepo returns a instance of dbrepo
func newItemRepo(db *elastic.Client) *itemRepo {
	return &itemRepo{
		db: db,
	}
}

//Index - to insert a index in a db
func (c *itemRepo) Index(index string, esType string, doc interface{}) (*elastic.IndexResponse, error) {

	ctx := context.Background()

	result, err := c.db.Index().
		Index(index).
		Type(esType).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}

func (c *itemRepo) Save(i entity.Item) (retVal entity.Item, restErr resterrors.RestErr) {

	result, err := c.Index(indexItem, esType, i)
	if err != nil {
		return retVal, resterrors.NewInternalServerError("Error when trying to save item - Database error")
	}

	i.ID = result.Id

	return i, nil
}
