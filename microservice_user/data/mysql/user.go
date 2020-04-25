package mysql

import (
	"database/sql"
	"fmt"

	"github.com/diegoclair/microservice_user/domain"
	"github.com/diegoclair/microservice_user/domain/entity"
	"github.com/diegoclair/microservice_user/logger"
	"github.com/diegoclair/microservice_user/utils/errors"
	"github.com/diegoclair/microservice_user/utils/mysqlutils"
)

type userDBClient struct {
	client *sql.DB
}

// newUserDBClient returns a instance of dbrepo
func newUserDBClient(client *sql.DB) *userDBClient {
	return &userDBClient{
		client: client,
	}
}

//GetByID - get a user by ID
func (s *userDBClient) GetByID(id int64) (*entity.User, *errors.RestErr) {

	query := `
		SELECT 	u.id,
				u.first_name,
				u.last_name,
				u.email,
				u.status,
				u.created_at

		FROM 	users 		u 
		WHERE 	u.id 		= ?;`

	stmt, err := s.client.Prepare(query)
	if err != nil {
		errorCode := "Error 0001: "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in GetByID", errorCode), err)
		return nil, errors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	var user entity.User

	// If we use result, err := stmt.Query(user.ID) instead QueryRow, then we need to close the connection defer result.Close() and check if we have some err
	// to get only one register on database, is better to use queryRow
	result := stmt.QueryRow(id)
	err = result.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
	)
	if err != nil {
		errorCode := "Error 0002: "
		logger.Error(fmt.Sprintf("%sError when trying to execute QueryRow in GetByID", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &user, nil
}

// GetUserByStatus return a list of all users by status
func (s *userDBClient) GetUserByStatus(status string) (users []entity.User, restErr *errors.RestErr) {

	query := `
			SELECT 	u.id,
					u.first_name,
					u.last_name,
					u.email,
					u.status,
					u.created_at

			FROM	users 		u 

			WHERE 	u.status 	= ?;`

	stmt, err := s.client.Prepare(query)
	if err != nil {
		errorCode := "Error 0003: "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in GetUserByStatus", errorCode), err)
		return nil, errors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		errorCode := "Error 0004: "
		logger.Error(fmt.Sprintf("%sError when trying to execute Query in GetUserByStatus", errorCode), err)
		return nil, errors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Status,
			&user.CreatedAt,
		)
		if err != nil {
			errorCode := "Error 0005: "
			logger.Error(fmt.Sprintf("%sError when trying to do For Scan in the Rows GetUserByStatus", errorCode), err)
			return nil, mysqlutils.HandleMySQLError(errorCode, err)
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No users matching with the status : %s", status))
	}

	return users, nil
}

// Create - to create a user on database
func (s *userDBClient) Create(user entity.User) (*entity.User, *errors.RestErr) {

	query := `
		INSERT INTO users 
				(first_name, last_name, email, password, status, created_at	) 
		VALUES	(		?, 		 ?,  	  ?,       ?, 		?,	   	 ? 		);
		`

	// When you use prepare, you not already execute the query on database, it's like to validate the query first
	// its is (more fast) than when you get an error directly on your database
	stmt, err := s.client.Prepare(query)
	if err != nil {
		errorCode := "Error 0006: "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in the Create user", errorCode), err)
		return nil, errors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.CreatedAt)
	if err != nil {
		errorCode := "Error 0007: "
		logger.Error(fmt.Sprintf("%sError when trying to execute Query in the Create user", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		errorCode := "Error 0008: "
		logger.Error(fmt.Sprintf("%sError when trying to get LastInsertId in the Create user", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	user.ID = userID

	return &user, nil
}

// Update - to update a user on database
func (s *userDBClient) Update(user entity.User) (*entity.User, *errors.RestErr) {

	query := `
		UPDATE users
			SET	first_name 	= ?,
				last_name	= ?,
				email		= ?
			
		WHERE 	id			= ?;
	`

	stmt, err := s.client.Prepare(query)
	if err != nil {
		errorCode := "Error 0009: "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in the Update user", errorCode), err)
		return nil, errors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		errorCode := "Error 0010: "
		logger.Error(fmt.Sprintf("%sError when trying to execute Query in the Update user", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &user, nil
}

// Delete - to delete a user on database
func (s *userDBClient) Delete(id int64) *errors.RestErr {

	query := `
		DELETE FROM users
		WHERE 	id			= ?;
	`

	stmt, err := s.client.Prepare(query)
	if err != nil {
		errorCode := "Error 0011: "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in the Delete user", errorCode), err)
		return errors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		errorCode := "Error 0012: "
		logger.Error(fmt.Sprintf("%sError when trying to execute Query in the Delete user", errorCode), err)
		return mysqlutils.HandleMySQLError(errorCode, err)
	}

	return nil
}

//GetByEmailAndPassword - get a user by their email and password
func (s *userDBClient) GetByEmailAndPassword(user entity.User) (*entity.User, *errors.RestErr) {

	query := `
		SELECT 	u.id,
				u.first_name,
				u.last_name,
				u.email,
				u.status,
				u.created_at

		FROM	users 		u 

		WHERE 	u.email 	= ?
		  AND   u.password	= ?
		  AND   u.status	= ?;`

	stmt, err := s.client.Prepare(query)
	if err != nil {
		errorCode := "Error 0013: "
		logger.Error(fmt.Sprintf("%sError when trying to prepare the query statement in GetByEmailAndPassword", errorCode), err)
		return nil, errors.NewInternalServerError(fmt.Sprintf("%sDatabase error", errorCode))
	}
	defer stmt.Close()

	// If we use result, err := stmt.Query(user.ID) instead QueryRow, then we need to close the connection defer result.Close() and check if we have some err
	// to get only one register on database, is better to use queryRow
	fmt.Println(user.Email, user.Password)
	result := stmt.QueryRow(user.Email, user.Password, domain.StatusActive)
	err = result.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
	)
	if err != nil {
		errorCode := "Error 0014: "
		logger.Error(fmt.Sprintf("%sError when trying to execute QueryRow in GetByEmailAndPassword", errorCode), err)
		return nil, mysqlutils.HandleMySQLError(errorCode, err)
	}

	return &user, nil
}
