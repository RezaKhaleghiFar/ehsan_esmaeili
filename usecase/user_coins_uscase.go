package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)


type UsercoinUsecase interface {
	InsertUser(user *model.User_Coin) (use *model.GetaUser, err error )
}

type UsercoinUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var UsercoinUse *UserUsecaseSqlServer

func NewUsercoinUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertUsercoin(user *model.User_Coin) (use *model.GetaUser, err error ) {
	return u.userRepositorySqlServer.InsertUsercoin(user)

}
