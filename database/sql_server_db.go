package database

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	sqlserver_host     string = "."
	sqlserver_port     int    = 1433
	sqlserver_username string = "esi"
	sqlserver_password string = "1"
	sqlserver_database string = "Samsonet"
)

func ConnectToSqlServerDB() (*sql.DB, error) {
	connString := fmt.Sprintf("user=%s;password=%s;database=%s;",
		sqlserver_username, sqlserver_password, sqlserver_database)
	condb, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Errorf("error is db ", err)

		return nil, err
	}
	return condb, nil

	// var t string = "reza is " + "good"
	// var t2 string = fmt.Sprintf("reza is %d", "good")
	// fmt.Println(t)
	// fmt.Println(t2)
}
