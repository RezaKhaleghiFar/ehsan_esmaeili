package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)


type AlertUsecase interface {
	InsertUser(user *model.Samsicoin) (use *model.GetaUser, err error )
}

type AlertUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var Alert *UserUsecaseSqlServer

func NewAlertUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertAlert(user *model.Alert) (use *model.GetaUser, err error ) {
	return u.userRepositorySqlServer.InsertAlert(user)

}
