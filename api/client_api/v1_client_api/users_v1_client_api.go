package v1_client_api

import (
	"ehsan_esmaeili/function"
	"ehsan_esmaeili/middleware"
	"ehsan_esmaeili/model"
	"ehsan_esmaeili/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"math/rand"
	"strings"
	"strconv"
	"io/ioutil"
	"github.com/google/uuid"
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
	var u getAid
	_ = json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("body - json --> all : ", u)
	fmt.Println("get one json")
	fmt.Println("body - json --> name : ", u.User_Id)
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

type getAid struct {
	User_Id         string
	Price           int
	Aid_description string
	Transaction_id  string
}
type getUser struct {
	User_Phon       string
	Name            string
	Last_Name       string
	Post_Cod        string
	State           string
	City            string
	Locition        string
	Shaba           string
	Energy_Count    int
	Samsicoin_Count float64
	Alert_Count     int
}
type getUser_coin struct {
	Coin_Id       int
	User_Id     string
}
type getTransaction struct {
	Transaction_User_Id       string
	Transaction_Price     int
	Transaction_Time      string
}
type getSamsicoin struct {
	User_Id        string
	Description    string
	Low_High       bool
	Low_High_Count int
	Transaction_Id string
	Typ            string
}
type getEnergy struct {
	User_Id        string
	Description    string
	Low_High       bool
	Low_High_Count int
}
type getClearing struct {
	User_Id         string
	Time            string
	Samsicoin_Count float32
	Price           int
}

func AddAidDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var u getAid
	_ = json.NewDecoder(r.Body).Decode(&u)
	id, err := usecase.UserUse.InsertAid(&model.Aid{

		User_Id:         u.User_Id,
		Aid_description: u.Aid_description,
		Pric:            u.Price,
		Transaction_id:  u.Transaction_id,
	})
	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)

	return
}

func AddUserDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var u model.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	uuidWithHyphen := uuid.New()
	id, err := usecase.UserUse.InsertUser(&model.User{

		User_Phon:       u.User_Phon,
		Name:            u.Name,
		Last_Name:       u.Last_Name,
		Post_Cod:        u.Post_Cod,
		State:           u.State,
		City:            u.City,
		Locition:        u.Locition,
		Shaba:           u.Shaba,
		Energy_Count:    u.Energy_Count,
		Samsicoin_Count: u.Samsicoin_Count,
		Alert_Count:     u.Alert_Count,
		User_Id:         uuidWithHyphen.String(),
		Active: false,
	})
	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("error")
	   fmt.Println(err.Error())
		return
	}
	url := "http://rest.ippanel.com/v1/messages"
	method := "POST"
	//fmt.Print((rand.Intn(99999-11111)+11111),",")
	kod := strconv.Itoa(rand.Intn(99999-11111)+11111)
	fmt.Println(kod)
	fmt.Println(r.FormValue("Recipients"))
	payload := strings.NewReader(
		`{`+""+`"originator": "+983000505",`+""+`"recipients":`+""+`[`+""+`"`+u.User_Phon+`"`+""+`],`+""+`"message":`+""+`"`+kod+`"`+""+`}`+""+` `)
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	  return
	}
	req.Header.Add("Authorization", "AccessKey P05MzJyAA1SRIzOOMFzPb95Bcn4aCB_xaps_zoPcOYQ=")
	req.Header.Add("Content-Type", "application/json")
  
	res, err := client.Do(req)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println(string(body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
	
	return
}

func AddUser_CoinDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var u getUser_coin

	_ = json.NewDecoder(r.Body).Decode(&u)
	id, err := usecase.UserUse.InsertUsercoin(&model.User_Coin{
		User_Id:        u.User_Id,
		Coin_Id:       u.Coin_Id,
		
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
func AddTransactionDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var u getTransaction
	_ = json.NewDecoder(r.Body).Decode(&u)
	uuidWithHyphen := uuid.New()
	id, err := usecase.UserUse.InsertTransaction(&model.Transaction{
		Transaction_Id:   uuidWithHyphen.String(),
		Transaction_User_Id:        u.Transaction_User_Id,
		Transaction_Price:       u.Transaction_Price,
		Transaction_Time:        u.Transaction_Time,
		
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
func AddSamsicoinDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var u getSamsicoin

	_ = json.NewDecoder(r.Body).Decode(&u)
	id, err := usecase.UserUse.InsertSamsicoin(&model.Samsicoin{
		Low_High:        u.Low_High,
		Description:       u.Description,
		User_Id: u.User_Id,
		Low_High_Count: u.Low_High_Count,
		Transaction_Id: u.Transaction_Id,
		Typ: u.Typ,
		
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
func AddEnergyDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var g getEnergy

	_ = json.NewDecoder(r.Body).Decode(&g)
	id, err := usecase.UserUse.InsertEnergy(&model.Energy{
		Low_High:        g.Low_High,
		Description:       g.Description,
		User_Id: g.User_Id,
		Low_High_Count: g.Low_High_Count,
	
		
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
func AddClearingDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var g getClearing

	_ = json.NewDecoder(r.Body).Decode(&g)
	id, err := usecase.UserUse.InsertClearing(&model.Clearing{
		User_Id: g.User_Id,
		Samsicoin_Count:        g.Samsicoin_Count,
		Price:       g.Price,
		Time: g.Time,
	
		
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
func AddBuy_SamsicoinDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var g model.Buy_Samsicoin

	_ = json.NewDecoder(r.Body).Decode(&g)
	id, err := usecase.UserUse.InsertBuy_Samsicoin(&model.Buy_Samsicoin{
		User_Id: g.User_Id,
		Samsicoin_Count:        g.Samsicoin_Count,
		Pric:       g.Pric,
		Transaction_Id: g.Transaction_Id,
	
		
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
func AddBuy_PackDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var g model.Buy_Pack

	_ = json.NewDecoder(r.Body).Decode(&g)
	id, err := usecase.UserUse.InsertBuy_Pack(&model.Buy_Pack{
		User_Id: g.User_Id,
		Pack_Id:        g.Pack_Id,
		Time:       g.Time,
		Transaction_Id: g.Transaction_Id,
	
		
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
func AddBuy_ChargDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var g model.Buy_Charg
	_ = json.NewDecoder(r.Body).Decode(&g)
	id, err := usecase.UserUse.InsertBuy_Charg(&model.Buy_Charg{
		User_Id: g.User_Id,
		Phone_Number:       g.Phone_Number,
		Pric: g.Pric,
	
		
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
func AddAlertDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	var g model.Alert

	_ = json.NewDecoder(r.Body).Decode(&g)
	id, err := usecase.UserUse.InsertAlert(&model.Alert{
		User_Id: g.User_Id,
		Coin_Id:       g.Coin_Id,
		Pric: g.Pric,
		Time: g.Time,
		Activ: g.Activ,
	
		
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
func SMS(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	err := middleware.CheckAuthentication(w, r)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
		return
	}
	url := "http://rest.ippanel.com/v1/messages"
	method := "POST"
	//fmt.Print((rand.Intn(99999-11111)+11111),",")
	kod := strconv.Itoa(rand.Intn(99999-11111)+11111)
	fmt.Println(kod)
	fmt.Println(r.FormValue("Recipients"))
	payload := strings.NewReader(
		`{`+""+`"originator": "+983000505",`+""+`"recipients":`+""+`[`+""+r.FormValue("Recipients")+""+`],`+""+`"message":`+""+`"`+kod+`"`+""+`}`+""+` `)
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	  return
	}
	req.Header.Add("Authorization", "AccessKey P05MzJyAA1SRIzOOMFzPb95Bcn4aCB_xaps_zoPcOYQ=")
	req.Header.Add("Content-Type", "application/json")
  
	res, err := client.Do(req)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println(string(body))
  }