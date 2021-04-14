package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)


type ClearingUsecase interface {
	InsertUser(user *model.Samsicoin) (use *model.GetaUser, err error )
}

type ClearingUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var Clearing *UserUsecaseSqlServer

func NewClearingUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertClearing(user *model.Clearing) (use *model.GetaUser, err error ) {
	return u.userRepositorySqlServer.InsertClearing(user)

}
