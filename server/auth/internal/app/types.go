package app

import (
	"smolf-auth/internal/repository/db"
	"smolf-auth/internal/usecase/user"
)

type RepositoryParam struct{}

type Repositories struct {
	Db *db.Repository
}

type Usecases struct {
	User *user.Usecase
}
