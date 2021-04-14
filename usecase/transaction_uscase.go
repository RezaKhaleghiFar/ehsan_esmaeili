package usecase

import (
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/repository"
)


type TransactionUsecase interface {
	InsertUser(user *model.Transaction) (use *model.GetaUser, err error )
}

type TransactionUsecaseSqlServer struct {
	userRepositorySqlServer *repository.UserRepositorySqlServer
}

var TransactionUse *UserUsecaseSqlServer

func NewTransactionUsecaseSqlServer(userRepositorySqlServer *repository.UserRepositorySqlServer) *UserUsecaseSqlServer {
	return &UserUsecaseSqlServer{
		userRepositorySqlServer: userRepositorySqlServer,
	}
}
func (u *UserUsecaseSqlServer) InsertTransaction(user *model.Transaction) (use *model.GetaUser, err error ) {
	return u.userRepositorySqlServer.InsertTransaction(user)

}
