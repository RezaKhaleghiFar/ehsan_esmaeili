package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)


type Buy_PackUsecase interface {
	InsertUser(user *model.Samsicoin) (use *model.GetaUser, err error )
}

type Buy_PackUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var Buy_Pack *UserUsecaseSqlServer

func NewBuy_PackUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertBuy_Pack(user *model.Buy_Pack) (use *model.GetaUser, err error ) {
	return u.userRepositorySqlServer.InsertBuy_Pack(user)

}
