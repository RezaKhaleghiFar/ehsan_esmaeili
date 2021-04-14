package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)


type SamsicoinUsecase interface {
	InsertUser(user *model.Samsicoin) (use *model.GetaUser, err error )
}

type SamsicoinUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var Samsicoin *UserUsecaseSqlServer

func NewSamsicoinUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertSamsicoin(user *model.Samsicoin) (use *model.GetaUser, err error ) {
	return u.userRepositorySqlServer.InsertSamsicoin(user)

}
