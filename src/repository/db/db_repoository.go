package db

import (
	"errors"
	"github.com/andrestor2/bookstore_oauth-api/src/clients/cassandra"
	"github.com/andrestor2/bookstore_oauth-api/src/domain/access_token"
	"github.com/andrestor2/bookstore_utils-go/rest_errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token =?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires = ? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(token access_token.AccessToken) rest_errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) rest_errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, rest_errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSessions().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no access token found with given id")
		}
		return nil, rest_errors.NewInternalServerError("error with cassandra's query", errors.New("database error"))
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) rest_errors.RestErr {
	if err := cassandra.GetSessions().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error with cassandra's query", errors.New("database error"))
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) rest_errors.RestErr {
	if err := cassandra.GetSessions().Query(queryUpdateExpires,
		at.AccessToken,
		at.Expires,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error cassandra's expiration query", errors.New("database error"))
	}
	return nil
}
