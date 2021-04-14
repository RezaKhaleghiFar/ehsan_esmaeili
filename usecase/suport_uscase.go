package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)


type SuportUsecase interface {
	InsertUser(user *model.Samsicoin) (use *model.GetaUser, err error )
}

type SuportUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var Suport *UserUsecaseSqlServer

func NewSuportUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertSuport(user *model.Suport) (use *model.GetaUser, err error ) {
	return u.userRepositorySqlServer.InsertSuport(user)

}
