package data

import (
	"github.com/diegoclair/microservice_oauth/data/cassandra"
	"github.com/diegoclair/microservice_oauth/domain/contract"
)

// Connect returns a instace of cassandra db
func Connect() (contract.RepoManager, error) {
	return cassandra.Instance()
}
