package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)


type Buy_ChargUsecase interface {
	InsertUser(user *model.Buy_Charg) (use *model.GetaUser, err error )
}

type Buy_ChargUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var Buy_Charg *UserUsecaseSqlServer

func NewBuy_ChargUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertBuy_Charg(user *model.Buy_Charg) (use *model.GetaUser, err error ) {
	return u.userRepositorySqlServer.InsertBuy_Charg(user)

}
