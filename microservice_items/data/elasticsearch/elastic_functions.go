package elasticsearch

import (
	"context"

	"github.com/olivere/elastic"
)

type elasticRepo struct {
	db *elastic.Client
}

// newElasticRepo returns a instance of dbrepo
func newElasticRepo(db *elastic.Client) *elasticRepo {
	return &elasticRepo{
		db: db,
	}
}

// CreateIndex - create new index document in elasticsearch
func (c *elasticRepo) CreateIndex(docIndex string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()

	result, err := c.db.Index().
		Index(docIndex).
		Type(docType).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDocumentByID - get a document by id
func (c *elasticRepo) GetDocumentByID(docIndex string, docType string, docID string) (*elastic.GetResult, error) {
	ctx := context.Background()

	result, err := c.db.Get().
		Index(docIndex).
		Type(docType).
		Id(docID).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *elasticRepo) SearchDocuments(docIndex string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()

	result, err := c.db.Search().
		Index(docIndex).
		Query(query).
		//Sort("user", true). // sort by "user" field, ascending
		//From(0).Size(10).   // take documents 0-9
		RestTotalHitsAsInt(true).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
