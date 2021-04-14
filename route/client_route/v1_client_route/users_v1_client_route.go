package v1_client_route

import (
	"ehsan_esmaeili/api/client_api/v1_client_api"

	"github.com/julienschmidt/httprouter"
)

const path string = "/api/client/v1/users"

func Users_v1_client_Init(r *httprouter.Router) {
	r.GET("/api/client/v1/token", v1_client_api.GetToken)
	r.GET(path+"/:id/:name", v1_client_api.GetAllUsers)
	r.POST(path, v1_client_api.AddUsers)
	r.POST(path+"/aid", v1_client_api.AddAidDB)
	r.POST(path+"/user", v1_client_api.AddUserDB)
	r.POST(path+"/user_coin", v1_client_api.AddUser_CoinDB)
	r.POST(path+"/transaction", v1_client_api.AddTransactionDB)
	r.POST(path+"/samsicoin", v1_client_api.AddSamsicoinDB)
	r.POST(path+"/energy", v1_client_api.AddEnergyDB)
	r.POST(path+"/clearing", v1_client_api.AddClearingDB)
	r.POST(path+"/buy_samsicoin", v1_client_api.AddBuy_SamsicoinDB)
	r.POST(path+"/buy_pack", v1_client_api.AddBuy_PackDB)
	r.POST(path+"/buy_charg", v1_client_api.AddBuy_ChargDB)
	r.POST(path+"/alert", v1_client_api.AddAlertDB)
	r.POST(path+"/sms", v1_client_api.SMS)
	r.PUT(path+"/:id", v1_client_api.AddUsers)
	r.DELETE(path+"/:id", v1_client_api.AddUsers)
}
