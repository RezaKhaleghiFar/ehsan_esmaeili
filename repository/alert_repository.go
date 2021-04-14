package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type AlertRepository interface {
	Insert(user *model.Alert) (use *model.GetaUser, err error  )
}
type AlertRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewAlertRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}

func (r *UserRepositorySqlServer) InsertAlert(user *model.Alert) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec alert_insert ?,?,?,?,?",
	user.Coin_Id,
	user.Pric,
	user.User_Id,
	user.Activ,
	user.Time,
		
	).Scan( &a.Error_cod , &a.Id )
	fmt.Println(a.Id)
	if err != nil {
		
		return nil, err 
	}
		 return &a, nil 
}
	
	



