package db

import (
	"github.com/andrestor2/bookstore_oauth-api/src/clients/cassandra"
	"github.com/andrestor2/bookstore_oauth-api/src/domain/access_token"
	"github.com/andrestor2/bookstore_users-api/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (repository *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSessions()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
