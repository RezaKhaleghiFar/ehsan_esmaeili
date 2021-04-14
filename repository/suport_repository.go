package repository

import (
	"database/sql"
	"ehsan_esmaeili/model"
	"fmt"
)


type SuportRepository interface {
	Insert(user *model.Suport) (use *model.GetaUser, err error  )
}
type SuportRepositorySqlServer struct {
	db    *sql.DB //64b
	table string  //4b
	//68
}

func NewSuportRepositorySqlServer(table string, db *sql.DB) *UserRepositorySqlServer {
	return &UserRepositorySqlServer{
		db:    db,
		table: table,
	}
}

func (r *UserRepositorySqlServer) InsertSuport(user *model.Suport) (use *model.GetaUser, err error ) {
	var a model.GetaUser
    
	err = r.db.QueryRow(" Exec Suports_insert ?,?",
	user.Support_Name,
	user.Support_Link,
		
	).Scan( &a.Error_cod , &a.Id )
	fmt.Println(a.Id)
	if err != nil {
		
		return nil, err 
	}

	
		 return &a, nil 
}
	
	



