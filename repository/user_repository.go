package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type UserRepository interface {
	Insert(user *model.User) (use *model.GetaUser, err error  )
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

func (r *UserRepositorySqlServer) Insert(user *model.User) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec users_insert ?,?,?,?,?,?,?,?,?,?,?,?,?",
		user.User_Phon,
		user.Name,
		user.Last_Name,
		user.Post_Cod,
		user.State,
		user.City,
		user.Locition,
		user.Shaba,
		user.Energy_Count,
		user.Samsicoin_Count,
		user.Alert_Count,
		user.User_Id,
		user.Active
	).Scan(&a.Error_cod , &a.Id)
	fmt.Println(user.User_Phon)
	//fmt.Sprintln(data)
	//a, err := data.()
	 fmt.Println(err)
	// //fmt.Println(a)
	// fmt.Println(userId)
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
	
	



