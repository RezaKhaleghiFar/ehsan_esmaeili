package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type ClearingRepository interface {
	Insert(user *model.Clearing) (use *model.GetaUser, err error  )
}
type ClearingRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewClearingRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}

func (r *UserRepositorySqlServer) InsertClearing(user *model.Clearing) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec clearing_insert ?,?,?,?",
	user.User_Id,
	user.Samsicoin_Count,
	user.Price,
	user.Time,
		
	).Scan( &a.Error_cod , &a.Id )
	fmt.Println(a.Id)
	if err != nil {
		
		return nil, err 
	}
		 return &a, nil 
}
	
	



