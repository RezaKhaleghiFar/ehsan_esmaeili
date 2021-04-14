package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)

type AidUsecase interface {
	InsertAid(user *model.Aid) (userId int, err error)
}

type AidUsecaseSqlServer struct {
	aidRepositorySqlServer *repository.UserRepositorySqlServer
}

var AidUse *AidUsecaseSqlServer

func NewUserAidcaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertAid(aid *model.Aid) (userId int, err error) {
	return u.userRepositorySqlServer.InsertAid(aid)

}
