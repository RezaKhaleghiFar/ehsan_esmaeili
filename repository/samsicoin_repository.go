package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type SamsicoinRepository interface {
	Insert(user *model.Samsicoin) (use *model.GetaUser, err error  )
}
type SamsicoinRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewSamsicoinRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}

func (r *UserRepositorySqlServer) InsertSamsicoin(user *model.Samsicoin) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec samsicoin_insert ?,?,?,?,?,?",
	user.Low_High,
	user.Description,
	user.User_Id,	
	user.Low_High_Count,
	user.Transaction_Id,
	user.Typ,
		
	).Scan( &a.Error_cod , &a.Id )
	fmt.Println(a.Id)
	//fmt.Sprintln(data)
	//a, err := data.()
	// fmt.Println(user.Coin_Id)
	// //fmt.Println(a)
//	 fmt.Println(user.User_Id)
	if err != nil {
		
		return nil, err 
	}
	// else if errorcod != 0{
	// 	return "", nil 
	// }
//	else{
	
		 return &a, nil 
	//}

	
 	
		//err.Scan(errorcod)
}
	
	



