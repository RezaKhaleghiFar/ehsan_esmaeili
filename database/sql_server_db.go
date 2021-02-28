package database

import (
	"database/sql"
	"fmt"
)

const (
	sqlserver_host     string = "."
	sqlserver_port     int    = 1433
	sqlserver_username string = "sa"
	sqlserver_password string = "sa"
	sqlserver_database string = ""
)

func ConnectToSqlServerDB() (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		sqlserver_host, sqlserver_username, sqlserver_password, sqlserver_port, sqlserver_database)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Errorf("error is db ", err)
		return nil, err
	}
	return conn, nil

	// var t string = "reza is " + "good"
	// var t2 string = fmt.Sprintf("reza is %d", "good")
	// fmt.Println(t)
	// fmt.Println(t2)
}
