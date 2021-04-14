package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type Buy_ChargRepository interface {
	Insert(user *model.Buy_Charg) (use *model.GetaUser, err error  )
}
type Buy_ChargRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewBuy_ChargRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}

func (r *UserRepositorySqlServer) InsertBuy_Charg(user *model.Buy_Charg) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec buy_charg_insert ?,?,?",
	user.User_Id,
	user.Phone_Number,
	user.Pric,
		
	).Scan( &a.Error_cod , &a.Id )
	fmt.Println(a.Id)
	if err != nil {
		
		return nil, err 
	}
		 return &a, nil 
}
	
	



