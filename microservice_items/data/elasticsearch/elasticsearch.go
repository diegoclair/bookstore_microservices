package elasticsearch

import (
	"time"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/microservice_items/domain/contract"
	"github.com/olivere/elastic"
)

// DBManager is the EslasticSearch connection manager
type DBManager struct {
	db *elastic.Client
}

//Instance retunrs an instance of a RepoManager
func Instance() (contract.RepoManager, error) {

	log := logger.GetLogger()
	logger.Info("Configuring the database...")

	client, err := elastic.NewClient(
		elastic.SetURL("http://db:9200"), //db is the elastic service  in docker-compose
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		return nil, err
	}

	logger.Info("Database successfully configured...")

	instance := &DBManager{
		db: client,
	}
	return instance, nil
}

//Elastic receive db session
func (c *DBManager) Elastic() contract.ElasticRepo {
	return newElasticRepo(c.db)
}

//Item returns the item set
func (c *DBManager) Item() contract.ItemRepo {
	return newItemRepo(c.Elastic())
}
