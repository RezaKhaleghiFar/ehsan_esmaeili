package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
)

type UserRepository interface {
	Insert(user *model.User) (userId int, err error)
}
type UserRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewUserRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}
func (r *UserRepositorySqlServer) Insert(user *model.User) (userId int, err error) {
	err = r.db.QueryRow("Exec public.sp_point_insert($1,$2,$3,$4)",
		user.Id,
		user.Name,
		user.Family,

		userId).Scan(&userId)

	return userId, err

}
