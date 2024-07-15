package user

type RepositoryGRPCProvider interface{}

type Usecase struct {
	repoGRPC RepositoryGRPCProvider
}

func NewUsecase(repoGRPC RepositoryGRPCProvider) *Usecase {
	return &Usecase{
		repoGRPC: repoGRPC,
	}
}
