package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)


type EnergyUsecase interface {
	InsertUser(user *model.Samsicoin) (use *model.GetaUser, err error )
}

type EnergyUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var Energy *UserUsecaseSqlServer

func NewEnergyUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertEnergy(user *model.Energy) (use *model.GetaUser, err error ) {
	return u.userRepositorySqlServer.InsertEnergy(user)

}
