package elasticsearch

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_items/domain"
	"github.com/diegoclair/microservice_items/domain/contract"
	"github.com/diegoclair/microservice_items/domain/entity"
)

type itemRepo struct {
	esr contract.ElasticRepo
}

// newItemRepo returns a instance of dbrepo
func newItemRepo(elasticRepo contract.ElasticRepo) *itemRepo {
	return &itemRepo{
		esr: elasticRepo,
	}
}

// CreateItem - to insert a new item in elasticsearch
func (c *itemRepo) CreateItem(item entity.Item) (*entity.Item, resterrors.RestErr) {

	result, err := c.esr.CreateIndex(domain.DocIndex, domain.DocType, item)
	if err != nil {
		logger.Error("Error when trying to create index document", err)
		return nil, resterrors.NewInternalServerError("Error when trying to save item - Database error")
	}

	item.ID = result.Id

	return &item, nil
}

// GetByID - get item by id
func (c *itemRepo) GetByID(id string) (*entity.Item, resterrors.RestErr) {

	result, err := c.esr.GetDocumentByID(domain.DocIndex, domain.DocType, id)
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to get id %s", id), err)

		if strings.Contains(err.Error(), "404") {
			return nil, resterrors.NewNotFoundError("No item found with id: " + id)
		}
		return nil, resterrors.NewInternalServerError(fmt.Sprintf("Database error - Error when trying to get id %s", id))
	}

	var item entity.Item
	err = json.Unmarshal(*result.Source, &item)
	if err != nil {
		logger.Error("Error to parse database response: ", err)
		return nil, resterrors.NewInternalServerError("Error to parse database response")
	}

	item.ID = result.Id

	return &item, nil
}

func (c *itemRepo) SearchItems(query entity.EsQuery) ([]entity.Item, resterrors.RestErr) {

	finalQuery := query.Build()
	result, err := c.esr.SearchDocuments(domain.DocIndex, finalQuery)
	if err != nil {
		logger.Error("Error when trying to search documents: ", err)
		return nil, resterrors.NewInternalServerError("Error when trying to search documents - Database error")
	}

	var items []entity.Item
	for _, hit := range result.Hits.Hits {
		var item entity.Item

		err := json.Unmarshal(*hit.Source, &item)
		if err != nil {
			logger.Error("Error when trying to unmarshal documents result to item struct: ", err)
			return nil, resterrors.NewInternalServerError("Error when trying to unmarshal documents result to item struct")
		}

		item.ID = hit.Id
		items = append(items, item)
	}

	if len(items) == 0 {
		return nil, resterrors.NewNotFoundError("No items found matching the given criteria")
	}

	return items, nil
}
