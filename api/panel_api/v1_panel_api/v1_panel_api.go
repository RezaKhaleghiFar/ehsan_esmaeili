package v1_client_api
import (
//	"ehsan_esmaeili/function"
	"ehsan_esmaeili/middleware"
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/usecase"
	"encoding/json"
	"fmt"
	"net/http"

//	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)
func AddSuportDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var g model.Suport

	_ = json.NewDecoder(r.Body).Decode(&g)
	id, err := usecase.UserUse.InsertSuport(&model.Suport{
		Support_Name:        g.Support_Name,
		Support_Link:       g.Support_Link,
	})

	if err != nil {
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("error")
	   fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
	
	return
}