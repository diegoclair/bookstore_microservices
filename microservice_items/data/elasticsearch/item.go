package elasticsearch

import "github.com/olivere/elastic"

type itemRepo struct {
	db *elastic.Client
}

// newItemRepo returns a instance of dbrepo
func newItemRepo(db *elastic.Client) *itemRepo {
	return &itemRepo{
		db: db,
	}
}
