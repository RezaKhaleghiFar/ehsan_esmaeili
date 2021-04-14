package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type User_CoinRepository interface {
	Insert(user *model.User_Coin) (use *model.GetaUser, err error  )
}
type User_coinRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewUser_CoinRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}

func (r *UserRepositorySqlServer) InsertUsercoin(user *model.User_Coin) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec user_coins_insert ?,?",
		
	user.User_Id,
	user.Coin_Id,
		
	).Scan(&a.Id)
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
	
	



