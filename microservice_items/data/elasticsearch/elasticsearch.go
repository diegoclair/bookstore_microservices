package elasticsearch

import (
	"time"

	"github.com/diegoclair/microservice_items/domain/contract"
	"github.com/diegoclair/microservice_items/logger"
	"github.com/olivere/elastic"
)

// DBManager is the EslasticSearch connection manager
type DBManager struct {
	db *elastic.Client
}

//Instance retunrs an instance of a RepoManager
func Instance() (contract.RepoManager, error) {

	logger.Info("Configuring the database...")

	client, err := elastic.NewClient(
		elastic.SetURL("http://db:9200"), //db is the elastic service  in docker-compose
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
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

//Item returns the item set
func (c *DBManager) Item() contract.ItemRepo {
	return newItemRepo(c.db)
}
