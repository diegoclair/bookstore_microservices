package cassandra

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservice_oauth/domain/entity"
	"github.com/gocql/gocql"
)

type accessTokenDBSession struct {
	session *gocql.Session
}

// newAccessTokenDBSession returns a instance of dbrepo
func newAccessTokenDBSession(session *gocql.Session) *accessTokenDBSession {
	return &accessTokenDBSession{
		session: session,
	}
}

func (s *accessTokenDBSession) GetByID(id string) (*entity.AccessToken, resterrors.RestErr) {

	query := `
		SELECT 	access_token, 
				user_id, 
				client_id, 
				expires
		FROM access_token 		
		WHERE access_token 	= ?;
		`
	var result entity.AccessToken
	err := s.session.Query(query, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, resterrors.NewNotFoundError("Error 0002: No access token found with given id")
		}
		return nil, resterrors.NewInternalServerError("Error 0001: " + err.Error())
	}

	return &result, nil
}

func (s *accessTokenDBSession) Create(at entity.AccessToken) resterrors.RestErr {

	query := `
		INSERT INTO	access_token
			(
				access_token, 
			 	user_id, 
			 	client_id, 
				expires
			)
			VALUES(
				?,
				?,
				?,
				?
			);
		`

	err := s.session.Query(query,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires,
	).Exec()

	if err != nil {
		return resterrors.NewInternalServerError("Error 0003: " + err.Error())
	}

	return nil
}

func (s *accessTokenDBSession) UpdateExpirationTime(at entity.AccessToken) resterrors.RestErr {

	query := `
		UPDATE	access_token
		SET		expires 		= ?
		WHERE 	access_token 	= ?
			
		`
	err := s.session.Query(query,
		at.Expires,
		at.AccessToken,
	).Exec()

	if err != nil {
		return resterrors.NewInternalServerError("Error 0004: " + err.Error())
	}

	return nil
}
