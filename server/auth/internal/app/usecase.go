package app

import "smolf-auth/internal/usecase/user"

func NewUsecase(repositories *Repositories) *Usecases {
	return &Usecases{
		User: user.NewUsecase(repositories.Db),
	}
}
