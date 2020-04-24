package cassandra

import (
	"github.com/diegoclair/bookstore_oauth-api/domain/contract"
	"github.com/gocql/gocql"
)

// DBManager is the Cassandra connection manager
type DBManager struct {
	session *gocql.Session
}

//Instance retunrs an instance of a RepoManager
func Instance() (contract.RepoManager, error) {
	cluster := getDBConfig()

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	instance := &DBManager{
		session: session,
	}

	return instance, nil
}

func getDBConfig() *gocql.ClusterConfig {
	// Config Cassandra cluster
	var cluster *gocql.ClusterConfig
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	return cluster
}

//AccessToken returns a session to use cassadra querys
func (c *DBManager) AccessToken() contract.AccessTokenRepo {
	return newAccessTokenDBSession(c.session)
}
