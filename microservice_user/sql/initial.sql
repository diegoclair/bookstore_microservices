Use sdb_users;


CREATE TABLE IF NOT EXISTS `users` (
	id INT AUTO_INCREMENT,
	first_name VARCHAR(30) NOT NULL,
	last_name VARCHAR(30) NOT NULL,
	email VARCHAR(50),
	password VARCHAR(100),
	status VARCHAR(30),
	created_at TIMESTAMP,
	PRIMARY KEY (id)
);

