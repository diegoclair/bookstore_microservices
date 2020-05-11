package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table users",
			Script: `CREATE TABLE IF NOT EXISTS users (
				id INT AUTO_INCREMENT,
				first_name VARCHAR(30) NOT NULL,
				last_name VARCHAR(30) NOT NULL,
				email VARCHAR(50),
				password VARCHAR(100),
				created_at TIMESTAMP,
				PRIMARY KEY (id)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     2,
			Description: "Adding column status to table users",
			Script:      "ALTER TABLE users ADD status VARCHAR(30) AFTER password;",
		},
	}
)
