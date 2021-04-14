package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)


type Buy_SamsicoinUsecase interface {
	InsertUser(user *model.Samsicoin) (use *model.GetaUser, err error )
}

type Buy_SamsicoinUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var Buy_Samsicoin *UserUsecaseSqlServer

func NewBuy_SamsicoinUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertBuy_Samsicoin(user *model.Buy_Samsicoin) (use *model.GetaUser, err error ) {
	return u.userRepositorySqlServer.InsertBuy_Samsicoin(user)

}
