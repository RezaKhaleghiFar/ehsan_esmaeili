package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type EnergyRepository interface {
	Insert(user *model.Energy) (use *model.GetaUser, err error  )
}
type EnergyRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewEnergyRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}

func (r *UserRepositorySqlServer) InsertEnergy(user *model.Energy) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec energys_insert ?,?,?,?",
	user.Low_High,
	user.Description,
	user.User_Id,
	user.Low_High_Count,
		
	).Scan( &a.Error_cod , &a.Id )
	fmt.Println(a.Id)
	//fmt.Sprintln(data)
	//a, err := data.()
	//  fmt.Println(user.Low_High)
	//  fmt.Println(user.User_Id)
	//  fmt.Println(user.Description)
	// fmt.Println(user.Low_High_Count)
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
	
	



