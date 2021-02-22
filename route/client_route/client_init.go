package client_route

import (
	"ehsan_esmaeili/route/client_route/v1_client_route"

	"github.com/julienschmidt/httprouter"
)

func ClientInit(r *httprouter.Router) {
	v1_client_route.V1ClientInit(r)
}
