package auth

type UserUsecase interface{}

type Handler struct {
	UserUC UserUsecase
}

func NewHandler(userUC UserUsecase) *Handler {
	return &Handler{
		UserUC: userUC,
	}
}
