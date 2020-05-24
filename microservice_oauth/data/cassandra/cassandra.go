package cassandra

import (
	"log"
	"time"

	"github.com/diegoclair/microservice_oauth/domain/contract"
	"github.com/diegoclair/microservice_oauth/domain/entity"
	"github.com/diegoclair/microservice_oauth/infra/config"
	"github.com/gocql/gocql"
)

// DBManager is the Cassandra connection manager
type DBManager struct {
	session *gocql.Session
}

//Instance retunrs an instance of a RepoManager
func Instance() (contract.RepoManager, error) {
	log.Println("Connecting to database...")

	log.Println("Getting configs")
	cfg := config.GetDBConfig()

	err := createInitialKeyspace(cfg)
	if err != nil {
		return nil, err
	}

	log.Println("Creating database Session...")

	cluster := getDBConfig(cfg)

	//now that we have our keyspace oauth created, we can create a new session
	cluster.Keyspace = cfg.DBName

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	log.Println("Creating initial table and data...")
	err = createInitialTableAndData(cfg, session)
	if err != nil {
		return nil, err
	}

	log.Println("Database successfully configured...")
	instance := &DBManager{
		session: session,
	}

	return instance, nil
}

func createInitialKeyspace(cfg entity.InitialConfig) error {

	log.Println("Creating database initial Session...")

	//use initialCluster without keyspace to create our keyspace oauth
	initialCluster := getDBConfig(cfg)

	initialSession, err := initialCluster.CreateSession()
	if err != nil {
		return err
	}

	// Check if the table already exists. Create if table does not exist
	log.Println("Creating new keyspace:", cfg.DBName)

	initialSession.Query(`CREATE KEYSPACE IF NOT EXISTS oauth 
		WITH replication = {'class' : 'SimpleStrategy', 'replication_factor' : 1};
	`).Exec()

	initialSession.Close()

	log.Println("Keyspace and table created successfully")
	return nil
}

func createInitialTableAndData(cfg entity.InitialConfig, session *gocql.Session) error {
	keySpaceMeta, err := session.KeyspaceMetadata(cfg.DBName)
	if err != nil {
		log.Println("Error keySpaceMeta: ", err)
	}

	if _, exists := keySpaceMeta.Tables["access_token"]; exists != true {

		log.Println("Creating new table access_token")

		err := session.Query(`
			CREATE TABLE IF NOT EXISTS access_token (
				access_token text, 
				user_id int, 
				client_id int, 
				expires int,
				PRIMARY KEY (access_token));
		`).Exec()
		if err != nil {
			log.Println("Error to create table", err)
		}

		//create a function to input data
	}
	return nil
}

func getDBConfig(cfg entity.InitialConfig) *gocql.ClusterConfig {

	// Config Cassandra cluster
	var cluster *gocql.ClusterConfig

	// Provide the cassandra cluster instance here.
	cluster = gocql.NewCluster(cfg.Host)

	// The authenticator is needed if password authentication is
	// enabled for your Cassandra installation. If not, this can
	// be removed.
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	// gocql requires the keyspace to be provided before the session is created.
	// In future there might be provisions to do this later.
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10

	return cluster
}

//AccessToken returns a session to use cassadra querys
func (c *DBManager) AccessToken() contract.AccessTokenRepo {
	return newAccessTokenDBSession(c.session)
}
