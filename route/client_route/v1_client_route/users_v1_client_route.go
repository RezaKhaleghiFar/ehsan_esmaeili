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
	r.PUT(path+"/:id", v1_client_api.AddUsers)
	r.DELETE(path+"/:id", v1_client_api.AddUsers)
}
