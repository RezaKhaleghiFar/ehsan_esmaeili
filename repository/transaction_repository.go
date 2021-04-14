package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type TransactionRepository interface {
	Insert(user *model.Transaction) (use *model.GetaUser, err error  )
}
type TransactionRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewTransactionRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}

func (r *UserRepositorySqlServer) InsertTransaction(user *model.Transaction) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec transaction_insert ?,?,?,?,?,?",
	user.Transaction_Id,
	user.Transaction_Time,
	user.Transaction_Price,	
	user.Transaction_User_Id,
	user.Tracking_Number,
	user.Transaction_IsSuccessful,
		
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
	
	



