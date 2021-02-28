package main

import (
	"ehsan_esmaeili/config"
	"ehsan_esmaeili/database"
	"ehsan_esmaeili/route"
	"ehsan_esmaeili/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	fmt.Println("Server Run On the : %s", config.Server_Url)
	db, err := database.ConnectToSqlServerDB()
	if err != nil {
		fmt.Errorf("eroro sql", err)
	}
	usecase.UsecaseInit(db)

	route.RouterInit(router)
	log.Fatal(http.ListenAndServe(":"+config.Server_Port, router))

}
