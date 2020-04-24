package data

import (
	"github.com/diegoclair/bookstore_users-api/data/mysql"
	"github.com/diegoclair/bookstore_users-api/domain/contract"
)

// Connect returns a instace of cassandra db
func Connect() (contract.RepoManager, error) {
	return mysql.Instance()
}
