package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)

type UserUsecase interface {
	InsertUser(user *model.User)
}

type UserUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var UserUse *UserUsecaseSqlServer

func NewUserUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertUser(user *model.User) {
	u.InsertUser(user)
}
