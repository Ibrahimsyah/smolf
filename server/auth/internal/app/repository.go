package app

import "smolf-auth/internal/repository/db"

func NewRepositories(param RepositoryParam) *Repositories {
	return &Repositories{
		Db: db.NewRepository(),
	}
}
