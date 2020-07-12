package elasticsearch

import (
	"log"
	"time"

	"github.com/diegoclair/microservice_items/domain/contract"
	"github.com/olivere/elastic"
)

// DBManager is the EslasticSearch connection manager
type DBManager struct {
	db *elastic.Client
}

//Instance retunrs an instance of a RepoManager
func Instance() (contract.RepoManager, error) {

	log.Println("Connecting to database...")

	client, err := elastic.NewClient(
		elastic.SetURL("http://db:9200"), //db is the elastic service  in docker-compose
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		return nil, err
	}

	log.Println("Database successfully configured...")

	instance := &DBManager{
		db: client,
	}
	return instance, nil
}

/* func (c *esClient) Index(interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
} */

//Item returns the item set
func (c *DBManager) Item() contract.ItemRepo {
	return newItemRepo(c.db)
}
