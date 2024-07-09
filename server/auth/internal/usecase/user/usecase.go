package user

import "smolf-auth/internal/repository/db"

type Usecase struct {
	Db *db.Repository
}

func NewUsecase(db *db.Repository) *Usecase {
	return &Usecase{
		Db: db,
	}
}
