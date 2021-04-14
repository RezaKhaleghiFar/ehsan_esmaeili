package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
)

type AidRepository interface {
	Insert(user *model.Aid) (userId int, err error)
}
type AidRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewAidRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}
func (r *UserRepositorySqlServer) InsertAid(user *model.Aid) (userId int, err error) {
	err = r.db.QueryRow(" Exec aid_insert ?,?,?,?",
		user.User_Id,
		user.Pric,
		user.Aid_description,
		user.Transaction_id,
	).Scan(&userId)
	// fmt.Println(err)
	// //fmt.Sprintln(data)
	// //a, err := data.()
	// fmt.Println(err)
	// //fmt.Println(a)
	// fmt.Println(userId)
	if err != nil {
		return 0, err
	}
	return userId, nil

}
