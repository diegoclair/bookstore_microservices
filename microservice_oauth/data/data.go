package data

import (
	"github.com/diegoclair/bookstore_oauth-api/data/cassandra"
	"github.com/diegoclair/bookstore_oauth-api/domain/contract"
)

// Connect returns a instace of cassandra db
func Connect() (contract.RepoManager, error) {
	return cassandra.Instance()
}
