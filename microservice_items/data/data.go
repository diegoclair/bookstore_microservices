package data

import (
	"github.com/diegoclair/microservice_items/data/elasticsearch"
	"github.com/diegoclair/microservice_items/domain/contract"
)

// Connect returns a instace of cassandra db
func Connect() (contract.RepoManager, error) {
	return elasticsearch.Instance()
}
