package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type Buy_SamsicoinRepository interface {
	Insert(user *model.Buy_Samsicoin) (use *model.GetaUser, err error  )
}
type Buy_SamsicoinRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewBuy_SamsicoinRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}

func (r *UserRepositorySqlServer) InsertBuy_Samsicoin(user *model.Buy_Samsicoin) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec buy_samsicoin_insert ?,?,?,?",
	user.User_Id,
	user.Samsicoin_Count,
	user.Pric,
	user.Transaction_Id,
		
	).Scan( &a.Error_cod , &a.Id )
	fmt.Println(a.Id)
	if err != nil {
		
		return nil, err 
	}
		 return &a, nil 
}
	
	



