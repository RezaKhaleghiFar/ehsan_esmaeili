package v1_client_route

import "github.com/julienschmidt/httprouter"

func V1ClientInit(r *httprouter.Router) {
	Users_v1_client_Init(r)
}
