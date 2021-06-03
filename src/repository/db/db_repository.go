package db

import (
	"github.com/TriTranDev/bookstore_oauth-api/src/domain/access_token"
	"github.com/TriTranDev/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string2 string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalserverError("database connection not implement")
	//return nil, nil
}
