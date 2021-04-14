package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type Buy_PackRepository interface {
	Insert(user *model.Buy_Pack) (use *model.GetaUser, err error  )
}
type Buy_PackRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewBuy_PackRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}

func (r *UserRepositorySqlServer) InsertBuy_Pack(user *model.Buy_Pack) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec buy_pack_insert ?,?,?,?",
	user.Pack_Id,
	user.User_Id,
	user.Transaction_Id,
	user.Time,
		
	).Scan( &a.Error_cod , &a.Id )
	fmt.Println(a.Id)
	if err != nil {
		
		return nil, err 
	}
		 return &a, nil 
}
	
	



