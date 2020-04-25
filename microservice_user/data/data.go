package data

import (
	"github.com/diegoclair/microservice_user/data/mysql"
	"github.com/diegoclair/microservice_user/domain/contract"
)

// Connect returns a instace of cassandra db
func Connect() (contract.RepoManager, error) {
	return mysql.Instance()
}
