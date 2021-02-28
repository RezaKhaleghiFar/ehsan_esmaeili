package usecase

import (
	"database/sql"
	"ehsan_esmaeili/repository"
)

func UsecaseInit(db *sql.DB) {
	userRepo := repository.NewUserRepositorySqlServer("user", db)
	UserUse = NewUserUsecaseSqlServer(userRepo)
}
