package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/diegoclair/bookstore_users-api/domain/contract"
	"github.com/diegoclair/bookstore_users-api/infra/config"
	_ "github.com/go-sql-driver/mysql" //Used to connect to database
)

// DBManager is the MySQL connection manager
type DBManager struct {
	client *sql.DB
}

//Instance retunrs an instance of a RepoManager
func Instance() (contract.RepoManager, error) {
	cfg := config.GetDBConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		cfg.Username, cfg.Password, cfg.Host, cfg.Schema,
	)

	log.Println("Connecting to database...")

	client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(); err != nil {
		return nil, err
	}
	log.Println("Database successfully configured")

	instance := &DBManager{
		client: client,
	}

	return instance, nil
}

//Ping returns a session to use cassadra querys
func (c *DBManager) Ping() contract.PingRepo {
	return nil
}

//User returns a session to use cassadra querys
func (c *DBManager) User() contract.UserRepo {
	return newUserDBClient(c.client)
}
