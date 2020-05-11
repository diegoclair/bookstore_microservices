package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/diegoclair/microservice_user/domain/contract"
	"github.com/diegoclair/microservice_user/infra/config"
	_ "github.com/go-sql-driver/mysql" //Used to connect to database
)

// DBManager is the MySQL connection manager
type DBManager struct {
	db *sql.DB
}

var Migration = `CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT,
	first_name VARCHAR(30) NOT NULL,
	last_name VARCHAR(30) NOT NULL,
	email VARCHAR(50),
	password VARCHAR(100),
	status VARCHAR(30),
	created_at TIMESTAMP,
	PRIMARY KEY (id)
);`

//Instance retunrs an instance of a RepoManager
func Instance() (contract.RepoManager, error) {
	cfg := config.GetDBConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	log.Println("Connecting to database...")

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	_, err = db.Query(Migration)
	if err != nil {
		log.Println("failed to run migrations", err.Error())
		return nil, err
	}
	log.Println("Database successfully configured")

	instance := &DBManager{
		db: db,
	}

	return instance, nil
}

//Ping returns a session to use cassadra querys
func (c *DBManager) Ping() contract.PingRepo {
	return nil
}

//User returns a session to use cassadra querys
func (c *DBManager) User() contract.UserRepo {
	return newUserDBClient(c.db)
}
