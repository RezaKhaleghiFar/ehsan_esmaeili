package main

import (
	"ehsan_esmaeili/config"
	"ehsan_esmaeili/route"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	fmt.Println("Server Run On the : %s", config.Server_Url)
	route.RouterInit(router)
	log.Fatal(http.ListenAndServe(":"+config.Server_Port, router))

}
