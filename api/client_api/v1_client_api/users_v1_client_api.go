package v1_client_api

import (
	"ehsan_esmaeili/function"
	"ehsan_esmaeili/middleware"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		return
	}
	//header
	fmt.Println("All Header")
	fmt.Println(r.Header)
	fmt.Println("get one Header")
	fmt.Println("header --> User-Agent : ", r.Header.Get("User-Agent"))
	fmt.Println("****************************")

	//query -> ex. http://localhost:8080/api/client/v1/users?id=10&name=reza
	fmt.Println("All Query")
	fmt.Println(r.URL.Query())
	fmt.Println("get one Query")
	fmt.Println("query --> id : ", r.URL.Query().Get("id"))
	fmt.Println("query --> name : ", r.URL.Query().Get("name"))
	fmt.Println("****************************")

	//path -> ex. http://localhost:8080/api/client/v1/users/id/name
	//path -> ex. http://localhost:8080/api/client/v1/users/20/ali
	fmt.Println("all path")
	fmt.Println(p)
	fmt.Println("get one path")
	fmt.Println("path --> id : ", p.ByName("id"))
	fmt.Println("path --> name : ", p.ByName("name"))
	fmt.Println("****************************")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("")
	return
}

func AddUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//header
	fmt.Println("All Header")
	fmt.Println(r.Header)
	fmt.Println("get one Header")
	fmt.Println("header --> User-Agent : ", r.Header.Get("User-Agent"))
	fmt.Println("****************************")

	//formdata -> ex. POST: http://localhost:8080/api/client/v1/users
	fmt.Println("get one Form Data")
	fmt.Println("body - form data --> name : ", r.FormValue("name"))
	fmt.Println("body - form data --> name : ", r.FormValue("name"))

	fmt.Println("****************************")
	fmt.Println("get model json")
	var u getUser
	_ = json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("body - json --> all : ", u)
	fmt.Println("get one json")
	fmt.Println("body - json --> name : ", u.Name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("")
	return
}
func GetToken(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := function.GenerateJWT()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
	return
}

type getUser struct {
	Name   string
	Family string
}
